package main

import (
	"homeworks/hw8/internal/handler"
	"homeworks/hw8/internal/service"
	"log"
)

func main() {
	srvs := service.New()
	hndlr := handler.New(*srvs)

	router := hndlr.InitRouter()

	log.Println("server started")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}
