package services

import (
	"github.com/GuilhermePC09/api-rest-blog-go/repository"
)

func ListUsersService() []repository.UserInfo {
	ListOfUsers := repository.UserSqlSelect()

	return ListOfUsers
}
