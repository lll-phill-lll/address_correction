package address

import (
	// "fmt"
	"github.com/lll-phill-lll/address_correction/internal/fiasdata"
	"github.com/lll-phill-lll/address_correction/logger"
	parser "github.com/openvenues/gopostal/parser"
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
//     world_region: currently only used for appending “West Indies” after the country name, a pattern frequently used in the English-speaking Caribbean e.g. “Jamaica, West Indies”
func Correct(address string) string {
	parsed_values := parser.ParseAddress(address)
	logger.Info.Println("Parsed", address, "into", parsed_values)

	street_name, street_type := "", ""
	house_number := ""
	city := "красноярск"
	road := ""

	for _, parsed_value := range parsed_values {
		if parsed_value.Label == "road" {
			road = parsed_value.Value
			logger.Info.Println("Found road:", road)
			street_name, street_type = SpiltRoadIntoNameAndType(road)
		}
		if parsed_value.Label == "house_number" {
			house_number = parsed_value.Value
			logger.Info.Println("Found house_number:", house_number)
		}

	}

	if road != "" {
		logger.Info.Println("Found road:", road)
		street_name, street_type = SpiltRoadIntoNameAndType(road)
		street_type = StreetTypeToCanonical(street_type)
	}

	korpus := "ANY"
	logger.Info.Println("city, street_type, street_name, house_number, korpus:", city, street_type, street_name, house_number, korpus)
	fias := fiasdata.Storage.GetFias(city, street_type, street_name, house_number, "ANY")

	return fias // fmt.Sprintln(parsed)
}

func SpiltRoadIntoNameAndType(road string) (string, string) {
	splitted_road := strings.FieldsFunc(road, Split)
	logger.Info.Println("Splitted:", road, "to", splitted_road)

	if len(splitted_road) == 1 {
		// TODO consider returning ANY
		return road, ""
	}

	street_type := ""
	street_name := ""

	if IsStreetType(splitted_road[1]) {
		street_type = splitted_road[1]
		street_name = splitted_road[0]
	} else if IsStreetType(splitted_road[0]) {
		street_type = splitted_road[0]
		street_name = splitted_road[1]
	}

	return street_name, street_type
}

func Split(r rune) bool {
	return r == ' ' || r == ','
}

func IsStreetType(street_type string) bool {
	return true
}

func StreetTypeToCanonical(street_type string) string {
	if street_type == "набережная" {
		return "наб"
	}

	return street_type
}
