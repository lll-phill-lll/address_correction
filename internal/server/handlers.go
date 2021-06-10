package server

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/lll-phill-lll/address_correction/api"
    "github.com/lll-phill-lll/address_correction/pkg/address"
    "github.com/lll-phill-lll/address_correction/logger"
)


func SetHandlers() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/correct", correct)
}

func hello(w http.ResponseWriter, r *http.Request) {
    logRequest(r)
    fmt.Fprintf(w, "hello\n")
}

func correct(w http.ResponseWriter, r *http.Request) {
    logRequest(r)
    if r.Method != "POST" {
        fmt.Fprintf(w, "wrong method\n")
    }

    request, err := parseRequestBody(r)
    if err != nil {
        fmt.Fprintf(w, err.Error() + "\n")
        return
    }


    w.Header().Set("Content-Type", "application/json")

    response := api.Response{CorrectAddress: address.Correct(request.InitialAddress), FIAS: "123"}

    err = json.NewEncoder(w).Encode(response)
    if err != nil {
        logger.Error.Println("Can't marshal response", err)
        fmt.Fprintf(w, "can't marshal response\n")
        return
    }
}
