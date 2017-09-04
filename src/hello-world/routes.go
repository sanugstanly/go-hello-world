package helloworld

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


func Print_Hello_World(w http.ResponseWriter, r *http.Request) {
   payload := []byte("Hello World!")
   w.Write(payload)
}
