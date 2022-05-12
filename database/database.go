package database

import (
	"database/sql"
	"fmt"

	"github.com/GuilhermePC09/api-rest-blog-go/infra/dbconfig"
)

var (
	Db  *sql.DB
	Err error
)

func CheckErr(errp error) {
	if errp != nil {
		panic(errp.Error())
	}
}

func InitDatabase() {
	fmt.Printf("Accessing %s ... ", dbconfig.DbName)

	Db, Err = sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	if Err != nil {
		panic(Err.Error())
	} else {
		fmt.Println("Connected!")
	}

	return
}
