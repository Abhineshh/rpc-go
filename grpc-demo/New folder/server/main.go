package main

import (
	"log"
	"net"
	dangit "grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct{
	dangit.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp",port)
	if err != nil {
		log.Fatalf("Failed to start the server %v",err)
	}
	grpcServer := grpc.NewServer()

	dangit.RegisterGreetServiceServer(grpcServer,&helloServer{})

	 err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Failed to start the server %v",err)
	}
}