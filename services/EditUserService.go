package services

import (
	"errors"

	"github.com/GuilhermePC09/api-rest-blog-go/repository"
)

func EditUserService(id int64, email string, editType string, editInfo string) error {

	userExists := repository.FindUser(email)

	if !userExists {
		return errors.New("user dont exist")
	}

	if editType == "name" {
		repository.UserSqlUpdateName(id, editInfo)
		return nil
	}

	if editType == "email" {
		repository.UserSqlUpdateEmail(id, editInfo)
		return nil
	}

	if editType == "password" {
		repository.UserSqlUpdatePasword(id, editInfo)
		return nil
	} else {
		return errors.New("type arent selected")
	}
}
