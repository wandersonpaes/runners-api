package userRunner

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err = user.Prepare("register"); err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	userTable := NewUserConnection(db)
	user.ID, err = userTable.create(user)
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func SearchAll(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	db, err := database.Conectar()
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userTable := NewUserConnection(db)
	users, err := userTable.search(nameOrNick)
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func Search(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userTable := NewUserConnection(db)
	user, err := userTable.searchByID(userID)
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func Update(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

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

	if err = user.Prepare("edition"); err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userTable := NewUserConnection(db)
	if err = userTable.update(userID, user); err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userTable := NewUserConnection(db)
	if err = userTable.delete(userID); err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}
