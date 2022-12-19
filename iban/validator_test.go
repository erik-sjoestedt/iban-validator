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

func TestEncodeIban(t *testing.T) {
	input := "AY97123BZ456"
	expected := "1231135456103497"
	actual := encodeIban(input)
	if actual != expected {
		t.Log(input + " should be encoded to " + expected + "; was " + actual)
		t.Fail()
	}
}

func TestEncodeChar(t *testing.T) {
	chars_and_codes := map[byte]string{
		byte('A'): "10",
		byte('B'): "11",
		byte('Y'): "34",
		byte('Z'): "35",
	}
	for char, code := range chars_and_codes {
		actual := encodeChar(char)
		if actual != code {
			t.Log(string(char) + " should be encoded to " + code + "; was " + actual)
			t.Fail()
		}
	}
}

func TestMod97(t *testing.T) {
	inputs_and_expected := map[string]int{
		"321428291":                    70,
		"32142829100":                  16,
		"3214282912345698765432161182": 1,
	}
	for input, expected := range inputs_and_expected {
		actual := mod97(input)
		if actual != expected {
			t.Logf("%s mod 97 should be %d; was %d", input, expected, actual)
			t.Fail()
		}
	}
}
