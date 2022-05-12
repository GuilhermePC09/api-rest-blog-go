package repository

import (
	"fmt"

	"github.com/GuilhermePC09/api-rest-blog-go/database"
	"github.com/GuilhermePC09/api-rest-blog-go/infra/dbconfig"
	_ "github.com/lib/pq"
)

type PostInfo interface{}

func FindPost(id string) bool {

	var exists bool = false

	sqlStatement, Err := database.Db.Query("SELECT postid FROM " + dbconfig.TablePost)
	database.CheckErr(Err)

	for sqlStatement.Next() {
		var posts dbconfig.PostTable

		Err = sqlStatement.Scan(&posts.IdPost)
		database.CheckErr(Err)

		if posts.IdPost == id {
			exists = true
		}
	}
	return exists
}

func PostSqlSelect() []PostInfo {

	postList := make([]PostInfo, 0)

	sqlStatement, Err := database.Db.Query("SELECT postid, userid, title, content, datetime FROM " + dbconfig.TablePost)
	database.CheckErr(Err)

	for sqlStatement.Next() {
		var posts dbconfig.PostTable

		Err = sqlStatement.Scan(&posts.IdPost, &posts.IdUser, &posts.Title, &posts.Content, &posts.DateTime)
		database.CheckErr(Err)

		postList = append(postList,
			posts.IdUser,
			posts.IdPost,
			posts.Title,
			posts.Content,
			posts.DateTime)
	}
	return postList
}

func PostSqlSelectId(id string) []PostInfo {
	var post dbconfig.PostTable
	postList := make([]PostInfo, 0)

	sqlStatement := fmt.Sprintf("SELECT postid, title, content FROM %s where postid = $1", dbconfig.TablePost)
	Err := database.Db.QueryRow(sqlStatement, id).Scan(&post.IdUser, &post.IdPost, &post.Title, &post.Content, &post.DateTime)

	database.CheckErr(Err)

	postList = append(postList,
		post.IdUser,
		post.IdPost,
		post.Title,
		post.Content,
		post.DateTime)

	return postList
}

func PostSqlInsert(postId string, userId int64, title string, content string, date string) int64 {

	sqlStatement := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2, $3, $4, $5)", dbconfig.TablePost)

	insert, Err := database.Db.Prepare(sqlStatement)
	database.CheckErr(Err)

	result, Err := insert.Exec(postId, userId, title, content, date)
	database.CheckErr(Err)

	affect, Err := result.RowsAffected()
	database.CheckErr(Err)

	return affect
}

func PostSqlUpdateContent(postId string, content string) int64 {
	sqlStatement := fmt.Sprintf("UPDATE %s SET content=$1 where postid=$2", dbconfig.TablePost)

	update, Err := database.Db.Prepare(sqlStatement)
	database.CheckErr(Err)

	result, Err := update.Exec(content, postId)
	database.CheckErr(Err)

	affect, Err := result.RowsAffected()
	database.CheckErr(Err)

	return affect
}

func PostSqlUpdateTitle(postId string, title string) int64 {
	sqlStatement := fmt.Sprintf("UPDATE %s SET title=$1 where postid=$2", dbconfig.TablePost)

	update, Err := database.Db.Prepare(sqlStatement)
	database.CheckErr(Err)

	result, Err := update.Exec(title, postId)
	database.CheckErr(Err)

	affect, Err := result.RowsAffected()
	database.CheckErr(Err)

	return affect
}

func PostSqlDelete(postId string) int64 {

	sqlStatement := fmt.Sprintf("delete from %s where postid=$1", dbconfig.TablePost)

	delete, Err := database.Db.Prepare(sqlStatement)
	database.CheckErr(Err)

	result, Err := delete.Exec(postId)
	database.CheckErr(Err)

	affect, Err := result.RowsAffected()
	database.CheckErr(Err)

	return affect
}
