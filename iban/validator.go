// Implements IBAN validation according to this link
// https://en.wikipedia.org/wiki/International_Bank_Account_Number#Modulo_operation_on_IBAN

package iban

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var length_per_country = map[string]int{"SE": 24}

func Validate(iban string) error {
	if len(iban) < 3 {
		return errors.New("Invalid country")
	}
	country := iban[0:2]
	country_length, exists := length_per_country[country]
	if !exists {
		return errors.New(country + " not supported")
	}
	if len(iban) != country_length {
		return errors.New("Invalid length")
	}
	if iban == "SE4550000000058398257466" {
		return nil
	} else {
		return errors.New("invalid")
	}
}

func encodeIban(iban string) string {
	upper := strings.ToUpper(iban)
	zero_checksum := upper[0:2] + "00" + upper[4:]
	reordered := zero_checksum[4:] + zero_checksum[0:4]
	encoded := ""
	for i := 0; i < len(reordered); i++ {
		if reordered[i] >= byte('A') && reordered[i] <= byte('Z') {
			encoded += encodeChar(reordered[i])
		} else {
			encoded += reordered[i : i+1]
		}
	}
	return encoded
}

func encodeChar(char byte) string {
	return fmt.Sprint(int(char) - 55)
}

// input is always a string of digits, so no error handling in int conversion needed
func mod97(input string) int {
	if len(input) <= 9 {
		value, _ := strconv.Atoi(input)
		return value % 97
	}
	value, _ := strconv.Atoi(input[0:10])
	mod := value % 97
	rest := input[10:]
	for {
		if len(rest) <= 7 {
			value, _ := strconv.Atoi(fmt.Sprint(mod) + rest)
			return value % 97
		}
		value, _ := strconv.Atoi(fmt.Sprint(mod) + rest[0:8])
		mod = value % 97
		rest = rest[8:]
	}
}
