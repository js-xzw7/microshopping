package main

import (
	"context"
	pb "frontend/proto"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const avoidNoopCurrencyConversionRPC = false

func (fe *FrontendServer) getCurrencies(ctx context.Context) ([]string, error) {
	currs, err := fe.currencyService.GetSupportedCurrencies(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	var res []string
	for _, c := range currs.CurrencyCodes {
		if _, ok := whitelistedCurrencies[c]; ok {
			res = append(res, c)
		}
	}
	return res, nil
}

func (fe *FrontendServer) getProducts(ctx context.Context) ([]*pb.Product, error) {
	res, err := fe.productCatalogService.ListProducts(ctx, &emptypb.Empty{})
	return res.GetProducts(), err
}

func (fe *FrontendServer) getProduct(ctx context.Context, id string) (*pb.Product, error) {
	res, err := fe.productCatalogService.GetProduct(ctx, &pb.GetProductRequest{Id: id})
	return res, err
}

func (fe *FrontendServer) getCart(ctx context.Context, userID string) ([]*pb.CartItem, error) {
	res, err := fe.cartService.GetCart(ctx, &pb.GetCartRequest{UserId: userID})
	return res.GetItems(), err
}

func (fe *FrontendServer) emptyCart(ctx context.Context, userID string) error {
	_, err := fe.cartService.EmptyCart(ctx, &pb.EmptyCartRequest{UserId: userID})
	return err
}

func (fe *FrontendServer) insertCart(ctx context.Context, userID, productID string, quantity int32) error {
	_, err := fe.cartService.AddItem(ctx, &pb.AddItemRequest{
		UserId: userID,
		Item: &pb.CartItem{
			ProductId: productID,
			Quantity:  quantity,
		},
	})

	return err
}

func (fe *FrontendServer) convertCurrency(ctx context.Context, money *pb.Money, currency string) (*pb.Money, error) {
	if avoidNoopCurrencyConversionRPC && money.GetCurrencyCode() == currency {
		return money, nil
	}

	return fe.currencyService.Convert(ctx, &pb.CurrencyConversionRequest{
		From:   money,
		ToCode: currency,
	})
}

func (fe *FrontendServer) getShippingQuote(ctx context.Context, items []*pb.CartItem, currency string) (*pb.Money, error) {
	quote, err := fe.shippingService.GetQuote(ctx, &pb.GetQuoteRequest{Address: nil, Items: items})
	if err != nil {
		return nil, err
	}

	localized, err := fe.convertCurrency(ctx, quote.GetCostUsd(), currency)
	return localized, errors.Wrap(err, "failed to convert currency for shipping cost")
}

func (fe *FrontendServer) getRecommendations(ctx context.Context, userID string, productIDs []string) ([]*pb.Product, error) {
	res, err := fe.recommendationService.ListRecommendations(ctx, &pb.ListRecommendationsRequest{
		UserId:     userID,
		ProductIds: productIDs,
	})

	if err != nil {
		return nil, err
	}

	out := make([]*pb.Product, len(res.GetProductIds()))
	for i, v := range res.GetProductIds() {
		p, err := fe.getProduct(ctx, v)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get recommended product info (#%s)", v)
		}
		out[i] = p
	}

	if len(out) > 4 {
		out = out[:4]
	}

	return out, err
}

func (fe *FrontendServer) getAd(ctx context.Context, ctxKeys []string) ([]*pb.Ad, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	res, err := fe.adService.GetAds(ctx, &pb.AdRequest{
		ContextKeys: ctxKeys,
	})

	return res.GetAds(), errors.Wrap(err, "failed to get ads")
}

func (fe *FrontendServer) chooseAd(ctx context.Context, ctxKeys []string, log logrus.FieldLogger) *pb.Ad {
	ads, err := fe.getAd(ctx, ctxKeys)
	if err != nil {
		log.WithField("error", err).Warn("查询广告失败")
		return nil
	}

	if len(ads) == 0 {
		return nil
	}

	return ads[rand.Intn(len(ads))]
}
