package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Andrey-Kachow/xchange-r8-microservice/internal/api"
)

func sampleHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Conent-Type", "text/html")
	writer.Write([]byte("<h1>Hello World!</h1>"))
}

func main() {
	fmt.Println("Starting the app")
	http.HandleFunc("/", sampleHandler)
	http.HandleFunc("/rate", api.RateHandler)
	http.HandleFunc("/convert", api.ConvertHandler)
	http.HandleFunc("/rates", api.RateHandler)
	http.HandleFunc("/historical", api.HistoricalHandler)
	http.HandleFunc("/currencies", api.CurrenciesHandler)
	http.HandleFunc("/health", api.HealthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
