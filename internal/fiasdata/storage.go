package fiasdata

import (
	"github.com/lll-phill-lll/address_correction/logger"
	"github.com/lll-phill-lll/address_correction/pkg/strops"
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

func (s *storage) GetFias(city string,
	street_type string,
	street string,
	house_num string,
	korpus string) string {
	logger.Info.Println("Storage size:", len(s.Addresses))
	wasRestart := false

	for true {

		for _, address := range s.Addresses {
			if strops.LowerEqualWithErrors(address.City, city, 0) || city == "ANY" {
				if address.StreetType == strings.ToLower(street_type) || street_type == "ANY" {
					if address.FormalName == strings.ToLower(street) || street == "ANY" ||
						wasRestart && (strings.Contains(address.FormalName, strings.ToLower(street)) || strings.Contains(strings.ToLower(street), address.FormalName)) {
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
		if wasRestart {
			break
		}
		wasRestart = true
	}
	return "Not found"
}
