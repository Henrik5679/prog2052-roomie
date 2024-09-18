package main

import (
    "errors"
    "log"
    "net/http"
    "os"
    "roomie/internal"
)

func main() {
    // --- Establish Handler Functions ---
    // TODO: Every page has proper header initialization and such
    http.HandleFunc(internal.PATH_DEFAULT,   internal.DefaultHandler)
    http.HandleFunc(internal.PATH_FORM,      internal.FormHandler)
    http.HandleFunc(internal.PATH_USERINPUT, internal.UserInputHandler)

    // --- Parse File System ---
    // TODO Got this part of the code from perplexity.ai
    // I'd like to investigate for how it works
    fs := http.FileServer(http.Dir("presentation"))
    http.Handle("/presentation/", http.StripPrefix("/presentation/", fs))

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

