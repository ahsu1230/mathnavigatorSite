package main
import (
  "fmt"
  "net/http"
  "github.com/ahsu1230/mathnavigatorSite/tree/master/services/orion/controllers"
)

func main() {
  fmt.Println("Orion service starting...")

  http.HandleFunc("/", controllers.sayHello)
  if err := http.ListenAndServe(":8080", nil);
  err != nil {
    panic(err)
  }
}
