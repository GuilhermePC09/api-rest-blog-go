package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GuilhermePC09/api-rest-blog-go/services"
)

type UserRequest struct {
	Id       int64
	Name     string
	Email    string
	Password string
}

type EditRequest struct {
	Id       int64
	Email    string
	Type     string
	EditInfo string
}

type DeleteRequest struct {
	Id int64
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

	var editUser EditRequest

	err := json.NewDecoder(r.Body).Decode(&editUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	services.EditUserService(editUser.Id, editUser.Email, editUser.Type, editUser.EditInfo)
	response := make(map[string]string)
	response["message"] = "Usuário alterado com sucesso"

	checkResponse, err2 := json.Marshal(response)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	w.Write(checkResponse)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var deleteUser DeleteRequest

	err := json.NewDecoder(r.Body).Decode(&deleteUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	services.DeleteUserService(deleteUser.Id)

	response := make(map[string]string)
	response["message"] = "Usuário deletado com sucesso"

	checkResponse, err2 := json.Marshal(response)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	w.Write(checkResponse)
}
