package dbconfig

import (
	"fmt"
	"time"
)

type UserTable struct {
	IdUser       int
	UserName     string
	UserEmail    string
	UserPassword string
}

type PostTable struct {
	IdUser   int
	IdPost   int
	Title    string
	Content  []byte
	DateTime time.Time
}

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "Gui090900!"
const DbName = "BlogDatabase"
const TableName = "posts"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
