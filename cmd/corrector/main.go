package main

import (
    "os"
    "net/http"

    "github.com/lll-phill-lll/address_correction/logger"
    "github.com/lll-phill-lll/address_correction/internal/server"
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
