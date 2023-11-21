package main

import (
	"homeworks/hw14/internal/main_service/handler"
	"net/http"
)

func main() {
	handler.InitRouter()
	http.ListenAndServe(":8080", nil)

}
