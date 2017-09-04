package routes

import(
  "fmt"
  "net/http"
)

func Serve(port string)  {
  err := http.ListenAndServe(":"+port, HelloR)
  fmt.Print(err)
}
