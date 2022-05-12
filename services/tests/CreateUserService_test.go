package services

import (
	"fmt"
	"testing"

	"github.com/GuilhermePC09/api-rest-blog-go/database"
	"github.com/GuilhermePC09/api-rest-blog-go/repository"
	"github.com/GuilhermePC09/api-rest-blog-go/services"
)

type trueResultInfoUser interface{}

func TestCreateUserService(t *testing.T) {

	database.InitDatabase()
	defer database.Db.Close()

	userName := "Nome de teste"
	userEmail := "Email de teste"
	userPassword := "Senha de teste"

	expectedResult := []trueResultInfoUser{"Nome de teste", "Email de teste", "Senha de teste"}

	test, err := services.CreateUserService(userName, userEmail, userPassword)

	if err != nil {
		panic(err.Error())
	}

	if test[1] != expectedResult[0] || test[2] != expectedResult[1] {
		t.Fatalf("Valores encontrados diferentes do esperado")
		id := fmt.Sprint(test[1])
		repository.PostSqlDelete(string(id))
	} else {
		fmt.Println("Tudo ocorreu como esperado!!!")
	}
}
