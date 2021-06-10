package api

type Request struct {
	InitialAddress string `json:"initial_address"`
}

type Response struct {
	CorrectAddress string `json:"corrected_address"`
	FIAS           string `json:"fias"`
}
