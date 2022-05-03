package services

import (
	"errors"

	"github.com/GuilhermePC09/api-rest-blog-go/repository"
)

func EditUserService(editType string, userId int64, email string, name string, password string) ([]repository.UserInfo, error) {

	userAlreadyExists := repository.FindUserId(userId)

	if !userAlreadyExists {
		return nil, errors.New("User dont exist")
	}
	if editType == "name" {
		repository.UserSqlUpdateName(userId, name)
		editedUser := repository.UserSqlSelectId(userId)
		return editedUser, nil
	}

	if editType == "email" {
		repository.UserSqlUpdateEmail(userId, email)
		editedUser := repository.UserSqlSelectId(userId)
		return editedUser, nil
	}

	if editType == "password" {
		repository.UserSqlUpdatePasword(userId, password)
		editedUser := repository.UserSqlSelectId(userId)
		return editedUser, nil
	} else {
		return nil, errors.New("Type arent selected")
	}
}
