package server

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"

    "github.com/lll-phill-lll/address_correction/api"
    "github.com/lll-phill-lll/address_correction/pkg/address"
    "github.com/lll-phill-lll/address_correction/logger"
)


func SetHandlers() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/correct", correct)
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func correct(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        fmt.Fprintf(w, "wrong method\n")
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        logger.Error.Println("Can't read body", err)
        fmt.Fprintf(w, "can't read body\n")
        return
    }

    var request api.Request

    err = json.Unmarshal(body, &request)
    if err != nil {
        logger.Error.Println("Can't unmarshal body", err)
        fmt.Fprintf(w, "can't unmarshal body\n")
        return
    }

    if request.InitialAddress == "" {
        logger.Error.Println("Can't empty address", err)
        fmt.Fprintf(w, "can't empty address\n")
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

    for name, headers := range r.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}
