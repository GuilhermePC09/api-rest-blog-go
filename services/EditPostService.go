package services

import (
	"errors"

	"github.com/GuilhermePC09/api-rest-blog-go/repository"
)

type PostInfo interface{}

func EditPostService(editType string, postId string, title string, content string) ([]repository.PostInfo, error) {

	postExists := repository.FindPost(postId)

	if !postExists {
		return nil, errors.New("post dont exist")
	}

	if editType == "title" {
		repository.PostSqlUpdateTitle(postId, title)
		editedPost := repository.PostSqlSelectId(postId)
		return editedPost, nil
	}

	if editType == "content" {
		repository.PostSqlUpdateContent(postId, content)
		editedPost := repository.PostSqlSelectId(postId)
		return editedPost, nil
	} else {
		return nil, errors.New("type arent selected")
	}
}
