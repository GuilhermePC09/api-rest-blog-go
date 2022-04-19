package main

import (
	"database/sql"
	"fmt"

	"github.com/GuilhermePC09/api-rest-blog-go/gopostgres/dbconfig"
	"github.com/GuilhermePC09/api-rest-blog-go/repository"
	"github.com/GuilhermePC09/api-rest-blog-go/routes"
	_ "github.com/lib/pq"
)

var (
	err error
)

func main() {
	fmt.Printf("Accessing %s ... ", dbconfig.DbName)

	repository.Db, err = sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer repository.Db.Close()

	routes.HandleRequest()
}
