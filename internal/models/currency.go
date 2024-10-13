package models

type Currency struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type CurrencyList struct {
	Currencies []Currency `json:"currencies"`
}
