package main

import (
      "fmt"
      "log"
      "net/http"
)
var version string

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
  fmt.Fprinln(w,version)
}

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
