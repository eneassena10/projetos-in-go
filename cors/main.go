package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	mux.HandleFunc("/hello", handlerHello())

	handler := cors.Handler(mux)
	log.Println("listennig...")
	http.ListenAndServe(":8080", handler)
}

func handlerHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		writer.Write([]byte("Hello There!"))
	}
}
