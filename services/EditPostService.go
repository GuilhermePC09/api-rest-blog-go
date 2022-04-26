package services

import (
	"errors"

	"github.com/GuilhermePC09/api-rest-blog-go/repository"
)

func EditPostService(editType string, postId int64, title string, content string) error {

	postExists := repository.FindPost(postId)

	if !postExists {
		return errors.New("post dont exist")
	}

	if editType == "title" {
		repository.PostSqlUpdateTitle(postId, title)
		return nil
	}

	if editType == "content" {
		repository.PostSqlUpdateContent(postId, content)
		return nil
	} else {
		return errors.New("type arent selected")
	}
}
