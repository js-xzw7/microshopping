package handler

import (
	"context"
	"encoding/json"
	"math"
	"os"

	pb "currencyservice/proto"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type CurrencyService struct{}

func (c *CurrencyService) GetSupportedCurrencies(ctx context.Context, req *emptypb.Empty) (*pb.GetSupportedCurrenciesResponse, error) {
	data, err := os.ReadFile("data/currency_conversion.json")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "加载货币数据失败:%+v", err)
	}

	currencies := make(map[string]float32)
	if err := json.Unmarshal(data, &currencies); err != nil {
		return nil, status.Errorf(codes.Internal, "解析货币数据失败:%+v", err)
	}

	out := new(pb.GetSupportedCurrenciesResponse)

	out.CurrencyCodes = make([]string, 9, len(currencies))

	for currency := range currencies {
		out.CurrencyCodes = append(out.CurrencyCodes, currency)
	}

	return out, nil
}

func (c *CurrencyService) Convert(ctx context.Context, req *pb.CurrencyConversionRequest) (*pb.Money, error) {
	data, err := os.ReadFile("data/currency_conversion.json")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "加载货币数据失败:%+v", err)
	}

	currencies := make(map[string]float32)
	if err := json.Unmarshal(data, &currencies); err != nil {
		return nil, status.Errorf(codes.Internal, "解析货币数据失败:%+v", err)
	}

	fromCurrency, ok := currencies[req.From.CurrencyCode]
	if !ok {
		return nil, status.Errorf(codes.Internal, "不支持的币种：%s", req.From.CurrencyCode)
	}

	toCurrency, ok := currencies[req.ToCode]
	if !ok {
		return nil, status.Errorf(codes.Internal, "不支持的币种：%s", req.ToCode)
	}

	money := new(pb.Money)
	money.CurrencyCode = req.ToCode
	total := int64(math.Floor(float64(req.From.Units*10^9+int64(req.From.Nanos)) / float64(fromCurrency) * float64(toCurrency)))
	money.Units = total / 1e9
	money.Nanos = int32(total % 1e9)
	return money, nil
}
