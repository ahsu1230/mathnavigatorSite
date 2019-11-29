package main
import (
  "fmt"
  "net/http"
  "strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  w.Write([]byte(message))
}

func main() {
  fmt.Println("Orion service starting...")

  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8080", nil);
  err != nil {
    panic(err)
  }
}
