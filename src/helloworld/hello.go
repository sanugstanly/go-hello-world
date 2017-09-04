package helloworld

import (
"fmt"
"net/http"
"github.com/gorilla/mux"
)

func Print_Hello_World() {
  route.HandleFunc("/", Test)
}


func Test(w http.ResponseWriter, r *http.Request)  {
   payload := []byte("Hello World!")
   w.Write(payload)
}
