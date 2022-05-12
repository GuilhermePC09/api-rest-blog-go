package services

import (
	"fmt"
	"testing"

	"github.com/GuilhermePC09/api-rest-blog-go/database"
	"github.com/GuilhermePC09/api-rest-blog-go/services"
)

func TestDeletePostService(t *testing.T) {

	database.InitDatabase()
	defer database.Db.Close()

	testId := "b6c0dad0-0b79-45ac-b6ba-6dcbb0307c72"

	test := services.DeletePostService(testId)

	if test == 1 {
		fmt.Println("Feature com funcionamento correto")
	} else {
		if test == 0 {
			t.Fatalf("Nenhum post deletado")
		} else {
			t.Fatalf("Mais de um post deletado")
		}
	}
}
