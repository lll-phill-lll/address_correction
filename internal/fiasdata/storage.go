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
	Addresses []Address
}

func SetNewStorage(addresses []Address) {
	Storage = storage{addresses}
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

func (s *storage) GetFias(city string,
	street_type string,
	street string,
	house_num string,
	korpus string) string {
	logger.Info.Println("Storage size:", len(s.Addresses))

	for i := 0; i != 3; i++ {

		for _, address := range s.Addresses {
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
