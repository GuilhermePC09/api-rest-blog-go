package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GuilhermePC09/api-rest-blog-go/services"
)

type PostRequest struct {
	Id      int64
	Title   string
	Content string
}

type EditPostRequest struct {
	Id      int64
	Title   string
	Content string
	Type    string
}

type DeletePostRequest struct {
	Id int64
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post PostRequest

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	checkedPost, err2 := services.CreatePostService(0, post.Title, post.Content)

	if err2 != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(checkedPost)
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	checkedPostList := services.ListPostsService()

	jsonPosts, err := json.MarshalIndent(checkedPostList, "", "")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(jsonPosts))
}

func EditPost(w http.ResponseWriter, r *http.Request) {

	var editPost EditPostRequest

	err := json.NewDecoder(r.Body).Decode(&editPost)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	services.EditPostService(editPost.Type, editPost.Id, editPost.Title, editPost.Content)
	response := make(map[string]string)
	response["message"] = "Usuário alterado com sucesso"

	checkedResponse, err2 := json.Marshal(response)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	w.Write(checkedResponse)
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	var deleteUser DeletePostRequest

	err := json.NewDecoder(r.Body).Decode(&deleteUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := make(map[string]string)
	response["message"] = "Post deletado com sucesso"

	checkResponse, err2 := json.Marshal(response)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(checkResponse)
}
