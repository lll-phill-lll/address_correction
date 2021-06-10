package main

import (
    "os"
    "net/http"

    "github.com/lll-phill-lll/address_correction/logger"
    "github.com/lll-phill-lll/address_correction/internal/server"
)

func main() {
    logger.SetLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
    logger.Info.Println("Hello, world")

    port := ":8080"

    server.SetHandlers()
	logger.Info.Println("Start Listening on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		logger.Error.Println(err.Error())
	}
}
