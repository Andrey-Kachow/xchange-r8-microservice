package models

import (
	"encoding/csv"
	"fmt"
	"os"
)

func GetAllSupportedCurrencyList() (*CurrencyList, error) {

	file, err := os.Open("resources/currencies.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	csvRecords, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil, err
	}

	var currencies []Currency
	for _, record := range csvRecords {
		currency := Currency{
			Code: record[0],
			Name: record[1],
		}
		currencies = append(currencies, currency)
	}
	return &CurrencyList{Currencies: currencies}, nil
}
