package repository

import (
	"database/sql"
	"fmt"

	"github.com/GuilhermePC09/api-rest-blog-go/gopostgres/dbconfig"
	_ "github.com/lib/pq"
)

var (
	Db   *sql.DB
	errp error
)

func checkErrPosts(errp error) {
	if errp != nil {
		panic(errp.Error())
	}
}

func PostSqlSelect() {
	sqlStatement, errp := Db.Query("SELECT postid, title, Content FROM " + dbconfig.TableName)
	checkErrPosts(errp)

	for sqlStatement.Next() {
		var posts dbconfig.PostTable

		errp = sqlStatement.Scan(&posts.IdPost, &posts.Title, &posts.Content)
		checkErrPosts(errp)

		fmt.Printf("%d\t%s\t%s \n", posts.IdPost, posts.Title, posts.Content)
	}
}

func PostSqlSelectId(id int) (int, string, []byte) {
	var post dbconfig.PostTable

	sqlStatement := fmt.Sprintf("SELECT postid, title, content FROM %s where postid = $1", dbconfig.TableName)
	errp = Db.QueryRow(sqlStatement, id).Scan(&post.IdPost, &post.Title, &post.Content)
	checkErrPosts(errp)

	return post.IdPost, post.Title, post.Content
}

func PostSqlInsert(postId int, userId int, title string, content []byte) int64 {

	sqlStatement := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2, $3, $4)", dbconfig.TableName)

	insert, errp := Db.Prepare(sqlStatement)
	checkErrPosts(errp)

	result, errp := insert.Exec(postId, userId, title, content)
	checkErrPosts(errp)

	affect, errp := result.RowsAffected()
	checkErrPosts(errp)

	return affect
}

func PostSqlUpdateContent(postId int, content []byte) int64 {
	sqlStatement := fmt.Sprintf("UPDATE %s SET content=$1 where postid=$2", dbconfig.TableName)

	update, errp := Db.Prepare(sqlStatement)
	checkErrPosts(errp)

	result, errp := update.Exec(content, postId)
	checkErrPosts(errp)

	affect, errp := result.RowsAffected()
	checkErrPosts(errp)

	return affect
}

func PostSqlUpdateTitle(postId int, title string) int64 {
	sqlStatement := fmt.Sprintf("UPDATE %s SET title=$1 where postid=$2", dbconfig.TableName)

	update, errp := Db.Prepare(sqlStatement)
	checkErrPosts(errp)

	result, errp := update.Exec(title, postId)
	checkErrPosts(errp)

	affect, errp := result.RowsAffected()
	checkErrPosts(errp)

	return affect
}

func PostSqlDelete(postId int) int64 {

	sqlStatement := fmt.Sprintf("delete from %s where postid=$1", dbconfig.TableName)

	delete, errp := Db.Prepare(sqlStatement)
	checkErrPosts(errp)

	result, errp := delete.Exec(postId)
	checkErrPosts(errp)

	affect, errp := result.RowsAffected()
	checkErrPosts(errp)

	return affect
}
