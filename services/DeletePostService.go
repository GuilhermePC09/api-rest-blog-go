package services

import "github.com/GuilhermePC09/api-rest-blog-go/repository"

func DeletePostService(id int64) {

	repository.PostSqlDelete(id)

}
