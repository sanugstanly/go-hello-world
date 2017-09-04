package routes

import (
"fmt"
"net/http"
"github.com/gorilla/mux"
)

func Route()  {
  r := mux.NewRouter()
  err := http.ListenAndServe(":9080", r)
  fmt.Print(err)
}
