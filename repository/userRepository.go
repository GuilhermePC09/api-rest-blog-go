package repository

import (
	"database/sql"
	"fmt"

	"github.com/GuilhermePC09/api-rest-blog-go/gopostgres/dbconfig"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	erru error
)

func checkErrUser(erru error) {
	if erru != nil {
		panic(erru.Error())
	}
}

func UserSqlSelect() {
	sqlStatement, erru := db.Query("SELECT userid, name, email, password FROM " + dbconfig.TableName)
	checkErrUser(erru)

	for sqlStatement.Next() {
		var users dbconfig.UserTable

		erru = sqlStatement.Scan(&users.IdUser, &users.UserName, &users.UserEmail, &users.UserPassword)
		checkErrUser(erru)

		fmt.Printf("%d\t%s\t%s \n", users.IdUser, users.UserName, users.UserEmail)
	}
}

func UserSqlSelectId(id int) (int, string, string, string) {
	var user dbconfig.UserTable

	sqlStatement := fmt.Sprintf("SELECT postid, title, content FROM %s where postid = $1", dbconfig.TableName)
	erru = db.QueryRow(sqlStatement, id).Scan(&user.IdUser, &user.UserName, &user.UserEmail, &user.UserPassword)
	checkErrUser(erru)

	return user.IdUser, user.UserName, user.UserEmail, user.UserPassword
}

func UserSqlInsert(userId int, userName string, userEmail string, userPassword string) int64 {

	sqlStatement := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2, $3, $4)", dbconfig.TableName)

	insert, erru := db.Prepare(sqlStatement)
	checkErrUser(erru)

	result, erru := insert.Exec(userId, userName, userEmail, userPassword)
	checkErrUser(erru)

	affect, erru := result.RowsAffected()
	checkErrUser(erru)

	return affect
}

func UserSqlUpdatePasword(userId int, password string) int64 {
	sqlStatement := fmt.Sprintf("UPDATE %s SET password=$1 where postid=$2", dbconfig.TableName)

	update, erru := db.Prepare(sqlStatement)
	checkErrUser(erru)

	result, erru := update.Exec(password, userId)
	checkErrUser(erru)

	affect, erru := result.RowsAffected()
	checkErrUser(erru)

	return affect
}

func UserSqlUpdateName(userId int, name string) int64 {
	sqlStatement := fmt.Sprintf("UPDATE %s SET name=$1 where postid=$2", dbconfig.TableName)

	update, erru := db.Prepare(sqlStatement)
	checkErrUser(erru)

	result, erru := update.Exec(name, userId)
	checkErrUser(erru)

	affect, erru := result.RowsAffected()
	checkErrUser(erru)

	return affect
}

func UserSqlUpdateEmail(userId int, email string) int64 {
	sqlStatement := fmt.Sprintf("UPDATE %s SET email=$1 where postid=$2", dbconfig.TableName)

	update, erru := db.Prepare(sqlStatement)
	checkErrUser(erru)

	result, erru := update.Exec(email, userId)
	checkErrUser(erru)

	affect, erru := result.RowsAffected()
	checkErrUser(erru)

	return affect
}

func UserSqlDelete(userId int) int64 {

	sqlStatement := fmt.Sprintf("delete from %s where userid=$1", dbconfig.TableName)

	delete, erru := db.Prepare(sqlStatement)
	checkErrUser(erru)

	result, erru := delete.Exec(userId)
	checkErrUser(erru)

	affect, erru := result.RowsAffected()
	checkErrUser(erru)

	return affect
}
