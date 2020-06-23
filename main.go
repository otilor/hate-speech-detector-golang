package main

import (
	"fmt"
	"github.com/GabielFemi/hate-speech-detector-golang/hateSpeech"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", hateSpeech.Index)
	fmt.Println("Hate Speech Detector Application using Go")

	server := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: router,
	}

	log.Fatalln(server.ListenAndServe())
}