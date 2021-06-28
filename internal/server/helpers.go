package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/lll-phill-lll/address_correction/api"
	"github.com/lll-phill-lll/address_correction/logger"
)

func logRequest(r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		logger.Error.Println(err)
	}

	logger.Info.Println(string(requestDump))
}

func parseRequestBody(r *http.Request) (api.Request, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error.Println(err)
		return api.Request{}, errors.New("can't read body")
	}

	var request api.Request

	err = json.Unmarshal(body, &request)
	if err != nil {
		logger.Error.Println(err)
		return api.Request{}, errors.New("can't unmarshal body")
	}

	if request.InitialAddress == "" {
		err = errors.New("empty initial body")
		logger.Error.Println(err)
		return api.Request{}, err
	}

	return request, nil
}
