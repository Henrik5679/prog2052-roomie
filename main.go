package main

import (
    "fmt"
    "log"
    "net/http"
)

func defaultPage(w http.ResponseWriter, r *http.Request) {
    log.Println("Serving a client on " + r.URL.Path)
    fmt.Fprintf(w, "<h1>Hello Firebase!</h1>")
}

func main () {
    http.HandleFunc("/", defaultPage)
    log.Println("Starting server on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
