package webservice

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"farooque.in/WebServicesRaw/storage"
	"farooque.in/WebServicesRaw/webutils"
)

var funcTable map[string]func(w http.ResponseWriter, r *http.Request)

func init() {

	funcTable = make(map[string]func(w http.ResponseWriter, r *http.Request))

	funcTable[http.MethodGet] = handleList
	funcTable[http.MethodPost] = handleAddUpdate
	funcTable[http.MethodDelete] = handleRemove
}

func handleList(w http.ResponseWriter, r *http.Request) {
	webutils.HandleSuccess(&w, storage.Get())
}

func handleAddUpdate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	byteData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		webutils.HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
		return
	}

	var item storage.TodoItem

	err = json.Unmarshal(byteData, &item)

	if err != nil {
		webutils.HandleError(&w, 500, "Internal Server Error", "Error unmarhsalling JSON", err)
		return
	}

	if item.Title == "" {
		webutils.HandleError(&w, 400, "Bad Request", "Unmarshalled JSON didn't have required fields", nil)
		return
	}

	id := storage.AddUpdate(item)

	log.Println("Added/Updated item:", item)

	webutils.HandleSuccess(&w, id)
}

func handleRemove(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		webutils.HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
		return
	}

	var id int

	err = json.Unmarshal(requestBody, &id)

	if err != nil {
		webutils.HandleError(&w, 400, "Bad Request", "Error unmarshalling", err)
		return
	}

	if id == 0 {
		webutils.HandleError(&w, 500, "Bad Request", "ID not provided", nil)
		return
	}

	if storage.Remove(id) {
		webutils.HandleSuccess(&w, id)
	} else {
		webutils.HandleError(&w, 400, "Bad Request", "ID not found", nil)
	}
}

// HandleRequest implements a simple request router
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r.Method)

	f, ok := funcTable[r.Method]
	if !ok || f == nil {
		webutils.HandleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		return
	}

	f(w, r)
}
