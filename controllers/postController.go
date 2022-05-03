package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GuilhermePC09/api-rest-blog-go/services"
	"github.com/gorilla/mux"
)

type PostRequest struct {
	Id      int64
	IdUser  int64
	Title   string
	Content string
}

type EditPostRequest struct {
	Id      int64
	Title   string
	Content string
	Type    string
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post PostRequest

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	checkedPost, err2 := services.CreatePostService(post.IdUser, post.Title, post.Content)

	if err2 != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(checkedPost)
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	checkedPostList := services.ListPostsService()

	jsonPosts, err := json.MarshalIndent(checkedPostList, "", " ")

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

	edit, err2 := services.EditPostService(editPost.Type, editPost.Id, editPost.Title, editPost.Content)

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
func DeletePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["postId"]

	convId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	services.DeletePostService(convId)

	response := make(map[string]string)
	response["message"] = "Post deletado com sucesso"

	checkResponse, err2 := json.Marshal(response)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(checkResponse)
}
