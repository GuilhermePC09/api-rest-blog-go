package services

import (
	"errors"
	"time"

	"github.com/GuilhermePC09/api-rest-blog-go/repository"
)

type IPostRequest struct {
	UserId  int64
	PostId  int64
	Title   string
	Content string
}

func CreatePostService(userId int64, title string, content string) (IPostRequest, error) {

	if title == "" || content == "" {
		return IPostRequest{}, errors.New("missing information")
	}

	postId := time.Now().UnixNano() / (1 << 44)

	repository.PostSqlInsert(postId, userId, title, content)

	return IPostRequest{
		UserId:  userId,
		PostId:  postId,
		Title:   title,
		Content: content,
	}, nil
}
