package dbconfig

import (
	"fmt"
	"time"
)

type UserTable struct {
	IdUser       int64
	UserName     string
	UserEmail    string
	UserPassword string
}

type PostTable struct {
	IdUser   int64
	IdPost   int64
	Title    string
	Content  string
	DateTime time.Time
}

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "Gui090900!"
const DbName = "BlogDatabase"
const TablePost = "posts"
const TableUser = "users"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
