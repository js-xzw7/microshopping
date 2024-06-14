package handler

import (
	"bytes"
	"context"
	"log"
	"math/rand"
	pb "recommendationservice/proto"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger:", log.Lshortfile)
)

type RecommendationService struct {
	ProductCatalogService pb.ProductCatalogServiceClient
}

func (s *RecommendationService) ListRecommendations(ctx context.Context, req *pb.ListRecommendationsRequest) (res *pb.ListRecommendationsResponse, err error) {
	maxResponseCount := 5
	res = new(pb.ListRecommendationsResponse)

	//查询商品类别
	catalog, err := s.ProductCatalogService.ListProducts(ctx, &emptypb.Empty{})
	if err != nil {
		return res, err
	}

	filteredProductsIDs := make([]string, 0, len(catalog.Products))
	for _, p := range catalog.Products {
		if contains(p.Id, req.ProductIds) {
			continue
		}
		filteredProductsIDs = append(filteredProductsIDs, p.Id)
	}

	productIDs := sample(filteredProductsIDs, maxResponseCount)
	logger.Printf("[Recv ListRecommendations] product_ids=%v", productIDs)
	res.ProductIds = productIDs
	return res, nil
}

// 判断是否包含
func contains(target string, source []string) bool {
	for _, s := range source {
		if target == s {
			return true
		}
	}
	return false
}

// 示例
func sample(source []string, c int) []string {
	n := len(source)
	if n <= c {
		return source
	}

	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}

	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		indices[i], indices[j] = indices[j], indices[i]
	}

	result := make([]string, 0, c)
	for i := 0; i < c; i++ {
		result = append(result, source[indices[i]])
	}
	return result
}
