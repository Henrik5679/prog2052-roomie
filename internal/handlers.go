package internal

import (
    "html/template"
    "log"
    "net/http"
    "path/filepath"
    "time"
    "io/fs"
)

// A page object with some basic data while we work things out here in the beginning
var Homepage Page = Page {
    Title:    "My Super Duper Home Page",
    Article:  "Lipsum dipsum",
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

// Handler function which uses template.ParseFiles to establish the file
// system As of now, reads the whole file system, and NOT just the desired
// part, it seems.
// Called by Path: /
func DefaultHandler (w http.ResponseWriter, r *http.Request) {
    log.Println("Serving a client on " + r.URL.Path)
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    tmpl, err := template.ParseFiles(filepath.Join("presentation", "index.tmpl"))
    if err != nil {
        log.Println("Could not parse filesystem. ERROR:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    Homepage.DateTime = time.Now().Format(time.UnixDate)
    walkFilesystem(".") // Just for debuggin
    err = tmpl.Execute(w, Homepage)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

