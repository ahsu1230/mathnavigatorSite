package main
import (
  "fmt"
  "net/http"
  "orion/controllers"
)

func main() {
  fmt.Println("Orion service starting...")

  http.HandleFunc("/", controllers.SayHello)
  if err := http.ListenAndServe(":8080", nil);
  err != nil {
    panic(err)
  }
}
