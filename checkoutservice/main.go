package main

import (
	"fmt"
	"log"
	"net"

	"checkoutservice/handler"

	pb "checkoutservice/proto"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetGrpcConn(consulClient *api.Client, serviceName string, serviceTag string) *grpc.ClientConn {
	service, _, err := consulClient.Health().Service(serviceName, serviceTag, true, nil)
	if err != nil {
		log.Fatalf("获取健康服务错误：%v\n", err)
		return nil
	}

	s := service[0].Service
	address := fmt.Sprintf("%s:%d", s.Address, s.Port)

	grpcConn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("获取gprc服务连接错误:%v\n", err)
		return nil
	}

	return grpcConn
}

const PORT = 50020
const ADDRESS = "127.0.0.1"

func main() {
	ipaddr := fmt.Sprintf("%s:%d", ADDRESS, PORT)

	//------------------------注册到consul-----------------------------
	consulConfig := api.DefaultConfig()
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("consul client create err:%v", err)
	}

	reg := api.AgentServiceRegistration{
		Tags:    []string{"checkoutService"},
		Name:    "checkoutService",
		Port:    PORT,
		Address: ADDRESS,
	}

	if err = consulClient.Agent().ServiceRegister(&reg); err != nil {
		log.Fatalf("consul client register error:%v", err)
	}

	//---------------------------grpc service--------------------------
	listen, err := net.Listen("tcp", ipaddr)
	if err != nil {
		log.Fatalf("tcp listen error:%v", err)
	}
	defer listen.Close()

	log.Default().Printf("tcp server start %v \n", ipaddr)

	grpcServer := grpc.NewServer()

	checkoutservice := &handler.CheckoutService{
		CartService:           pb.NewCartServiceClient(GetGrpcConn(consulClient, "cartService", "cartService")),
		CurrencyService:       pb.NewCurrencyServiceClient(GetGrpcConn(consulClient, "currencyService", "currencyService")),
		EmailService:          pb.NewEmailServiceClient(GetGrpcConn(consulClient, "emailService", "emailService")),
		PaymentService:        pb.NewPaymentServiceClient(GetGrpcConn(consulClient, "paymentService", "paymentService")),
		ProductCatalogService: pb.NewProductCatalogServiceClient(GetGrpcConn(consulClient, "productCatalogService", "productCatalogService")),
		ShippingService:       pb.NewShippingServiceClient(GetGrpcConn(consulClient, "shippingService", "shippingService")),
	}

	pb.RegisterCheckoutServiceServer(grpcServer, checkoutservice)
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("grpc server listen error:%v", err)
	}
}
