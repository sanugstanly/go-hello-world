package helloworld

import (
"fmt"
"net/http"
"github.com/gorilla/mux"
)

func Print_Hello_World() {
  r := mux.NewRouter()
  r.HandleFunc("/", Test)
  err := http.ListenAndServe(":9080", r)
  fmt.Print(err)
}


func Test(w http.ResponseWriter, r *http.Request)  {
   payload := []byte("Hello World!")
   w.Write(payload)
}
