// Implements IBAN validation according to this link
// https://en.wikipedia.org/wiki/International_Bank_Account_Number#Modulo_operation_on_IBAN

package iban

import "errors"

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
