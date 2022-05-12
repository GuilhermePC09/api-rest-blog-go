package services

import (
	"errors"
	"time"

	"github.com/GuilhermePC09/api-rest-blog-go/repository"
)

type IUserRequest struct {
	UserId       int64
	UserName     string
	UserEmail    string
	UserPassword string
}

type IUserRequestTest interface{}

func CreateUserService(userName string, userEmail string, userPassword string) ([]IUserRequestTest, error) {

	if userName == "" || userEmail == "" || userPassword == "" {
		return nil, errors.New("missing information")
	}

	userAlreadyExists := repository.FindUserEmail(userEmail)

	if userAlreadyExists {
		return nil, errors.New("user already exists")
	}

	UserList := make([]IUserRequestTest, 0)
	userId := time.Now().UnixNano() / (1 << 22)

	repository.UserSqlInsert(userId, userName, userEmail, userPassword)

	UserList = append(UserList, userId, userName, userEmail, userPassword)
	return UserList,
		nil
}
