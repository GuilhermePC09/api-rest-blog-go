package main

import (
	"github.com/GuilhermePC09/api-rest-blog-go/database"
	"github.com/GuilhermePC09/api-rest-blog-go/routes"
	_ "github.com/lib/pq"
)

var (
	err error
)

func main() {

	database.InitDatabase()
	defer database.Db.Close()

	routes.HandleRequest()
}
