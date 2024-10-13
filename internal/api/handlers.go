package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Andrey-Kachow/xchange-r8-microservice/internal/app"
	"github.com/Andrey-Kachow/xchange-r8-microservice/internal/models"
)

/*
Query Parameters:

	base (required): The base currency (e.g., USD).
	target (required): The target currency (e.g., EUR).
	return json example: {"target_amount": 123.48}
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

	rateProvider := app.GetAppContext().RatesProvider
	rate, err := rateProvider.GetRate(baseCurrency, targetCurrency)

	if err != nil {
		fmt.Println("Failed to get echange rate from provider:", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Println(rate)
	json.NewEncoder(writer).Encode(map[string]float64{
		"rate": rate,
	})
}

/*
Query Parameters:

	base (required): The currency to convert from (e.g., USD).
	target (required): The currency to convert to (e.g., GBP).
	amount (required): The amount to convert (e.g., 100).
*/
func ConvertHandler(writer http.ResponseWriter, request *http.Request) {

	queryParams := request.URL.Query()
	baseCurrency := queryParams.Get("base")
	targetCurrency := queryParams.Get("target")
	amountString := queryParams.Get("amount")

	if baseCurrency == "" || targetCurrency == "" || amountString == "" {
		fmt.Println("Error: base or target currency or amount is not provided")
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	amount, err := strconv.Atoi(amountString)

	if err != nil {
		fmt.Printf("Bad amount parameter: \"%s\"\n", amountString)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rateProvider := app.GetAppContext().RatesProvider
	rate, err := rateProvider.GetRate(baseCurrency, targetCurrency)
	if err != nil {
		fmt.Println("Failed to get echange rate from provider:", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	targetAmount := float64(amount) * rate
	json.NewEncoder(writer).Encode(map[string]float64{
		"target_amount": targetAmount,
	})
}

/*
Query Parameters:

	base (required): The currency to convert from (e.g., USD).
	targets (required): The comma separated list of currencies to convert to (e.g., GBP,EUR,JPY).

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

	queryParams := request.URL.Query()
	baseCurrency := queryParams.Get("base")
	targetCurrencyList := queryParams.Get("targets")

	if baseCurrency == "" || targetCurrencyList == "" {
		fmt.Println("Error: base or target currencies is not provided in query parameters")
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	targetCurrencies := strings.Split(targetCurrencyList, ",")
	for _, targetCurrency := range targetCurrencies {
		if !models.CurrencyIsSupported(targetCurrency) {
			fmt.Println("Currency is not supported:", targetCurrency)
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	rateProvider := app.GetAppContext().RatesProvider
	rates, err := rateProvider.GetRates(baseCurrency, targetCurrencies)
	if err != nil {
		fmt.Println("Failed to get echange rates from provider:", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	json.NewEncoder(writer).Encode(map[string]interface{}{
		"base":      baseCurrency,
		"rates":     rates,
		"timestamp": time.Now().Format(time.RFC3339),
	})
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

	currencyList := models.GetAllSupportedCurrencyList()

	jsonData, err := json.Marshal(currencyList)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	writer.Write(jsonData)
}

// return Json: { "status": "healthy" }
func HealthHandler(writer http.ResponseWriter, request *http.Request) {
	//
	// TODO: health check
	//
	json.NewEncoder(writer).Encode(map[string]string{
		"status": "healthy",
	})
}
