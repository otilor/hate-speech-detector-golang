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

	// Handle assets directory
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Handle routes
	router.HandleFunc("/", hateSpeech.Index)

	// Console message
	fmt.Println("Hate Speech Detector Application using Go")

	// Initialize server that creates an instance of http.Server
	server := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: router,
	}

	log.Fatalln(server.ListenAndServe())
}