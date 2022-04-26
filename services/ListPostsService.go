package services

import "github.com/GuilhermePC09/api-rest-blog-go/repository"

func ListPostsService() []*repository.PostInfo {
	ListOfPosts := repository.PostSqlSelect()

	return ListOfPosts
}
