// Package main is an example of how to create an API server using
// gorilla-mux. It creates and listens to some end points
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Response is an open map to return key value pairs back to the user
type Response map[string]interface{}

// Item is a sample type that may be requested
type Item struct {
	Id   string
	Name string
}

// String representation of the response map
func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

func main() {

	// Register a couple of routes.
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/items", itemsHandler)
	r.HandleFunc("/items/{id}", itemHandler)

	log.Println("Starting up a http server on port 8080...")
	log.Println("Listening to /, /items and /items/{id}")

	// Send all incoming requests to mux.DefaultRouter.
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))

}

// Handles requests to /
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, Response{"success": true, "message": "Welcome to the root!", "method": r.Method})
}

// Handles requests to /items

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var items [5]Item
	for i := 0; i < 5; i++ {
		itemdesc := fmt.Sprintf("dummy item %d", i)
		items[i] = Item{strconv.Itoa(i), itemdesc}
	}

	b, err := json.Marshal(items)

	resp := ""
	if err == nil {
		resp = string(b)
	}
	fmt.Fprint(w, resp)
}

// Handles requests to /items/{id}
func itemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	itemdesc := fmt.Sprintf("dummy item %s", vars["id"])
	i := Item{vars["id"], "dummy item " + itemdesc}
	b, err := json.Marshal(i)

	resp := ""
	if err == nil {
		resp = string(b)
	}
	fmt.Fprint(w, resp)
}
