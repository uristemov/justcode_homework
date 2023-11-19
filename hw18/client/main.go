package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"service/protobuf"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "localhost:4000", grpc.DialOption(grpc.WithInsecure()))
	if err != nil {
		log.Printf("grpc dial error %v", err)
		return
	}
	defer conn.Close()

	client := protobuf.NewDataServiceClient(conn)

	stream, err := client.SendData(ctx, new(empty.Empty))
	if err != nil {
		log.Fatalf("grpc send stream data error %v", err)
	}

	ticker := time.NewTicker(7 * time.Second)

	for {
		select {
		case <-ticker.C:
			data, err := stream.Recv()
			if err != nil {
				log.Fatalf("grpc get stream data error %v", err)
			}

			fmt.Println("Topic got from grpc: ", data.Topic)
			fmt.Println("Data got from grpc: ", data.Data)
		case <-ctx.Done():
			log.Println("context is done by timeout")
			break
		}

	}
}
