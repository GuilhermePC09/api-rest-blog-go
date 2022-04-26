package main

import (
	"database/sql"
	"fmt"

	"github.com/GuilhermePC09/api-rest-blog-go/database"
	"github.com/GuilhermePC09/api-rest-blog-go/infra/dbconfig"
	"github.com/GuilhermePC09/api-rest-blog-go/routes"
	_ "github.com/lib/pq"
)

var (
	err error
)

func main() {
	fmt.Printf("Accessing %s ... ", dbconfig.DbName)

	database.Db, err = sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer database.Db.Close()

	routes.HandleRequest()
}
