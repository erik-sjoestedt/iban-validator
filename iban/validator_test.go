package iban

import "testing"

func TestValidateValidIban(t *testing.T) {
	valid_ibans := [1]string{"SE4550000000058398257466"}
	for i := 0; i < len(valid_ibans); i++ {
		if Validate(valid_ibans[i]) != nil {
			t.Log(valid_ibans[i] + " should be valid IBAN")
			t.Fail()
		}
	}
}

func TestValidateInvalidIban(t *testing.T) {
	invalid_ibans := [1]string{
		"SE123", // incorrect SE length
	}
	for i := 0; i < len(invalid_ibans); i++ {
		if Validate(invalid_ibans[i]) == nil {
			t.Log(invalid_ibans[i] + " should be invalid IBAN")
			t.Fail()
		}
	}
}
