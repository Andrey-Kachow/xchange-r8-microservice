package app

import (
	"fmt"

	"github.com/Andrey-Kachow/xchange-r8-microservice/internal/xchange"
)

type AppContext struct {
	RatesProvider xchange.RateProvider
}

var appContext AppContext

func InitAppContext() error {
	rateProvider, err := xchange.CreateOpenExchangeRatesOrgRateProvider()
	if err != nil {
		fmt.Println("Failed to create rateProvider for appContext:", err)
		return err
	}

	appContext = AppContext{
		RatesProvider: rateProvider,
	}
	return nil
}

func GetAppContext() *AppContext {
	return &appContext
}
