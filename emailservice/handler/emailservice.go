package handler

import (
	"bytes"
	"context"
	pb "emailservice/proto"
	"log"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type DummyEmailService struct{}

// 日志
var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger:", log.Lshortfile)
)

func (s *DummyEmailService) SendOrderConfirmation(ctx context.Context, req *pb.SendOrderConfirmationRequest) (res *emptypb.Empty, err error) {
	logger.Printf("email send to:%s", req.Email)
	res = new(emptypb.Empty)
	return res, nil
}
