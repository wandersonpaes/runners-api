package user

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/wandersonpaes/runners-api/internal/pkg/database"
	"github.com/wandersonpaes/runners-api/internal/pkg/response"
)

func Create(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	userTable := newUserConnection(db)
	user.ID, err = userTable.create(user)
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
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
