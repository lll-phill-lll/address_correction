package fiasdata

import (
	"fmt"
	"github.com/lll-phill-lll/address_correction/logger"
	"github.com/lll-phill-lll/address_correction/pkg/strops"
	"regexp"
	"strings"
)

var Storage storage

// TODO use maps
type storage struct {
	Addresses map[string][]Address
}

func SetNewStorage(addresses []Address) {
	Storage = storage{make(map[string][]Address)}
	for _, address := range addresses {
		Storage.AddAddress(address)
	}
}

const (
	EXACT_MATCH int = iota
	SUBSTRING_MATCH
	ABBREVIATION_MATCH
)

func areStreetsEqual(correct string, given string, iter int) bool {
	if correct == "ANY" {
		return true
	}

	given = strings.ToLower(given)

	switch iter {
	case EXACT_MATCH:
		return correct == given
	case SUBSTRING_MATCH:
		return strings.Contains(correct, given) || strings.Contains(given, correct)
	case ABBREVIATION_MATCH:
		if strings.ContainsRune(given, '.') {
			regexString := fmt.Sprintf("^%s$", strings.ReplaceAll(given, ".", ".*"))
			reg, err := regexp.Compile(regexString)
			if err != nil {
				return false
			}
			return reg.MatchString(correct)
		}
		return false
	}

	return false
}

func (s *storage) AddAddress(address Address) {
	s.Addresses[address.City] = append(s.Addresses[address.City], address)
}

func (s *storage) FindCityWithErrors(city string, maxErrors int) string {

	for i := 0; i != maxErrors; i++ {
		for storedCity, _ := range s.Addresses {
			if strops.LowerEqualWithErrors(storedCity, city, i) {
				return storedCity
			}
		}
	}
	return ""
}

func (s *storage) GetFias(city string,
	street_type string,
	street string,
	house_num string,
	korpus string) string {

	cityAddresses, ok := s.Addresses[city]
	if !ok {
		correctedCity := s.FindCityWithErrors(city, 3)

		if correctedCity == "" {
			return "No such city loaded"
		}

		cityAddresses = s.Addresses[correctedCity]
	}

	for i := 0; i != 3; i++ {
		for _, address := range cityAddresses {
			if strops.LowerEqualWithErrors(address.City, city, 0) || city == "ANY" {
				if address.StreetType == strings.ToLower(street_type) || street_type == "ANY" {
					if areStreetsEqual(address.FormalName, street, i) {
						if address.HouseNum == strings.ToLower(house_num) || house_num == "ANY" {
							logger.Info.Println(address)
							if address.Korpus == strings.ToLower(korpus) || korpus == "ANY" {
								logger.Info.Println("Found address", address)
								return address.FIAS

							}
						}
					}

				}

			}
		}
	}
	return "Not found"
}
