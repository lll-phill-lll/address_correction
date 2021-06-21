package fiasdata

import (
    "github.com/lll-phill-lll/address_correction/pkg/csv"
	"github.com/lll-phill-lll/address_correction/logger"
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
func FromCSV(path string) ([]Address, error) {
    rows, err := csv.Read(path)
    if err != nil {
        logger.Error.Println("Can't read csv", err)
        return nil, err
    }

    // TODO add titles check, primary validation

    var addresses []Address

    for i, row := range rows {
        // skip table title
        if i == 0 {
            continue
        }

        address := Address{}
        address.CityLong = row[CITY]
        address.City = shortenCity(row[CITY])
        address.StreetType = row[STREETTYPE]
        address.FormalName = row[FORMALNAME]
        address.HouseNum = row[HOUSENUM]
        address.FIAS = row[HOUSEGUID]

        if row[KORPUS] != "NULL" {
            address.Corpus = row[KORPUS]
        } else if row[STROENIE] != "NULL" {
            address.Corpus = row[STROENIE]
        }

        addresses = append(addresses, address)
    }

    return addresses, nil
}

func shortenCity(city string) string {
    return city
}
