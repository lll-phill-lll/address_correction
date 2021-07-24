package api

type Request struct {
	InitialAddress string `json:"initial_address"`
	City           string `json:"city"`
}

type CorrectAddress struct {
	City            string `json:"city"`
	StreetType      string `json:"street_type"`
	StreetName      string `json:"street_name"`
	HouseNumber     string `json:"house_number"`
	Korpus          string `json:"korpus"`
	CompoundAddress string `json:"compound_address"`
}

type Response struct {
	CorrectedAddress CorrectAddress `json:"corrected_address"`
	FIAS             string         `json:"fias"`
}
