package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	StartDispatcher(5)

	r := mux.NewRouter()
	r.HandleFunc("/submit", SubmitJobHandler).Methods("POST")
	r.HandleFunc("/status/{id}", GetJobStatusHandler).Methods("GET")

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
