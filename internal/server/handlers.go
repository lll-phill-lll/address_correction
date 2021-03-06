package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/lll-phill-lll/address_correction/api"
	"github.com/lll-phill-lll/address_correction/logger"
	"github.com/lll-phill-lll/address_correction/pkg/address"
)

func GetMuxWithHandlers() *http.ServeMux {
	mux := http.NewServeMux()

	correctHandler := http.HandlerFunc(correct)
	mux.Handle("/correct", checkRequestHandler(correctHandler))

	helloHandler := http.HandlerFunc(hello)
	mux.Handle("/hello", helloHandler)

	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func correct(w http.ResponseWriter, r *http.Request) {
	request, err := parseRequestBody(r)
	if err != nil {
		fmt.Fprintf(w, err.Error()+"\n")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	initialAddress, err := url.QueryUnescape(request.InitialAddress)
	if err != nil {
		fmt.Fprintf(w, err.Error()+"\n")
		return
	}

	city, err := url.QueryUnescape(request.City)
	if err != nil {
		fmt.Fprintf(w, err.Error()+"\n")
		return
	}

	fias, correctedAddress := address.CorrectAndGetFIAS(initialAddress, city)

	response := api.Response{CorrectedAddress: correctedAddress, FIAS: fias}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Error.Println("Can't marshal response", err)
		fmt.Fprintf(w, "can't marshal response\n")
		return
	}
}
