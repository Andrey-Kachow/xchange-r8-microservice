package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

var supportedCurrencies CurrencyList

func GetAllSupportedCurrencyList() CurrencyList {
	return supportedCurrencies
}

func InitAllSupportedCurrencyList() error {
	file, err := os.Open("resources/currencies.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	csvRecords, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return err
	}

	var currencies []Currency
	for _, record := range csvRecords {
		currency := Currency{
			Code: record[0],
			Name: record[1],
		}
		currencies = append(currencies, currency)
	}

	// Sorting for easier Lookup later
	//
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Code < currencies[j].Code
	})

	supportedCurrencies = CurrencyList{Currencies: currencies}

	return nil
}

// CurrencyIsSupported searches sorted list of currencies using binary search agorithm
func CurrencyIsSupported(currencyCode string) bool {
	low := 0
	high := len(supportedCurrencies.Currencies)
	for low <= high {
		mid := (low + high) / 2
		midCode := supportedCurrencies.Currencies[mid].Code

		if currencyCode == midCode {
			return true
		}

		if currencyCode < midCode {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
