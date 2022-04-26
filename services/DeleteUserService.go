package services

import "github.com/GuilhermePC09/api-rest-blog-go/repository"

func DeleteUserService(id int64) {

	repository.UserSqlDelete(id)
}
