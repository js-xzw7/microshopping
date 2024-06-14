package handler

import (
	"bytes"
	"context"
	"log"
	pb "paymentservice/proto"
	"strconv"

	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger:", log.Lshortfile)
)

type PaymentService struct{}

func (s *PaymentService) Charge(ctx context.Context, req *pb.ChargeRequest) (res *pb.ChargeResponse, err error) {
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.FormatInt(int64(req.CreditCard.CreaditCardCvv), 10),
		Year:   strconv.FormatInt(int64(req.CreditCard.CreaditCardExpirationYear), 10),
		Month:  strconv.FormatInt(int64(req.CreditCard.CreaditCardExpirationMonth), 10),
	}

	res = new(pb.ChargeResponse)

	if err := card.Validate(); err != nil {
		return res, status.Errorf(codes.InvalidArgument, err.Error())
	}

	logger.Printf(`事务处理：%s, Amount: %s%d.%d`, req.CreditCard.CreditCardNumber, req.Amount.CurrencyCode, req.Amount.Units, req.Amount.Nanos)
	res.TransactionId = uuid.NewString()
	return res, nil
}
