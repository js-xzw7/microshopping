package main

import (
	"fmt"
	"log"
	"net"

	"paymentservice/handler"
	pb "paymentservice/proto"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const PORT = 50014
const ADDRESS = "127.0.0.1"

func main() {
	ipaddr := fmt.Sprintf("%s:%d", ADDRESS, PORT)

	//-------------------注册到consul----------------------
	consulConfig := api.DefaultConfig()
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("consul client create error:%v", err)
	}

	reg := api.AgentServiceRegistration{
		Tags:    []string{"paymentService"},
		Name:    "paymentsService",
		Port:    PORT,
		Address: ADDRESS,
	}

	if err = consulClient.Agent().ServiceRegister(&reg); err != nil {
		log.Fatalf("consul server register error:%v", err)
	}

	//--------------------grpc service------------------------
	listen, err := net.Listen("tcp", ipaddr)
	if err != nil {
		log.Fatalf("tcp listen error:%v", err)
	}
	defer listen.Close()
	log.Default().Printf("tcp server start: %s", ipaddr)

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, new(handler.PaymentService))

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("grpc server listen error:%v", err)
	}
}
