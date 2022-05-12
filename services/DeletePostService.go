package services

import "github.com/GuilhermePC09/api-rest-blog-go/repository"

func DeletePostService(id string) int64 {

	checkDelete := repository.PostSqlDelete(id)

	return checkDelete
}
