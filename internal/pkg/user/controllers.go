package user

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/wandersonpaes/runners-api/internal/pkg/database"
)

func Create(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Conectar()
	if err != nil {
		log.Fatal(err)
	}

	userTable := newUserConnection(db)
	userID, err := userTable.create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("User created on ID: %d", userID)))
}

func SearchAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching for all Users!"))
}

func Search(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching a User!"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating a User!"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a User!"))
}
