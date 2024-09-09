package main

import (
    "errors"
    _ "fmt"
    "html/template"
    "io/fs"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

// Page; a struct with which we implement HTML pages.
type Page struct { // TODO move to struct file
    Title    string
    Article  string
    DateTime string // This means when program starts
}

var Homepage Page = Page {
    Title:    "My Super Duper Home Page",
    Article:  "Lipsum dipsum",
    DateTime: time.Now().Format(time.UnixDate),
}

// walkFilesystem; a basic development function which prints the filesystem
// as the program sees it
func walkFilesystem (path string) {
    err := filepath.WalkDir (path, func(innerpath string, d fs.DirEntry, err error) error {
        if err != nil { return err }
        log.Println(innerpath)
        return nil
    })
    if err != nil {
        log.Fatalf("Error walking the path %q: %v\n", ".", err)
    }
}

// Handler which utilizes template.Parsefiles
// Parses entire filesystem, which is bad and wrong
// Called by path: /template
func templatePage (w http.ResponseWriter, r *http.Request) {
    log.Println("Serving a client on " + r.URL.Path)

    walkFilesystem(".")

    tmpl, err := template.ParseFiles("Frontend/index.tmpl")
    if err != nil { panic(err) }
    err = tmpl.Execute(w, Homepage)
    if err != nil { panic(err) }
}


// Handler function which uses template.ParseFiles to establish the file
// system As of now, reads the whole file system, and NOT just the desired
// part, it seems.
// Called by Path: /pic
// Relies on http.FileServer
func defaultHandler (w http.ResponseWriter, r *http.Request) {
    log.Println("Serving a client on " + r.URL.Path)
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    tmpl, err := template.ParseFiles(filepath.Join("Frontend", "index.tmpl")) // TODO it *should* just be connected to that "Frontend" is root, but that's not the case...
    if err != nil {
        log.Println("Could not parse filesystem. ERROR:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // walkFilesystem(".")
    err = tmpl.Execute(w, Homepage)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    // --- Establish Handler Functions ---
    // TODO: have the PATHs be constants
    http.HandleFunc("/", defaultHandler)

    // --- Parse File System ---
    // TODO Got this part of the code from perplexity.ai
    // I'd like to investigate for how it works
    fs := http.FileServer(http.Dir("Frontend"))
    http.Handle("/Frontend/", http.StripPrefix("/Frontend/", fs))

    // --- Start Server ---
    log.Println("Starting server on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if errors.Is(err, http.ErrServerClosed) {
        log.Printf("Server closed; all is well. I don't know in what situations you may see this message.")
    } else if err != nil {
        log.Printf("Error starting server: %s", err)
        os.Exit(1)
    }
}

