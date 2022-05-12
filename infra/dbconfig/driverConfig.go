package dbconfig

import (
	"fmt"
)

type UserTable struct {
	IdUser       int64
	UserName     string
	UserEmail    string
	UserPassword string
}

type PostTable struct {
	IdUser   int64
	IdPost   string
	Title    string
	Content  string
	DateTime string
}

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "olaisaac"
const DbName = "BlogDatabase"
const TablePost = "posts"
const TableUser = "users"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
