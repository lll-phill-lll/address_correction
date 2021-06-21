package main

import (
	"net/http"
	"os"

	"github.com/lll-phill-lll/address_correction/internal/server"
	"github.com/lll-phill-lll/address_correction/logger"
	"github.com/lll-phill-lll/address_correction/config"
)

func main() {
	logger.SetLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

    cfg, err := config.Read()
    if err != nil {
		logger.Error.Println(err.Error())
        return
    }

	mux := server.GetMuxWithHandlers()
	logger.Info.Println("Start Listening on port", cfg.Port)
    if err := http.ListenAndServe(":" + cfg.Port, mux); err != nil {
		logger.Error.Println(err.Error())
	}
}
