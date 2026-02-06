package main

import (
	"log"
	"net/http"
	"time"
	"github.com/Vivek-Patel3/httpBasic/pkg/server"
)


func main() {
	address := ":8080"
	handler := http.NewServeMux()

	srv := server.New()

	handler.HandleFunc("/", srv.HandleIndex)
	handler.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			srv.HandleUsers(w, r)
		case http.MethodPost, http.MethodPut:
			srv.HandleCreateUsers(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	s := &http.Server{
		Addr: address,
		Handler: handler,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Start server: %v\n", address)
	log.Println(s.ListenAndServe())
}
