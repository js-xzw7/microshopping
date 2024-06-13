package cartstore

import (
	pb "cartservice/proto"
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// 购物车接口
type CartStore interface {
	AddItem(ctx context.Context, userID, productID string, quantity int32, out *emptypb.Empty) (r *emptypb.Empty, err error)
	GetCart(ctx context.Context, userID string) (*pb.Cart, error)
	EmptyCart(ctx context.Context, userID string) (*emptypb.Empty, error)
}

// 实例化
func NewMemoryCartStore() CartStore {
	return &memoryCartStore{
		carts: make(map[string]map[string]int32),
	}
}
