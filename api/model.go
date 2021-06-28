package api

type Request struct {
	InitialAddress string `json:"initial_address"`
	City           string `json:"city"`
}

type Response struct {
	CorrectAddress string `json:"corrected_address"`
	FIAS           string `json:"fias"`
}
