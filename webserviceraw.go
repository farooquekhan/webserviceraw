package main

import (
	"log"
	"net/http"

	"github.com/farooquekhan/webserviceraw/webservice"
)

func main() {
	log.Println("Starting HTTP Server...")

	http.HandleFunc("/todos/", webservice.HandleRequest)

	var err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Printf("Server failed starting. Error: %s", err)
	}
}
