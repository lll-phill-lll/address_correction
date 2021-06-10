package server

import (
    "net/http"
)


func checkRequestHandler(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    logRequest(r)

    if r.Method != "POST" {
        http.Error(w, "Wrong method", http.StatusBadRequest)
        return
    }

    next.ServeHTTP(w, r)
  })
}
