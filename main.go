package main

import (
	"log"
	"net/http"
	"os"

	"jus.tw.cx/gcsfileserver/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s := server.Server{
		DirListPageSize: 100,
	}
	http.Handle("/", &s)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
