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

// expects city to be in s.Addresses
func (s *storage) findFIASInCity(city string,
	street_type string,
	street string,
	house_num string,
	korpus string) Address {

	for i := 0; i != 3; i++ {
		for _, address := range s.Addresses[city] {
			if address.StreetType == strings.ToLower(street_type) || street_type == "ANY" {
				if areStreetsEqual(address.FormalName, street, i) {
					if address.HouseNum == strings.ToLower(house_num) || house_num == "ANY" {
						logger.Info.Println(address)
						if address.Korpus == strings.ToLower(korpus) || korpus == "ANY" {
							logger.Info.Println("Found address", address)
							return address

						}
					}
				}
			}
		}
	}

	return Address{}
}

func (s *storage) GetAddress(city string,
	street_type string,
	street string,
	house_num string,
	korpus string) Address {

	if city != "ANY" {
		_, ok := s.Addresses[city]
		if !ok {
			correctedCity := s.FindCityWithErrors(city, 3)

			if correctedCity == "" {
				return Address{}
			}
			return s.findFIASInCity(correctedCity, street_type, street, house_num, korpus)
		} else {
			return s.findFIASInCity(city, street_type, street, house_num, korpus)
		}
	}

	for city, _ := range s.Addresses {
		address := s.findFIASInCity(city, street_type, street, house_num, korpus)
		if address.FIAS != "" {
			return address
		}
	}

	return Address{}
}
