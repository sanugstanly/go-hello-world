package sayhi

import (
  "routes"
  "net/http"
)

func Sayhi()  {
  routes.HelloR.HandleFunc("/hi/{name}", func (w http.ResponseWriter, r *http.Request)  {
    params := routes.GetParams(r)
    w.Write([]byte("Hi "+params["name"]))
  })
}
