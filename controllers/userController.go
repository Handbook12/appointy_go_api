package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-mongodb/models"
	. "go-rest-mongodb/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var usersRepository UsersRepository

var GetAllUsers = func(w http.ResponseWriter, r * http.Request) {
	users, err := usersRepository.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

var GetUserById = func(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln("ALL USERS")
}

var CreateUser = func(w http.ResponseWriter, r * http.Request) {
	
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	insertResult, err := usersRepository.Insert(user);
	if  err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	user.ID = insertResult.(primitive.ObjectID)
	respondWithJson(w, http.StatusCreated, user)
}

var UpdateUser = func(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln("User Updated")
}

var DeleteUser = func(w http.ResponseWriter, r * http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if err := usersRepository.Delete(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
