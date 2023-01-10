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

	mux.HandleFunc("/hello", func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		//fmt.Fprintln(w, "Hello there! fmt")
		w.Write([]byte("Hello"))
	})

	log.Println("Listening...")
	handler := cors.Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

/*func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:80")
	w.Header().Set("Access-Control-Max-Age", "15")
	fmt.Println(w, "Hello, there")
	w.Write([]byte("Hello, there"))
}*/
