package userRunner

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/wandersonpaes/runners-api/internal/pkg/database"
	"github.com/wandersonpaes/runners-api/internal/pkg/response"
	"github.com/wandersonpaes/runners-api/internal/pkg/security"
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

	userIdOnToken, err := security.ExtracUserID(r)
	if err != nil {
		response.ERR(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIdOnToken {
		response.ERR(w, http.StatusForbidden, errors.New("it's not possible update a user that's not yours"))
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

	userIdOnToken, err := security.ExtracUserID(r)
	if err != nil {
		response.ERR(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIdOnToken {
		response.ERR(w, http.StatusForbidden, errors.New("it's not possible delete a user that's not yours"))
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

func Follow(w http.ResponseWriter, r *http.Request) {
	followerID, err := security.ExtracUserID(r)
	if err != nil {
		response.ERR(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		response.ERR(w, http.StatusForbidden, errors.New("it's not possible follow yourself"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	followersTable := NewUserConnection(db)
	if err = followersTable.follow(userID, followerID); err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	followerID, err := security.ExtracUserID(r)
	if err != nil {
		response.ERR(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		response.ERR(w, http.StatusForbidden, errors.New("it's not possible unfollow yourself"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	followersTable := NewUserConnection(db)
	if err = followersTable.unfollow(userID, followerID); err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func SearchFollowers(w http.ResponseWriter, r *http.Request) {
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

	followersTable := NewUserConnection(db)
	followers, err := followersTable.searchFollowers(userID)
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

func SearchFollowing(w http.ResponseWriter, r *http.Request) {
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

	followersTable := NewUserConnection(db)
	followers, err := followersTable.searchFollowing(userID)
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIDOnToken, err := security.ExtracUserID(r)
	if err != nil {
		response.ERR(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

	if userIDOnToken != userID {
		response.ERR(w, http.StatusForbidden, errors.New("it's not possible to update password from another user"))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

	var newPassword NewPassword
	if err := json.Unmarshal(bodyRequest, &newPassword); err != nil {
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
	passwordSavedInDatabase, err := userTable.searchPassword(userID)
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(passwordSavedInDatabase, newPassword.Current); err != nil {
		response.ERR(w, http.StatusUnauthorized, errors.New("the current password is incorrect"))
		return
	}

	passwordWithHash, err := security.Hash(newPassword.New)
	if err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

	if err = userTable.updatePassword(userID, string(passwordWithHash)); err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
