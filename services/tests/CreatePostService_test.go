package services

import (
	"fmt"
	"testing"

	"github.com/GuilhermePC09/api-rest-blog-go/database"
	"github.com/GuilhermePC09/api-rest-blog-go/repository"
	"github.com/GuilhermePC09/api-rest-blog-go/services"
)

type trueResultInfoPost interface{}

func TestCreatePostService(t *testing.T) {

	database.InitDatabase()
	defer database.Db.Close()

	var testUserId int64 = 424623423
	testTitle := "Título de teste"
	testContent := "Esse é um texto de teste"

	expectedResult := []trueResultInfoPost{424623423, "Título de teste", "Esse é um texto de teste"}

	test, err := services.CreatePostService(testUserId, testTitle, testContent)

	if err != nil {
		panic(err.Error())
	}

	if test[2] != expectedResult[1] || test[3] != expectedResult[2] {
		t.Fatalf("Valores encontrados diferentes do esperado")
		id := fmt.Sprint(test[0])
		repository.PostSqlDelete(string(id))
	} else {
		fmt.Println("Tudo ocorreu como esperado!!!")
	}
}
