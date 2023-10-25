package main

import (
	"homeworks/hw11/internal/config"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.NewViperConfig()
	if err != nil {
		log.Fatal("config error: ", err)
		return
	}

	// logic (app running)

	// for showing working config http port
	log.Println("server started")
	http.ListenAndServe(cfg.HTTP.Port, nil)
}
