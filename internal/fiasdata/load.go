package fiasdata

import (
	"github.com/lll-phill-lll/address_correction/logger"
	"github.com/lll-phill-lll/address_correction/pkg/csv"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	HOUSEGUID int = iota
	CITY
	STREETTYPE
	FORMALNAME
	HOUSENUM
	KORPUS
	HOUSETYPE
	STROENIE
	STATUSCSV
)

// FromCSV returns addresses read from csv file.
// expected format:
// HOUSEGUID CITY STREETTYPE FORMALNAME HOUSENUM KORPUS HOUSETYPE STROENIE STATUSCSV
func FromCSV(folderPath string) error {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		logger.Error.Println("can't read path", err)
		return err
	}

	var addresses []Address

	for _, f := range files {
		fileName := f.Name()
		if filepath.Ext(fileName) != ".csv" {
			continue
		}

		path := filepath.Join(folderPath, fileName)

		rows, err := csv.Read(path)
		if err != nil {
			logger.Error.Println("Can't read csv", err)
			return err
		}

		// TODO add titles check, primary validation

		for i, row := range rows {
			// skip table title
			if i == 0 {
				continue
			}

			address := Address{}
			address.CityLong = strings.ToLower(row[CITY])
			address.City = strings.ToLower(shortenCity(address.CityLong))
			address.StreetType = strings.ToLower(row[STREETTYPE])
			address.FormalName = strings.ToLower(row[FORMALNAME])
			address.HouseNum = strings.ToLower(row[HOUSENUM])
			address.FIAS = row[HOUSEGUID]

			if row[KORPUS] != "NULL" {
				address.Korpus = row[KORPUS]
			} else if row[STROENIE] != "NULL" {
				address.Korpus = row[STROENIE]
			}

			addresses = append(addresses, address)
		}
	}

	SetNewStorage(addresses)

	return nil
}

func shortenCity(city string) string {
	splittedCity := strings.Split(city, ",")
	cityPart := splittedCity[len(splittedCity)-1]

	cityNameSplitted := strings.Split(cityPart, " ")
	cityName := cityNameSplitted[len(cityNameSplitted)-1]

	return cityName
}
