package main

import (
	"fmt"
	"log"
	"net"

	handler "cartservice/handler"
	pb "cartservice/proto"

	capi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const PORT = 50011
const ADDRESS = "127.0.0.1"

func main() {
	ipaddr := fmt.Sprintf("%s:%d", ADDRESS, PORT)

	//---------------注册consul-----------------
	consulConfig := capi.DefaultConfig()
	consulClent, err := capi.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("new consul client err: %+v", err)
	}

	reg := capi.AgentServiceRegistration{
		Tags:    []string{"cartservice"},
		Name:    "cartservice",
		Address: ADDRESS,
		Port:    PORT,
	}

	if err = consulClent.Agent().ServiceRegister(&reg); err != nil {
		log.Fatalf("consul register err: %+v", err)
	}

	//---------------grpc 服务-----------------
	listen, err := net.Listen("tcp", ipaddr)
	if err != nil {
		log.Fatalf("listen err: %+v", err)
	}
	defer listen.Close()

	log.Default().Printf("tcp server statrt: %s", ipaddr)

	grpcService := grpc.NewServer()
	pb.RegisterCartServiceServer(grpcService, new(handler.CartService))

	if err := grpcService.Serve(listen); err != nil {
		log.Fatalf("cartservice grpc service start err: %+v", err)
	}
}
