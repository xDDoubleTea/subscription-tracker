package core

import "errors"

const allCurrencyLen = 3

var supportedCurrency = [allCurrencyLen]string{"ntd", "jpy", "usd"}

type Currency struct {
	Name string
}

func NewCurrency(name string) (Currency, error) {
	if !ValidCurrency(name) {
		return Currency{}, errors.New("Currency should be one of ntd, jpy or usd.")
	}
	return Currency{Name: name}, nil
}

func ValidCurrency(input string) bool {

	supported := false
	for i := 0; i < allCurrencyLen; i++ {
		supported = supported || (supportedCurrency[i] == input)
	}
	return supported
}
