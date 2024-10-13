package xchange

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const openExchangeRatesURL = "https://openexchangerates.org/api/latest.json"
const errorRate = -1.0

type OpenExchangeRatesOrgRateProvider struct {
	appID string
}

func CreateOpenExchangeRatesOrgRateProvider() (*OpenExchangeRatesOrgRateProvider, error) {
	appID := os.Getenv("XCHANGE_R8_OPENEXCHANGERATES_APP_ID")
	if appID == "" {
		return nil, errors.New("openexchangerates.org APP ID is empty. Please set XCHANGE_R8_OPENEXCHANGERATES_APP_ID")
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

func (orp *OpenExchangeRatesOrgRateProvider) getLatest(baseCurrency string, targetCurrencies []string) (map[string]float64, error) {
	if baseCurrency != "USD" {
		return nil, errors.New("openxchangerates.org free tier only supports USD as base currency")
	}

	symbols := strings.Join(targetCurrencies, ",")

	queryParams := url.Values{}
	queryParams.Add("app_id", orp.appID)
	queryParams.Add("base", baseCurrency)
	queryParams.Add("symbols", symbols)
	requestURL := fmt.Sprintf("%s?%s", openExchangeRatesURL, queryParams.Encode())

	response, err := http.Get(requestURL)
	if err != nil {
		fmt.Println("Error sending the response", err)
		return nil, err
	}

	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var responseJson apiLatestJsonResponse
	err = json.Unmarshal(bodyBytes, &responseJson)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	return responseJson.Rates, nil
}

func (orp *OpenExchangeRatesOrgRateProvider) GetRate(baseCurrency string, targetCurrency string) (float64, error) {

	rates, err := orp.getLatest(baseCurrency, []string{targetCurrency})
	if err != nil {
		fmt.Println("Error accessing openexchangerates.org")
		return errorRate, err
	}

	targetRate, ok := rates[targetCurrency]
	if !ok {
		fmt.Println("Response missing the target currency")
		return errorRate, errors.New("failed to access openxchangerates.org/latest")
	}

	return targetRate, nil
}

func (oerorp *OpenExchangeRatesOrgRateProvider) GetRates(baseCurrency string, targetCurrencies []string) (map[string]float64, error) {
	return oerorp.getLatest(baseCurrency, targetCurrencies)
}
