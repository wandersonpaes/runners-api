package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/wandersonpaes/runners-api/internal/app"
)

func main() {
	fmt.Println("Runners API is running!")
	r := router.CreateMux()

	log.Fatal(http.ListenAndServe(":5000", r))
}
