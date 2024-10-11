package xchange

type RateProvider interface {
	GetRate(baseCurrency string, targetCurrency string) (float64, error)
}
