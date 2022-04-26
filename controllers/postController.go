package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermePC09/api-rest-blog-go/services"
)

type PostRequest struct {
	Title   string
	Content string
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
func ListPosts(w http.ResponseWriter, r *http.Request)  {}
func EditPost(w http.ResponseWriter, r *http.Request)   {}
func DeletePost(w http.ResponseWriter, r *http.Request) {}
