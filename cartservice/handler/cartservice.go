package handler

import (
	"cartservice/cartstore"
	pb "cartservice/proto"
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// 购物车服务结构体
type CartService struct {
	Store cartstore.CartStore
}

// 添加商品
func (s *CartService) AddItem(ctx context.Context, req *pb.AddItemRequest) (res *emptypb.Empty, err error) {
	res = new(emptypb.Empty)
	return s.Store.AddItem(ctx, req.UserId, req.Item.ProductId, req.Item.Quantity, res)
}

// 获得购物车
func (s *CartService) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.Cart, error) {
	cart, err := s.Store.GetCart(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	res := new(pb.Cart)
	res.Items = cart.Items
	res.UserId = cart.UserId
	return res, nil
}

// 清空购物车
func (s *CartService) EmptyCart(ctx context.Context, req *pb.EmptyCartRequest) (*emptypb.Empty, error) {
	return s.Store.EmptyCart(ctx, req.UserId)
}
