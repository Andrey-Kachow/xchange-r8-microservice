package main

import (
	"fmt"
	"log"
	"net/http"
)

func sampleHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Conent-Type", "text/html")
	writer.Write([]byte("<h1>Hello World!</h1>"))
}

func main() {
	fmt.Println("Starting the app")
	http.HandleFunc("/", sampleHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
