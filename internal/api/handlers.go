package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Andrey-Kachow/xchange-r8-microservice/internal/xchange"
)

func RateHandler(writer http.ResponseWriter, request *http.Request) {

	rateProvider, err := xchange.CreateOpenExchangeRatesOrgProvider()
	if err != nil {
		fmt.Println("Failed to create exchange rate provider", err)
		return
	}

	rate, err := rateProvider.GetRate("USD", "EUR")

	if err != nil {
		fmt.Println("Failed to get echange rate from provider:", err)
		return
	}

	rateString := strconv.FormatFloat(rate, 'g', 5, 64)

	fmt.Println(rateString)
	writer.Write([]byte(rateString))
}

func ConvertHandler(writer http.ResponseWriter, request *http.Request) {
}

func RatesHandler(writer http.ResponseWriter, request *http.Request) {
}

func HistoricalHandler(writer http.ResponseWriter, request *http.Request) {
}

func CurrenciesHandler(writer http.ResponseWriter, request *http.Request) {
}

func HealthHandler(writer http.ResponseWriter, request *http.Request) {
}
