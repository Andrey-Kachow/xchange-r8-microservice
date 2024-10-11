package xchange

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const openExchangeRatesURL = "https://openexchangerates.org/api/latest.json"
const errorRate = -1.0

type OpenExchangeRatesOrgRateProvider struct {
	appID string
}

func CreateOpenExchangeRatesOrgRateProvider() (*OpenExchangeRatesOrgRateProvider, error) {
	appID := os.Getenv("XCHANGE_R8_OPENEXCHANGERATES_APP_ID")
	if appID == "" {
		return nil, errors.New("openxchangerates.org APP ID is empty. Please set XCHANGE_R8_OPENEXCHANGERATES_APP_ID")
	}
	provider := OpenExchangeRatesOrgRateProvider{
		appID: appID,
	}
	return &provider, nil
}

type apiLatestJsonResponse struct {
	Disclaimer string             `json:"disclaimer"`
	License    string             `json:"license"`
	Timestamp  int                `json:"timestamp"`
	Base       string             `json:"base"`
	Rates      map[string]float64 `json:"rates"`
}

func (oerorp *OpenExchangeRatesOrgRateProvider) GetRate(baseCurrency string, targetCurrency string) (float64, error) {

	if baseCurrency != "USD" {
		return errorRate, errors.New("openxchangerates.org free tier only supports USD as base currency")
	}

	queryParams := url.Values{}
	queryParams.Add("app_id", oerorp.appID)
	queryParams.Add("base", baseCurrency)
	queryParams.Add("symbols", targetCurrency)
	requestURL := fmt.Sprintf("%s?%s", openExchangeRatesURL, queryParams.Encode())

	response, err := http.Get(requestURL)
	if err != nil {
		fmt.Println("Error sending the response", err)
		return errorRate, err
	}

	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return errorRate, err
	}

	var responseJson apiLatestJsonResponse
	err = json.Unmarshal(bodyBytes, &responseJson)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return errorRate, err
	}

	rate, ok := responseJson.Rates[targetCurrency]
	if !ok {
		return errorRate, errors.New("openxchangerates.org returned Json without target currency\n")
	}

	return rate, nil
}
