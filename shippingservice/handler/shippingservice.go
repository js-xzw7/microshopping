package handler

import (
	"bytes"
	"context"
	"fmt"
	"log"
	pb "shippingservice/proto"
)

type ShippingService struct{}

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger:", log.Lshortfile)
)

func (s *ShippingService) GetQuote(ctx context.Context, req *pb.GetQuoteRequest) (res *pb.GetQuoteResponse, err error) {
	logger.Print("[GetQuote] 收到请求")
	defer logger.Print("[GetQuote]完成请求")

	//1.根据商品数量生成报价
	res = new(pb.GetQuoteResponse)
	quote := CreateQuoteFromCount(0)

	//2.生成响应
	res.CostUsd = &pb.Money{
		CurrencyCode: "USD",
		Units:        int64(quote.Dollars),
		Nanos:        int32(quote.Cents * 10000000),
	}

	return res, nil
}

func (s *ShippingService) ShipOrder(ctx context.Context, req *pb.ShipOrderRequest) (res *pb.ShipOrderResponse, err error) {
	logger.Print("[ShipOrder] 收到请求")
	defer logger.Print("[ShipOrder]完成请求")

	//1.创建跟踪id
	res = new(pb.ShipOrderResponse)
	baseAddress := fmt.Sprintf("%s,%s,%s", req.Address.StreetAddress, req.Address.City, req.Address.State)
	id := CreateTrackingId(baseAddress)

	//2.生成响应
	res.TrackingId = id
	return res, nil
}
