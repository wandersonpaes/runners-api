package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wandersonpaes/runners-api/internal/app"
	"github.com/wandersonpaes/runners-api/internal/pkg/auth"
	"github.com/wandersonpaes/runners-api/internal/pkg/database"
)

func main() {
	database.SetUp()
	auth.SetUp()

	fmt.Println("Runners API is running!")
	r := app.CreateMux()

	log.Fatal(http.ListenAndServe(":5000", r))
}
