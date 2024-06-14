package main

import (
	"fmt"
	"log"
	"net"

	"shippingservice/handler"
	pb "shippingservice/proto"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const PORT = 50018
const ADDRESS = "127.0.0.1"

func main() {
	ipaddr := fmt.Sprintf("%s:%d", ADDRESS, PORT)

	//------------------注册到consult----------------------
	consulconfig := api.DefaultConfig()
	consulClient, err := api.NewClient(consulconfig)
	if err != nil {
		log.Fatalf("consul client create error:%v", err)
	}

	reg := api.AgentServiceRegistration{
		Tags:    []string{"shippingService"},
		Name:    "shippingService",
		Port:    PORT,
		Address: ADDRESS,
	}

	if err = consulClient.Agent().ServiceRegister(&reg); err != nil {
		log.Fatalf("consul client register error:%v", err)
	}

	//--------------------grpc service----------------------
	listen, err := net.Listen("tcp", ipaddr)
	if err != nil {
		log.Fatalf("tcp listen fatal:%v", err)
	}
	defer listen.Close()

	log.Default().Printf("tcp server start: %s", ipaddr)

	grpcService := grpc.NewServer()
	pb.RegisterShippingServiceServer(grpcService, new(handler.ShippingService))

	if err = grpcService.Serve(listen); err != nil {
		log.Fatalf("grpc service error:%v", err)
	}
}
