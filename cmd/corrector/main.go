package main

import (
	"net/http"
	"os"

	"github.com/lll-phill-lll/address_correction/internal/server"
	"github.com/lll-phill-lll/address_correction/logger"
)

func main() {
	logger.SetLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	port := ":8080"

	mux := server.GetMuxWithHandlers()
	logger.Info.Println("Start Listening on port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		logger.Error.Println(err.Error())
	}
}
