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
func TemplateHandler (w http.ResponseWriter, r *http.Request) {
    log.Println("Serving a client on " + r.URL.Path)

    walkFilesystem(".")

    tmpl, err := template.ParseFiles("presentation/index.tmpl")
    if err != nil { panic(err) }
    err = tmpl.Execute(w, Homepage)
    if err != nil { panic(err) }
}


// Handler function which uses template.ParseFiles to establish the file
// system As of now, reads the whole file system, and NOT just the desired
// part, it seems.
// Called by Path: /pic
// Relies on http.FileServer
func DefaultHandler (w http.ResponseWriter, r *http.Request) {
    log.Println("Serving a client on " + r.URL.Path)
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    tmpl, err := template.ParseFiles(filepath.Join("presentation", "index.tmpl")) // TODO it *should* just be connected to that "presentation" is root, but that's not the case...
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

