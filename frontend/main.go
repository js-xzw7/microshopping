package main

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "frontend/proto"
)

const (
	name    = "frontend"
	version = "1.0.0"

	defaultCurrency = "USD"
	cookieMaxAge    = 60 * 60 * 48

	cookiePrefix    = "shop_"
	cookieSessionID = cookiePrefix + "session-id"
	cookieCurrency  = cookiePrefix + "currency"
)

var (
	whitelistedCurrencies = map[string]bool{
		"USD": true,
		"EUR": true,
		"CAD": true,
		"JPY": true,
		"GBP": true,
		"TRY": true,
	}
)

// type ctxKeySessionID struct{}

// 前端server
type FrontendServer struct {
	adService             pb.AdServiceClient
	cartService           pb.CartServiceClient
	checkoutService       pb.CheckoutServiceClient
	currencyService       pb.CurrencyServiceClient
	productCatalogService pb.ProductCatalogServiceClient
	recommendationService pb.RecommendationServiceClient
	shippingService       pb.ShippingServiceClient
}

// 获得grpc连接
func GetGrpcConn(consulClient *api.Client, serviceName string, serviceTag string) *grpc.ClientConn {
	service, _, err_service := consulClient.Health().Service(serviceName, serviceTag, true, nil)
	if err_service != nil {
		fmt.Println("获取健康服务报错：", err_service)
		return nil
	}
	s := service[0].Service
	address := s.Address + ":" + strconv.Itoa(s.Port)
	//链接grpc服务
	grpcConn, _ := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	return grpcConn
}

func main() {

	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout

	//1.初始化consul配置
	consulConfig := api.DefaultConfig()
	//2.创建consul对象
	consulClient, err_consul := api.NewClient(consulConfig)
	if err_consul != nil {
		fmt.Println("consul创建对象报错：", err_consul)
		return
	}

	svc := FrontendServer{
		adService:             pb.NewAdServiceClient(GetGrpcConn(consulClient, "adService", "adService")),
		cartService:           pb.NewCartServiceClient(GetGrpcConn(consulClient, "cartService", "cartService")),
		checkoutService:       pb.NewCheckoutServiceClient(GetGrpcConn(consulClient, "checkoutService", "checkoutService")),
		currencyService:       pb.NewCurrencyServiceClient(GetGrpcConn(consulClient, "currencyService", "currencyService")),
		productCatalogService: pb.NewProductCatalogServiceClient(GetGrpcConn(consulClient, "productcatalogService", "productcatalogService")),
		recommendationService: pb.NewRecommendationServiceClient(GetGrpcConn(consulClient, "recommendationService", "recommendationService")),
		shippingService:       pb.NewShippingServiceClient(GetGrpcConn(consulClient, "shippingService", "shippingService")),
	}

	// recomRes, err := svc.recommendationService.ListRecommendations(context.TODO(), &pb.ListRecommendationsRequest{
	// 	UserId:     "1",
	// 	ProductIds: []string{"1"},
	// })

	// if err != nil {
	// 	fmt.Println("获取推荐商品报错：", err)
	// }

	// fmt.Printf("获取推荐商品：%v\n", recomRes)

	r := gin.Default()

	r.FuncMap = template.FuncMap{
		"renderMoney":        renderMoney,
		"renderCurrencyLogo": renderCurrencyLogo,
	}

	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")

	r.Use(setCookie)
	// 首页
	r.GET("/", svc.HomeHandler)
	// 商品
	r.GET("/product/:id", svc.ProductHandler)
	// 查看购物车
	r.GET("/cart", svc.viewCartHandler)
	// 添加购物车
	r.POST("/cart", svc.addToCartHandler)
	// 清空购物车
	r.POST("/cart/empty", svc.emptyCartHandler)
	// 设置货币种类
	r.POST("/setCurrency", svc.setCurrencyHandler)
	// // 退出登录
	r.GET("/logout", svc.logoutHandler)
	// // 结账
	r.POST("/cart/checkout", svc.placeOrderHandler)

	if err := r.Run(":8052"); err != nil {
		log.Fatalf("gin启动失败: %v", err)
	}
}
