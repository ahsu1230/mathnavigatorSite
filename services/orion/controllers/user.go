package controllers
import (
  "net/http"
  "strings"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello with mod: " + message
  w.Write([]byte(message))
}
