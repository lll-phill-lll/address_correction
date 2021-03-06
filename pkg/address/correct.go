package address

import (
	"github.com/lll-phill-lll/address_correction/api"
	"github.com/lll-phill-lll/address_correction/internal/fiasdata"
	"github.com/lll-phill-lll/address_correction/logger"
	parser "github.com/openvenues/gopostal/parser"
	"regexp"
	"strings"
)

// Possible keys in response:
//     house: venue name e.g. "Brooklyn Academy of Music", and building names e.g. "Empire State Building"
//     category: for category queries like "restaurants", etc.
//     near: phrases like "in", "near", etc. used after a category phrase to help with parsing queries like "restaurants in Brooklyn"
//     house_number: usually refers to the external (street-facing) building number. In some countries this may be a compount, hyphenated number which also includes an apartment number, or a block number (a la Japan), but libpostal will just call it the house_number for simplicity.
//     road: street name(s)
//     unit: an apartment, unit, office, lot, or other secondary unit designator
//     level: expressions indicating a floor number e.g. "3rd Floor", "Ground Floor", etc.
//     staircase: numbered/lettered staircase
//     entrance: numbered/lettered entrance
//     po_box: post office box: typically found in non-physical (mail-only) addresses
//     postcode: postal codes used for mail sorting
//     suburb: usually an unofficial neighborhood name like "Harlem", "South Bronx", or "Crown Heights"
//     city_district: these are usually boroughs or districts within a city that serve some official purpose e.g. "Brooklyn" or "Hackney" or "Bratislava IV"
//     city: any human settlement including cities, towns, villages, hamlets, localities, etc.
//     island: named islands e.g. "Maui"
//     state_district: usually a second-level administrative division or county.
//     state: a first-level administrative division. Scotland, Northern Ireland, Wales, and England in the UK are mapped to "state" as well (convention used in OSM, GeoPlanet, etc.)
//     country_region: informal subdivision of a country without any political status
//     country: sovereign nations and their dependent territories, anything with an ISO-3166 code.
//     world_region: currently only used for appending ???West Indies??? after the country name, a pattern frequently used in the English-speaking Caribbean e.g. ???Jamaica, West Indies???
func CorrectAndGetFIAS(address string, city string) (string, api.CorrectAddress) {
	// TODO move to preprocess function
	address = strings.ReplaceAll(address, "(", " ")
	address = strings.ReplaceAll(address, ")", " ")
	address = strings.ReplaceAll(address, "???", " ")
	address = strings.ReplaceAll(address, " ?????? ", " ???????????? ")
	address = strings.ReplaceAll(address, " ??????.", " ???????????? ")
	address = strings.ReplaceAll(address, " ?????? ", " ?????????????? ")

	space := regexp.MustCompile(`\s+`)
	address = space.ReplaceAllString(address, " ")

	parsed_values := parser.ParseAddress(address)
	logger.Info.Println("Parsed", address, "into", parsed_values)

	street_name, street_type := "", ""
	house_number := ""
	road := ""
	city_with_error := ""

	for _, parsed_value := range parsed_values {
		if parsed_value.Label == "road" {
			road = parsed_value.Value
			logger.Info.Println("Found road:", road)
		}
		if parsed_value.Label == "house_number" {
			house_number = parsed_value.Value
			logger.Info.Println("Found house_number:", house_number)
		}
		if parsed_value.Label == "city" {
			if city == "" || city == "ANY" {
				city = parsed_value.Value
			}
		}
		if parsed_value.Label == "house" {
			city_with_error = parsed_value.Value
		}
	}

	city = getCityName(city)
	if city == "" {
		city = getCityName(city_with_error)
	}
	if city != "ANY" {
		city = strings.ToLower(city)
	}

	if city == "" {
		city = "ANY"
	}

	if road != "" {
		logger.Info.Println("Found road:", road)
		street_name, street_type = SpiltRoadIntoNameAndType(road)
		logger.Info.Println("Split results:", street_name, street_type)
		street_type = StreetTypeToCanonical(street_type)
		if street_type == "" {
			street_type = "ANY"
		}
	}

	house_number, korpus := SplitHouseNumberToNumberAndKorpus(house_number)
	if korpus == "" {
		korpus = "ANY"
	}

	foundAddress := fiasdata.Storage.GetAddress(city, street_type, street_name, house_number, korpus)
	if foundAddress.FIAS != "" {
		compoundAddress := makeCompoundAddress(foundAddress.City, foundAddress.StreetType, foundAddress.FormalName, foundAddress.HouseNum, foundAddress.Korpus)
		correctedAddress := api.CorrectAddress{foundAddress.City, foundAddress.StreetType, foundAddress.FormalName, foundAddress.HouseNum, foundAddress.Korpus, compoundAddress}
		logger.Info.Println(correctedAddress)
		return foundAddress.FIAS, correctedAddress

	}

	if city == "ANY" {
		city = ""
	}
	if street_type == "ANY" {
		street_type = ""
	}
	if street_name == "ANY" {
		street_name = ""
	}
	if house_number == "ANY" {
		house_number = ""
	}
	if korpus == "ANY" {
		korpus = ""
	}

	compoundAddress := makeCompoundAddress(city, street_type, street_name, house_number, korpus)
	correctedAddress := api.CorrectAddress{city, street_type, street_name, house_number, korpus, compoundAddress}
	logger.Info.Println(correctedAddress)

	return "Not Found", correctedAddress
}

