package services

import (
	"fmt"
	"testing"

	"github.com/GuilhermePC09/api-rest-blog-go/database"
	"github.com/GuilhermePC09/api-rest-blog-go/services"
)

func TestListUsersServece(t *testing.T) {

	database.InitDatabase()
	defer database.Db.Close()

	ListOfUsersTest := services.ListUsersService()

	if ListOfUsersTest != nil {
		fmt.Println("Feature com funcionamento correto")
	} else {
		t.Fatalf("Lista não encontrada")
	}
}
