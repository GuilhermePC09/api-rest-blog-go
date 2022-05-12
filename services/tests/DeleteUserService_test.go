package services

import (
	"fmt"
	"testing"

	"github.com/GuilhermePC09/api-rest-blog-go/database"
	"github.com/GuilhermePC09/api-rest-blog-go/services"
)

func TestDeleteUserService(t *testing.T) {

	database.InitDatabase()
	defer database.Db.Close()

	testId := int64(393815335927)

	test := services.DeleteUserService(testId)

	if test == 1 {
		fmt.Println("Feature com funcionamento correto")
	} else {
		if test == 0 {
			t.Fatalf("Nenhum user deletado")
		} else {
			t.Fatalf("Mais de um user deletado")
		}
	}
}
