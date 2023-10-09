package main

import (
	"homeworks/hw7/service_a/handler"
	"log"
)

func main() {
	router := handler.NewRouter()

	log.Println("server started")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}
