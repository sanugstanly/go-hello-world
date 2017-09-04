package routes

import (
"fmt"
"net/http"
"github.com/gorilla/mux"
)

func Route()  {
  route := mux.NewRouter()

  err := http.ListenAndServe(":9080", r)
  fmt.Print(err)
}
