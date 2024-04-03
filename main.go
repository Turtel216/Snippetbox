package main 

import (
  "log"
  "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  w.Write([]byte("Hello from Snippetbox"))

}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Create a snippet..."))
}

func snippetDisplay(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Display a specific snippet"))
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet/create", snippetCreate)
  mux.HandleFunc("/snippet/view", snippetDisplay)

  log.Print("Starting server on 4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}


