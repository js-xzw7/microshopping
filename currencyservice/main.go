package main

import (
	"fmt"
	"log"
	"net"

	"currencyservice/handler"
	pb "currencyservice/proto"

	capi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const PORT = 50012
const ADDRESS = "127.0.0.1"

func main() {
	ipadd := fmt.Sprintf("%s:%d", ADDRESS, PORT)

	//------------------consul 注册---------------
	consulConfig := capi.DefaultConfig()
	consulClient, err := capi.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("new consul client err: %+v", err)
	}

	//consul代理注册
	reg := capi.AgentServiceRegistration{
		Tags:    []string{"currencyService"},
		Name:    "currencyService",
		Address: ADDRESS,
		Port:    PORT,
	}

	err = consulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatalf("consul register err: %+v", err)
	}

	//---------------grpc 服务-----------------
	listen, err := net.Listen("tcp", ipadd)
	if err != nil {
		log.Fatalf("tcp listen err: %+v", err)
	}
	defer listen.Close()
	log.Default().Printf("tcp server statrt: %s", ipadd)

	grpcServer := grpc.NewServer()

	pb.RegisterCurrencyServiceServer(grpcServer, new(handler.CurrencyService))

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("currencyservice set listen err: %+v", err)
	}

}
