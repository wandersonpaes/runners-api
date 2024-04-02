package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wandersonpaes/runners-api/internal/app"
)

func main() {
	fmt.Println("Runners API is running!")
	log.Fatal(http.ListenAndServe(":5000", app.Start()))
}
