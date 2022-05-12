package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GuilhermePC09/api-rest-blog-go/services"
	"github.com/gorilla/mux"
)

type UserRequest struct {
	Id       int64
	Name     string
	Email    string
	Password string
}

type EditUserRequest struct {
	Id       int64
	Name     string
	Email    string
	Type     string
	Password string
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user UserRequest

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	checkedUser, err2 := services.CreateUserService(user.Name, user.Email, user.Password)

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(checkedUser)
}
func ListUsers(w http.ResponseWriter, r *http.Request) {

	checkedUserList := services.ListUsersService()

	jsonUsers, err := json.MarshalIndent(checkedUserList, "", " ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(jsonUsers))

}

func EditUser(w http.ResponseWriter, r *http.Request) {

	var editUser EditUserRequest

	err := json.NewDecoder(r.Body).Decode(&editUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	edit, err2 := services.EditUserService(editUser.Type, editUser.Id, editUser.Email, editUser.Name, editUser.Password)

	if err2 != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err3 := json.MarshalIndent(edit, "", " ")

	if err3 != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, string(response))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["userId"]

	services.DeletePostService(id)

	response := make(map[string]string)
	response["message"] = "User deletado com sucesso"

	checkResponse, err2 := json.Marshal(response)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	w.Write(checkResponse)
}
