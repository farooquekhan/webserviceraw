package main

import (
	"log"
	"net/http"

	"farooque.in/WebServicesRaw/webservice"
)

func main() {
	log.Println("Starting HTTP Server...")

	http.HandleFunc("/todos/", webservice.HandleRequest)

	var err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Printf("Server failed starting. Error: %s", err)
	}
}
