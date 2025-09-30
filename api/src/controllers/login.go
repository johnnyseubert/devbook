package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/johnnyseubert/devbook/src/database"
	"github.com/johnnyseubert/devbook/src/models"
	"github.com/johnnyseubert/devbook/src/repositories"
	"github.com/johnnyseubert/devbook/src/responses"
	"github.com/johnnyseubert/devbook/src/security"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UsersRepository(db)
	userSavedOnDb, err := repository.GetByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.Compare(userSavedOnDb.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := security.GenerateToken(userSavedOnDb.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
