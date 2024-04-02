package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wandersonpaes/runners-api/internal/pkg/database"
	"github.com/wandersonpaes/runners-api/internal/pkg/response"
	"github.com/wandersonpaes/runners-api/internal/pkg/security"
	"github.com/wandersonpaes/runners-api/internal/pkg/userRunner"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user userRunner.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userTable := userRunner.NewUserConnection(db)
	userSalveInDatabase, err := userTable.SearchByEmail(user.Email)
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(userSalveInDatabase.Password, user.Password); err != nil {
		response.ERR(w, http.StatusUnauthorized, err)
		return
	}

	token, err := CreateToken(userSalveInDatabase.ID)
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(token)

	fmt.Println(userSalveInDatabase)
	w.Write([]byte("You are logged!"))
}
