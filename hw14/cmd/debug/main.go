package main

import (
	"homeworks/hw14/internal/debug/handler"
	"net/http"
)

func main() {
	router := handler.InitRouter()

	http.ListenAndServe(":8081", router)
}
