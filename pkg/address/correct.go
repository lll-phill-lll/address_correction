package address

import (
    "fmt"
    parser "github.com/openvenues/gopostal/parser"
)

func Correct(address string) string {
    parsed := parser.ParseAddress(address)
    return fmt.Sprintln(parsed)
}
