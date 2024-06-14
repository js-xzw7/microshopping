package main

import (
	"fmt"
	"log"
	"net"

	"emailservice/handler"
	pb "emailservice/proto"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const PORT = 50013
const ADDRESS = "127.0.0.1"

func main() {
	ipaddr := fmt.Sprintf("%s:%v", ADDRESS, PORT)

	//--------------注册到consul------------------
	consulConfig := api.DefaultConfig()
	consulCleint, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("consult client create fatal:%v", err)
	}

	reg := api.AgentServiceRegistration{
		Tags:    []string{"emailService"},
		Name:    "emailService",
		Port:    PORT,
		Address: ADDRESS,
	}

	if err = consulCleint.Agent().ServiceRegister(&reg); err != nil {
		log.Fatalf("consult register error:%v", err)
	}

	//------------------gprc service-----------------------------
	listen, err := net.Listen("tcp", ipaddr)
	if err != nil {
		log.Fatalf("tcp server listen error:%v", err)
	}
	defer listen.Close()

	log.Default().Printf("tcp server start: %s", ipaddr)

	grpcServer := grpc.NewServer()
	pb.RegisterEmailServiceServer(grpcServer, new(handler.DummyEmailService))

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("grpc server listen error:%v", err)
	}

}
