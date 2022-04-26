package repository

import (
	"fmt"

	"github.com/GuilhermePC09/api-rest-blog-go/database"
	"github.com/GuilhermePC09/api-rest-blog-go/infra/dbconfig"
	_ "github.com/lib/pq"
)

type UserInfo struct {
	UserId       int64
	UserName     string
	UserEmail    string
	UserPassword string
}

func FindUser(email string) bool {
	var exists bool = false

	sqlStatement, Err := database.Db.Query("SELECT email FROM " + dbconfig.TableUser)
	database.CheckErr(Err)

	for sqlStatement.Next() {
		var emails dbconfig.UserTable

		Err = sqlStatement.Scan(&emails.UserEmail)
		database.CheckErr(Err)

		if emails.UserEmail == email {
			exists = true
		}
	}
	return exists
}

func UserSqlSelect() []UserInfo {

	userList := make([]UserInfo, 0)

	sqlStatement, Err := database.Db.Query("SELECT userid, name, email, password FROM " + dbconfig.TableUser)
	database.CheckErr(Err)

	for sqlStatement.Next() {
		var users dbconfig.UserTable

		Err = sqlStatement.Scan(&users.IdUser, &users.UserName, &users.UserEmail, &users.UserPassword)
		database.CheckErr(Err)

		createUser := UserInfo{
			UserId:       users.IdUser,
			UserName:     users.UserName,
			UserEmail:    users.UserEmail,
			UserPassword: users.UserPassword,
		}

		userList = append(userList, createUser)
	}
	return userList
}

func UserSqlSelectId(id int64) (int64, string, string, string) {
	var user dbconfig.UserTable

	sqlStatement := fmt.Sprintf("SELECT postid, title, content FROM %s where userid = $1", dbconfig.TableUser)
	Err := database.Db.QueryRow(sqlStatement, id).Scan(&user.IdUser, &user.UserName, &user.UserEmail, &user.UserPassword)
	database.CheckErr(Err)

	return user.IdUser, user.UserName, user.UserEmail, user.UserPassword
}

func UserSqlInsert(userId int64, userName string, userEmail string, userPassword string) int64 {

	sqlStatement := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2, $3, $4)", dbconfig.TableUser)

	insert, Err := database.Db.Prepare(sqlStatement)
	database.CheckErr(Err)

	result, Err := insert.Exec(userId, userName, userEmail, userPassword)
	database.CheckErr(Err)

	affect, Err := result.RowsAffected()
	database.CheckErr(Err)

	return affect
}

func UserSqlUpdatePasword(userId int64, password string) int64 {
	sqlStatement := fmt.Sprintf("UPDATE %s SET password=$1 where userid=$2", dbconfig.TableUser)

	update, Err := database.Db.Prepare(sqlStatement)
	database.CheckErr(Err)

	result, Err := update.Exec(password, userId)
	database.CheckErr(Err)

	affect, Err := result.RowsAffected()
	database.CheckErr(Err)

	return affect
}

func UserSqlUpdateName(userId int64, name string) int64 {
	sqlStatement := fmt.Sprintf("UPDATE %s SET name=$1 where userid=$2", dbconfig.TableUser)

	update, Err := database.Db.Prepare(sqlStatement)
	database.CheckErr(Err)

	result, Err := update.Exec(name, userId)
	database.CheckErr(Err)

	affect, Err := result.RowsAffected()
	database.CheckErr(Err)

	return affect
}

func UserSqlUpdateEmail(userId int64, email string) int64 {
	sqlStatement := fmt.Sprintf("UPDATE %s SET email=$1 where userid=$2", dbconfig.TableUser)

	update, Err := database.Db.Prepare(sqlStatement)
	database.CheckErr(Err)

	result, Err := update.Exec(email, userId)
	database.CheckErr(Err)

	affect, Err := result.RowsAffected()
	database.CheckErr(Err)

	return affect
}

func UserSqlDelete(userId int64) int64 {

	sqlStatement := fmt.Sprintf("delete from %s where userid=$1", dbconfig.TableUser)

	delete, Err := database.Db.Prepare(sqlStatement)
	database.CheckErr(Err)

	result, Err := delete.Exec(userId)
	database.CheckErr(Err)

	affect, Err := result.RowsAffected()
	database.CheckErr(Err)

	return affect
}
