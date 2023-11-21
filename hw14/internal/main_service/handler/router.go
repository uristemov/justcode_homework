package handler

import "net/http"

func InitRouter() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Main Server"))
	})
}
