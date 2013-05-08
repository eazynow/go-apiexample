package main

import (
    "log"
    "fmt"
    "net/http"
    "encoding/json"
    "strconv"
    "github.com/gorilla/mux"
)

type Response map[string]interface{}

type Item struct {
    Id string
    Name string
}

func (i Item) String() (s string) {
        b, err := json.Marshal(i)
        if err != nil {
                s = ""
                return
        }
        s = string(b)
        return
}

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

    // Send all incoming requests to mux.DefaultRouter.
    http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":8080", r))    
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    fmt.Fprint(w, Response{"success": true, "message": "Welcome to the root!", "method":r.Method})
}

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
