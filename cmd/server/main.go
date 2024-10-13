package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Andrey-Kachow/xchange-r8-microservice/internal/api"
	"github.com/Andrey-Kachow/xchange-r8-microservice/internal/app"
	"github.com/Andrey-Kachow/xchange-r8-microservice/internal/models"
)

func sampleHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Conent-Type", "text/html")
	writer.Write([]byte("<h1>Hello World!</h1>"))
}

func initApp() {
	err := models.InitAllSupportedCurrencyList()
	if err != nil {
		log.Fatal("Failed to initialize app:", err)
	}
	err = app.InitAppContext()
	if err != nil {
		log.Fatal("Failed to initialize app Context:", err)
	}
}

func main() {
	fmt.Println("Starting the app")
	initApp()
	http.HandleFunc("/", sampleHandler)
	http.HandleFunc("/rate", api.RateHandler)
	http.HandleFunc("/convert", api.ConvertHandler)
	http.HandleFunc("/rates", api.RatesHandler)
	http.HandleFunc("/historical", api.HistoricalHandler)
	http.HandleFunc("/currencies", api.CurrenciesHandler)
	http.HandleFunc("/health", api.HealthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
