package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Andrey-Kachow/xchange-r8-microservice/internal/xchange"
)

/*
Query Parameters:

	base (required): The base currency (e.g., USD).
	target (required): The target currency (e.g., EUR).
*/
func RateHandler(writer http.ResponseWriter, request *http.Request) {

	queryParams := request.URL.Query()
	baseCurrency := queryParams.Get("base")
	targetCurrency := queryParams.Get("target")

	if baseCurrency == "" || targetCurrency == "" {
		fmt.Println("Error: base or target currency is not provided")
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rateProvider, err := xchange.CreateOpenExchangeRatesOrgRateProvider()
	if err != nil {
		fmt.Println("Failed to create exchange rate provider", err)
		return
	}

	rate, err := rateProvider.GetRate(baseCurrency, targetCurrency)

	if err != nil {
		fmt.Println("Failed to get echange rate from provider:", err)
		return
	}

	rateString := strconv.FormatFloat(rate, 'g', 5, 64)

	fmt.Println(rateString)
	writer.Write([]byte(rateString))
}

/*
Query Parameters:

	base (required): The currency to convert from (e.g., USD).
	target (required): The currency to convert to (e.g., GBP).
	amount (required): The amount to convert (e.g., 100).
*/
func ConvertHandler(writer http.ResponseWriter, request *http.Request) {
}

/*
Return JSON example

	{
		"base": "USD",
		"rates": {
			"EUR": 0.85,
			"GBP": 0.725,
			"JPY": 110.15
		},
		"timestamp": "2023-09-01T10:00:00Z"
	}
*/
func RatesHandler(writer http.ResponseWriter, request *http.Request) {
}

/*
Query Parameters:

	base (required): The base currency (e.g., USD).
	target (required): The target currency (e.g., CAD).
	start_date (required): The start date for historical data (e.g., 2023-01-01).
	end_date (optional): The end date for historical data (e.g., 2023-01-31). If not provided, it will return data up to the current date.

Example Return Json:

	{
		"base": "USD",
		"target": "CAD",
		"rates": [
			{"date": "2023-01-01", "rate": 1.25},
			{"date": "2023-01-02", "rate": 1.26},
			{"date": "2023-01-03", "rate": 1.24}
		]
	}
*/
func HistoricalHandler(writer http.ResponseWriter, request *http.Request) {
}

/*
Return Json:

	{
		"currencies": [
			{"code": "USD", "name": "United States Dollar"},
			{"code": "EUR", "name": "Euro"},
			{"code": "JPY", "name": "Japanese Yen"},
			...
		]
	}
*/
func CurrenciesHandler(writer http.ResponseWriter, request *http.Request) {
}

// return Json: { "status": "healthy" }
func HealthHandler(writer http.ResponseWriter, request *http.Request) {
}
