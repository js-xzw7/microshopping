package main

import (
	"fmt"
	"log"
	"net"

	"recommendationservice/handler"
	pb "recommendationservice/proto"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const PORT = 50016
const ADDRESS = "127.0.0.1"

func GetGrpcConn(consulClient *api.Client, serviceName string, serviceTag string) *grpc.ClientConn {
	service, _, err := consulClient.Health().Service(serviceName, serviceTag, true, nil)
	if err != nil {
		log.Fatalf("获取健康服务错误：%v", err)
	}

	s := service[0].Service
	address := fmt.Sprintf("%s:%d", s.Address, s.Port)
	log.Default().Printf("servcie name: %v \n", serviceName)
	log.Default().Printf("address: %v\n", address)

	//连接grpc服务
	grpcConn, _ := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpcConn
}

func main() {
	ipaddr := fmt.Sprintf("%s:%d", ADDRESS, PORT)

	//---------------------注册到consul------------------
	consulConfig := api.DefaultConfig()
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("consul client error:%v", err)
	}

	reg := api.AgentServiceRegistration{
		Tags:    []string{"recommendationService"},
		Name:    "recommendationService",
		Port:    PORT,
		Address: ADDRESS,
	}

	if err = consulClient.Agent().ServiceRegister(&reg); err != nil {
		log.Fatalf("consul register error:%v", err)
	}

	//--------------------grpc service----------------------
	listen, err := net.Listen("tcp", ipaddr)
	if err != nil {
		log.Fatalf("tcp listen error:%v", err)
	}
	defer listen.Close()

	grpcService := grpc.NewServer()
	recommendationservice := &handler.RecommendationService{
		ProductCatalogService: pb.NewProductCatalogServiceClient(GetGrpcConn(consulClient, "productcatalogService", "productcatalogService")),
	}

	pb.RegisterRecommendationserviceServer(grpcService, recommendationservice)

	if err = grpcService.Serve(listen); err != nil {
		log.Fatalf("grpc server error:%v", err)
	}
}
