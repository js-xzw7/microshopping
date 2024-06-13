package cartstore

import (
	"context"
	"sync"

	pb "cartservice/proto"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// 数据保存在内存中的结构体，使用嵌套map保存
type memoryCartStore struct {
	sync.RWMutex
	carts map[string]map[string]int32
}

// 添加商品
func (m *memoryCartStore) AddItem(ctx context.Context, userId, productId string, quantity int32, req *emptypb.Empty) (res *emptypb.Empty, err error) {
	m.Lock()
	defer m.Unlock()

	if cart, ok := m.carts[userId]; ok {
		if currentQuantity, ok := cart[productId]; ok {
			cart[productId] = currentQuantity + quantity
		} else {
			cart[productId] = quantity
		}
	} else {
		m.carts[userId] = map[string]int32{productId: quantity}
	}

	return res, nil
}

// 清空购物车
func (m *memoryCartStore) EmptyCart(ctx context.Context, userId string) (res *emptypb.Empty, err error) {
	m.Lock()
	defer m.Unlock()

	delete(m.carts, userId)

	return res, nil
}

// 获取购物车
func (m *memoryCartStore) GetCart(ctx context.Context, userId string) (*pb.Cart, error) {
	m.RLock()
	defer m.RUnlock()

	if cart, ok := m.carts[userId]; ok {
		item := make([]*pb.CartItem, 0, len(cart))
		for productId, quantity := range cart {
			item = append(item, &pb.CartItem{ProductId: productId, Quantity: quantity})
		}

		return &pb.Cart{UserId: userId, Items: item}, nil
	}

	return &pb.Cart{UserId: userId}, nil
}
