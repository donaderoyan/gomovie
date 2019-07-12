package main

import (
  "fmt"
  "net/http"
)


func test(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "yuhuuu")
}

func main() {
  http.HandleFunc("/test", test)
  http.ListenAndServe(":8765", nil)
}
