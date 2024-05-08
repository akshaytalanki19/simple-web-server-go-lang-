package main

import (
  "fmt"
  "log"
  "net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
  if err := r.ParseForm(); err != nil { // Use r.ParseForm()
    fmt.Fprintf(w, "ParseForm() err: %v", err) // Use fmt.Fprintf with correct format specifier
    return
  }

  fmt.Fprintf(w, "Post request successful!\n")

  name := r.FormValue("name")
  address := r.FormValue("Address") // Use r.FormValue with correct capitalization

  fmt.Fprintf(w, "Name: %s\n", name)
  fmt.Fprintf(w, "Address: %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/hello" {
    http.Error(w, "404 not found", http.StatusNotFound) // Use http.StatusNotFound constant
    return
  }
  if r.Method != http.MethodGet {
    http.Error(w, "method is not supported", http.StatusNotFound) // Use http.StatusNotFound constant
    return
  }

  fmt.Fprintf(w, "hello!\n")
}

func main() {
  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileServer)
  http.HandleFunc("/form", formHandler)
  http.HandleFunc("/hello", helloHandler)

  fmt.Printf("starting server at port 8080\n")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}
