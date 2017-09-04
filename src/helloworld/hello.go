package helloworld

import (
"net/http"
"routes"
)

func Print_Hello_World() {
  routes.HelloR.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
     payload := []byte("Hello World!")
     w.Write(payload)
  })
}
