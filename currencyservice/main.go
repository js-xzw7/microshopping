package main

import (
	"fmt"
	"log"
	"net"

	"currencyservice/handler"
	pb "currencyservice/proto"

	// "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const PORT = 50012
const ADDRESS = "127.0.0.1"

func main() {
	ipadd := fmt.Sprintf("%s:%d", ADDRESS, PORT)

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

	//------------------consul 注册---------------
	// consulConfig := api.DefaultConfig()
	// consulClient, err := api.NewClient(consulConfig)
	// if err != nil {
	// 	log.Fatalf("new consul client err: %+v", err)
	// }

	// //consul代理注册
	// reg := api.AgentServiceRegistration{
	// 	Tags:    []string{"currencyservice"},
	// 	Name:    "currencyservice",
	// 	Address: ADDRESS,
	// 	Port:    PORT,
	// }

	// err = consulClient.Agent().ServiceRegister(&reg)
	// if err != nil {
	// 	log.Fatalf("consul register err: %+v", err)
	// }
}
