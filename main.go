package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", hateSpeech.Index)
	fmt.Println("Hate Speech Detector Application using Go")
	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: router,
	}

	log.Fatalln(server.ListenAndServe())
}