package controllers

import (
	"fmt"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {}
func ListPosts(w http.ResponseWriter, r *http.Request)  {}
func EditPost(w http.ResponseWriter, r *http.Request)   {}
func DeletePost(w http.ResponseWriter, r *http.Request) {}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}
