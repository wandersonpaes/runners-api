package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/wandersonpaes/runners-api/internal/app"
	"github.com/wandersonpaes/runners-api/internal/pkg/database"
)

func main() {
	database.SetUpDatabase()

	fmt.Println("Runners API is running!")
	r := router.CreateMux()

	log.Fatal(http.ListenAndServe(":5000", r))
}
