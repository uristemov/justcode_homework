package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"service/protobuf"
)

func main() {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Printf("listen tcp error %v", err)
	}

	grpcServer := grpc.NewServer()

	s := &Server{}

	protobuf.RegisterDataServiceServer(grpcServer, s)

	log.Printf("starting grpc server at 4000 port")
	log.Fatalf(grpcServer.Serve(lis).Error())
}
