package xchange

type RateProvider interface {
	GetRate(baseCurrency string, targetCurrency string) (float64, error)
	GetRates(baseCurrency string, targetCurrencies []string) (map[string]float64, error)
}
