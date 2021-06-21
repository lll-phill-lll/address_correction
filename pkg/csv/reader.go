package csv

import (
    "encoding/csv"
    "os"
)

func Read(path string) ([][]string, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        return nil, err
    }

    return records, nil
}
