package handler

import (
	"bytes"
	"context"
	"log"
	"os"
	pb "productcatalogservice/proto"
	"strings"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var reloadCatalog bool

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger:", log.Lshortfile)
)

// 商品分类结构体
type ProductCatalogService struct {
	sync.Mutex
	products []*pb.Product
}

// 读配置文件
func (s *ProductCatalogService) readCatalogFile() (*pb.ListProductsResponse, error) {
	s.Lock()
	defer s.Unlock()

	catalogJson, err := os.ReadFile("data/products.json")
	if err != nil {
		logger.Println("read catalog file error:", err)
		return nil, err
	}

	catalog := &pb.ListProductsResponse{}
	if err := protojson.Unmarshal(catalogJson, catalog); err != nil {
		logger.Println("unmarshal catalog file error:", err)
		return nil, err
	}

	return catalog, nil
}

// 解析配置文件
func (s *ProductCatalogService) parseCatalog() []*pb.Product {
	if reloadCatalog || len(s.products) == 0 {
		catalog, err := s.readCatalogFile()
		if err != nil {
			return []*pb.Product{}
		}

		s.products = catalog.Products
	}
	return s.products
}

// 商品列表
func (s *ProductCatalogService) ListProducts(ctx context.Context, req *emptypb.Empty) (res *pb.ListProductsResponse, err error) {
	res = new(pb.ListProductsResponse)
	res.Products = s.parseCatalog()
	return res, nil
}

// 获得单个商品
func (s *ProductCatalogService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (res *pb.Product, err error) {
	var found *pb.Product
	res = new(pb.Product)
	products := s.parseCatalog()
	for _, p := range products {
		if req.Id == p.Id {
			found = p
			break
		}
	}

	if found == nil {
		return res, status.Errorf(codes.NotFound, "no product with id %s", req.Id)
	}

	res.Id = found.Id
	res.Name = found.Name
	res.Description = found.Description
	res.PriceUsd = found.PriceUsd
	res.Picture = found.Picture
	res.Categories = found.Categories

	return res, nil
}

// 搜索商品
func (s *ProductCatalogService) SearchProducts(ctx context.Context, req *pb.SearchProductsRequest) (res *pb.SearchProductsResponse, err error) {
	var ps []*pb.Product
	res = new(pb.SearchProductsResponse)

	products := s.parseCatalog()
	for _, p := range products {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(req.Query)) || strings.Contains(strings.ToLower(p.Description), strings.ToLower(req.Query)) {
			ps = append(ps, p)
		}
	}
	res.Results = ps
	return res, nil
}

// 初始化
// func init() {
// 	sigs := make(chan os.Signal, 1)
// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
// 	go func() {
// 		for {
// 			sig := <-sigs
// 			logger.Printf("接收信号：%s\n", sig)
// 			if sig == syscall.SIGINT {
// 				reloadCatalog = true
// 				logger.Println("加载商品信息")
// 			} else {
// 				reloadCatalog = false
// 				logger.Println("不加载商品信息")
// 			}
// 		}
// 	}()
// }
