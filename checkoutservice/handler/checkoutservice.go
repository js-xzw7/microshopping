package handler

import (
	"bytes"
	pb "checkoutservice/proto"
	"context"
	"fmt"
	"log"

	"checkoutservice/money"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger:", log.Lshortfile)
)

type CheckoutService struct {
	CartService           pb.CartServiceClient
	CurrencyService       pb.CurrencyServiceClient
	EmailService          pb.EmailServiceClient
	PaymentService        pb.PaymentServiceClient
	ProductCatalogService pb.ProductCatalogServiceClient
	ShippingService       pb.ShippingServiceClient
}

// 准备订单
type orderPrep struct {
	orderItems            []*pb.OrderItem
	cartItems             []*pb.CartItem
	shippingCostLocalized *pb.Money
}

// 获得用户购物车
func (s *CheckoutService) getUserCart(ctx context.Context, userID string) ([]*pb.CartItem, error) {
	cart, err := s.CartService.GetCart(ctx, &pb.GetCartRequest{UserId: userID})
	if err != nil {
		return nil, fmt.Errorf("获取用户购物车失败：%v", err)
	}

	return cart.GetItems(), nil
}

// 清空购物车
func (s *CheckoutService) emptyUserCart(ctx context.Context, userID string) error {

	if _, err := s.CartService.EmptyCart(ctx, &pb.EmptyCartRequest{UserId: userID}); err != nil {
		return fmt.Errorf("清空购物车失败: %v", err)
	}

	return nil
}

// 货币转换
func (s *CheckoutService) convertCurrency(ctx context.Context, from *pb.Money, toCurrency string) (*pb.Money, error) {
	result, err := s.CurrencyService.Convert(ctx, &pb.CurrencyConversionRequest{From: from, ToCode: toCurrency})
	if err != nil {
		return nil, fmt.Errorf("价格转换失败：%v", err)
	}
	return result, nil
}

// 准备订单项
func (s *CheckoutService) prepOrderItems(ctx context.Context, items []*pb.CartItem, userCurrency string) ([]*pb.OrderItem, error) {
	res := make([]*pb.OrderItem, len(items))
	for i, item := range items {
		product, err := s.ProductCatalogService.GetProduct(ctx, &pb.GetProductRequest{Id: item.GetProductId()})
		if err != nil {
			return nil, fmt.Errorf("获取商品失败：%v", item.GetProductId())
		}

		price, err := s.convertCurrency(ctx, product.GetPriceUsd(), userCurrency)
		if err != nil {
			return nil, fmt.Errorf("价格转换失败：%q to %s", item.GetProductId(), userCurrency)
		}

		res[i] = &pb.OrderItem{Item: item, Cost: price}
	}

	return res, nil
}

// 配送配额
func (s *CheckoutService) quoteShipping(ctx context.Context, address *pb.Address, items []*pb.CartItem) (*pb.Money, error) {
	shippingQuote, err := s.ShippingService.GetQuote(ctx, &pb.GetQuoteRequest{Address: address, Items: items})
	if err != nil {
		return nil, fmt.Errorf("配送配额失败：%v", err)
	}

	return shippingQuote.GetCostUsd(), nil
}

// 准备订单和配送
func (s *CheckoutService) prepareOrderItemsAndShippingQuoteFromCart(ctx context.Context, userID, userCurrency string, address *pb.Address) (orderPrep, error) {
	var res orderPrep

	cartItems, err := s.getUserCart(ctx, userID)
	if err != nil {
		return res, fmt.Errorf("购物车错误：%v", err)
	}

	orderItems, err := s.prepOrderItems(ctx, cartItems, userCurrency)
	if err != nil {
		return res, fmt.Errorf("准备订单失败：%v", err)
	}

	shippingUSD, err := s.quoteShipping(ctx, address, cartItems)
	if err != nil {
		return res, fmt.Errorf("配送配额失败：%v", err)
	}

	shippingPrice, err := s.convertCurrency(ctx, shippingUSD, userCurrency)
	if err != nil {
		return res, fmt.Errorf("价格转换失败：%v", err)
	}

	res.shippingCostLocalized = shippingPrice
	res.cartItems = cartItems
	res.orderItems = orderItems
	return res, nil
}

// 结算卡
func (s *CheckoutService) chargeCard(ctx context.Context, amount *pb.Money, paymentInfo *pb.CreditCardInfo) (string, error) {
	paymentResp, err := s.PaymentService.Charge(ctx, &pb.ChargeRequest{
		Amount:     amount,
		CreditCard: paymentInfo,
	})

	if err != nil {
		return "", fmt.Errorf("不能更换卡：%v", err)
	}
	return paymentResp.GetTransactionId(), nil
}

// 配送订单
func (s *CheckoutService) ShipOrder(ctx context.Context, address *pb.Address, items []*pb.CartItem) (string, error) {
	resp, err := s.ShippingService.ShipOrder(ctx, &pb.ShipOrderRequest{Address: address, Items: items})
	if err != nil {
		return "", fmt.Errorf("配送订单失败：%v", err)
	}

	return resp.GetTrackingId(), nil
}

// 发送确认消息
func (s *CheckoutService) sendOrderConfirmation(ctx context.Context, email string, order *pb.OrderResult) error {
	_, err := s.EmailService.SendOrderConfirmation(ctx, &pb.SendOrderConfirmationRequest{
		Email: email,
		Order: order,
	})

	if err != nil {
		return err
	}

	return nil
}

// 下订单
func (s *CheckoutService) PlaceOrder(ctx context.Context, req *pb.PlaceOrderRequest) (res *pb.PlaceOrderResponse, err error) {
	logger.Printf("[PlaceOrder] user_id=%q user_currency=%q", req.UserId, req.UserCurrency)

	res = new(pb.PlaceOrderResponse)
	orderId, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "生成订单id失败")
	}

	prep, err := s.prepareOrderItemsAndShippingQuoteFromCart(ctx, req.UserId, req.UserCurrency, req.Address)
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	total := &pb.Money{CurrencyCode: req.UserCurrency, Units: 0, Nanos: 0}
	total = money.Must(money.Sum(total, prep.shippingCostLocalized))
	for _, item := range prep.orderItems {
		multPrice := money.MultiplySlow(item.Cost, uint32(item.GetItem().GetQuantity()))
		total = money.Must(money.Sum(total, multPrice))
	}

	txID, err := s.chargeCard(ctx, total, req.CreditCard)
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "更换卡失败：%v", err)
	}

	logger.Printf("付款交易 (transaction_id: %s)", txID)

	shippingTrackingID, err := s.ShipOrder(ctx, req.Address, prep.cartItems)
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "配送错误：%v", err)
	}

	if err := s.emptyUserCart(ctx, req.UserId); err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "清空购物车失败：%v", err)
	}

	orderRes := &pb.OrderResult{
		OrderId:            orderId.String(),
		ShippingTrackingId: shippingTrackingID,
		ShippingCost:       prep.shippingCostLocalized,
		ShippingAddress:    req.Address,
		Items:              prep.orderItems,
	}

	if err := s.sendOrderConfirmation(ctx, req.Email, orderRes); err != nil {
		logger.Printf("发送订单确认消息失败：%q: %+v", req.Email, err)
		log.Fatal(err)
	} else {
		logger.Printf("发送订单确认消息成功：%q", req.Email)
		log.Printf("发送订单确认消息成功：%s", req.Email)
	}

	res.Order = orderRes
	return res, nil
}
