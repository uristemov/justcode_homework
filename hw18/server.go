package main

import (
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"service/protobuf"
	"time"
)

type Server struct {
	protobuf.UnimplementedDataServiceServer
}

func (s *Server) SendData(req *emptypb.Empty, stream protobuf.DataService_SendDataServer) error {
	timer := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-stream.Context().Done():
			log.Println("stream is done!")
		case <-timer.C:
			err := stream.Send(&protobuf.Data{Topic: "Topic", Data: "Data of the message"})
			if err != nil {
				log.Printf(err.Error())
				return err
			}
		}
	}
}
