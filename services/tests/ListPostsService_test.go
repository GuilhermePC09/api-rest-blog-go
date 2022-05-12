package services

import (
	"fmt"
	"testing"

	"github.com/GuilhermePC09/api-rest-blog-go/database"
	"github.com/GuilhermePC09/api-rest-blog-go/services"
)

func TestListPostsServece(t *testing.T) {

	database.InitDatabase()
	defer database.Db.Close()

	ListOfPostsTest := services.ListPostsService()

	if ListOfPostsTest != nil {
		fmt.Println("Feature com funcionamento correto")
	} else {
		t.Fatalf("Lista não encontrada")
	}
}
