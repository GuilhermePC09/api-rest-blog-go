package services

import (
	"fmt"
	"testing"

	"github.com/GuilhermePC09/api-rest-blog-go/database"
	"github.com/GuilhermePC09/api-rest-blog-go/services"
)

type trueResultInfoEditPost interface{}

func TestEditPostService(t *testing.T) {

	database.InitDatabase()
	defer database.Db.Close()

	testType := "title"
	testPostId := "7a9c7c15-c49e-4cb7-b0be-47811f5ebd6e"
	testTitle := "Título de teste"
	testContent := "Esse é um texto de teste"

	expectedResult := []trueResultInfoEditPost{testPostId, testTitle, testContent}

	test, Err := services.EditPostService(testType, testPostId, testTitle, testContent)

	if Err != nil {
		panic(Err.Error())
	}

	if test[1] != expectedResult[0] {
		t.Fatalf("Valores encontrados diferentes do esperado")
	} else {
		fmt.Println("Tudo ocorreu como esperado!!!")
	}
}
