package services

import (
	"errors"
	"time"

	"github.com/GuilhermePC09/api-rest-blog-go/repository"
	"github.com/google/uuid"
)

type IPostRequest struct {
	UserId   int64
	PostId   string
	Title    string
	Content  string
	Datetime string
}

type IPostRequestTest interface{}

func CreatePostService(userId int64, title string, content string) ([]IPostRequestTest, error) {

	PostList := make([]IPostRequestTest, 0)

	if title == "" || content == "" {
		return nil, errors.New("missing information")
	}

	PostId := uuid.New()
	DateTime := time.Now().String()

	repository.PostSqlInsert(PostId.String(), userId, title, content, DateTime)

	PostList = append(PostList, userId, PostId.String(), title, content, DateTime)

	return PostList, nil
}
