//go:build ignore

package main

import (
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./index.html");
}

func main() {
    log.Print("Started application");
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
