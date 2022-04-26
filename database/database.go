package database

import (
	"database/sql"
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
