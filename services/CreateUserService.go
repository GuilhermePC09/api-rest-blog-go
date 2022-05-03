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

func CreateUserService(userName string, userEmail string, userPassword string) (IUserRequest, error) {

	if userName == "" || userEmail == "" || userPassword == "" {
		return IUserRequest{}, errors.New("missing information")
	}

	userAlreadyExists := repository.FindUserEmail(userEmail)

	if userAlreadyExists {
		return IUserRequest{}, errors.New("user already exists")
	}

	userId := time.Now().UnixNano() / (1 << 22)
	repository.UserSqlInsert(userId, userName, userEmail, userPassword)

	return IUserRequest{
			UserId:       userId,
			UserName:     userName,
			UserEmail:    userEmail,
			UserPassword: userPassword},
		nil
}
