package routes

import (
"github.com/gorilla/mux"
"net/http"
)

var HelloR =  mux.NewRouter()

func GetParams(r *http.Request) map[string]string {
  par := mux.Vars(r)
  return par
}
