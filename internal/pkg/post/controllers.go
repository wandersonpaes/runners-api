package post

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/wandersonpaes/runners-api/internal/pkg/database"
	"github.com/wandersonpaes/runners-api/internal/pkg/response"
	"github.com/wandersonpaes/runners-api/internal/pkg/security"
)

func Create(w http.ResponseWriter, r *http.Request) {
	userIdOnToken, err := security.ExtracUserID(r)
	if err != nil {
		response.ERR(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post Posts
	if err = json.Unmarshal(bodyRequest, &post); err != nil {
		response.ERR(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userIdOnToken

	db, err := database.Conectar()
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postTable := newPostConnection(db)
	post.ID, err = postTable.create(post)
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, post)
}

func SearchAll(w http.ResponseWriter, r *http.Request) {

}

func SearchOne(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