func SpiltRoadIntoNameAndType(road string) (string, string) {
	splittedRoad := strings.FieldsFunc(road, Split)
	logger.Info.Println("Splitted:", road, "to", splittedRoad)

	elementsNum := len(splittedRoad)
	logger.Info.Println("elements:", elementsNum)

	if len(splittedRoad) == 1 {
		// TODO consider returning ANY
		return road, ""
	}

	if IsStreetType(splittedRoad[elementsNum-1]) {
		return strings.Join(splittedRoad[:elementsNum-1], " "), splittedRoad[elementsNum-1]
	} else if IsStreetType(splittedRoad[0]) {
		return strings.Join(splittedRoad[1:], " "), splittedRoad[0]
	}

	return road, ""
}

func Split(r rune) bool {
	return r == ' ' || r == ','
}

func IsStreetType(street_type string) bool {
	return StreetTypeToCanonical(street_type) != ""
}

func StreetTypeToCanonical(street_type string) string {
	if street_type == "" {
		return street_type
	}

	street_type = strings.ReplaceAll(street_type, ".", "")
	// TODO: sort and find
	canonical_names := []string{"????", "??????", "????", "??????", "????????????", "??????-??", "??????????", "????-????", "??-??", "??????", "??", "??????", "??????"}

	for _, canonical_name := range canonical_names {
		if street_type == canonical_name {
			return street_type
		}
	}

	if street_type == "??????????" || street_type == "??????" || street_type == "????????" {
		return "??-??"
	}

	if street_type == "??????????????" || street_type == "??????" || street_type == "??????????" {
		return "??-??"
	}

	if street_type == "????????????" || street_type == "??????" {
		return "??????-??"
	}

	if street_type == "??????????" || street_type == "??" {
		return "????"
	}

	if street_type == "????????????????" || street_type == "????????????" {
		return "??????"
	}

	if street_type == "??????????????" || street_type == "??" || street_type == "????????" {
		return "????"
	}

	if street_type == "????????????????????" || street_type == "????????" || street_type == "??????????" {
		return "??????"
	}

	if street_type == "????????????" || street_type == "??????????" {
		return "????????????"
	}

	if street_type == "????????????????????" || street_type == "??????????" || street_type == "??????????????" || street_type == "????????????????" {
		return "??????"
	}

	if street_type == "????????????????" || street_type == "????" {
		return "????-????"
	}

	return ""
}

func SplitHouseNumberToNumberAndKorpus(houseNumber string) (string, string) {
	// TODO change to regex
	korpusLabels := []string{"??????.", "????????.", "????????????????", "????????????"}

	for _, korpusLabel := range korpusLabels {
		splittedHouseNumber := strings.Split(houseNumber, korpusLabel)
		if len(splittedHouseNumber) == 2 {
			return strings.TrimSpace(splittedHouseNumber[0]),
				strings.TrimSpace(splittedHouseNumber[1])
		}
	}
	return houseNumber, ""
}

func getCityName(city string) string {
	splitted := strings.Split(city, " ")
	if len(splitted) == 1 {
		return city
	}

	return splitted[1]
}

func makeCompoundAddress(city, streetType, streetName, houseNumber, korpus string) string {
	return strings.Join([]string{city, streetType, streetName, houseNumber, korpus}, ",")
}
