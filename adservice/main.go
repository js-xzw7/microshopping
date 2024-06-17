package main

import (
	"fmt"
	"log"
	"net"

	"adservice/handler"
	pb "adservice/proto"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const PROT = 50010
const ADDRESS = "127.0.0.1"

func main() {
	ipaddr := fmt.Sprintf("%s:%d", ADDRESS, PROT)

	//---------------注册consul-------------------
	consulConfig := api.DefaultConfig()
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("consul client create fatal:%v", err)
	}
	reg := api.AgentServiceRegistration{
		Tags:    []string{"adService"},
		Name:    "adService",
		Address: ADDRESS,
		Port:    PROT,
	}

	if err = consulClient.Agent().ServiceRegister(&reg); err != nil {
		log.Fatalf("consul service register err: %v", err)
	}

	//---------------grpc service--------------------
	listen, err := net.Listen("tcp", ipaddr)
	if err != nil {
		log.Fatalf("tcp listen err:%v", err)
	}
	defer listen.Close()
	log.Default().Printf("tcp server statrt: %s", ipaddr)

	grpcService := grpc.NewServer()
	pb.RegisterAdServiceServer(grpcService, new(handler.Adservice))

	if err = grpcService.Serve(listen); err != nil {
		log.Fatalf("grpc service start err:%v", err)
	}
}
