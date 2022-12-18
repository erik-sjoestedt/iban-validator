package iban

import "testing"

func TestValidateValidIban(t *testing.T) {
	is_valid := Validate("1")
	if !is_valid {
		t.Log("1 should be valid IBAN")
		t.Fail()
	}
}

func TestValidateInvalidIban(t *testing.T) {
	is_valid := Validate("2")
	if is_valid {
		t.Log("2 should be invalid IBAN")
		t.Fail()
	}
}
