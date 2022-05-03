package repository

import (
	"fmt"

	"github.com/GuilhermePC09/api-rest-blog-go/database"
	"github.com/GuilhermePC09/api-rest-blog-go/infra/dbconfig"
	_ "github.com/lib/pq"
)

type PostInfo struct {
	UserId  int64
	PostId  int64
	Title   string
	Content string
}

func FindPost(id int64) bool {

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

	sqlStatement, Err := database.Db.Query("SELECT postid, userid, title, content FROM " + dbconfig.TablePost)
	database.CheckErr(Err)

	for sqlStatement.Next() {
		var posts dbconfig.PostTable

		Err = sqlStatement.Scan(&posts.IdPost, &posts.IdUser, &posts.Title, &posts.Content)
		database.CheckErr(Err)

		createPost := PostInfo{
			UserId:  posts.IdUser,
			PostId:  posts.IdPost,
			Title:   posts.Title,
			Content: posts.Content,
		}

		postList = append(postList, createPost)
	}
	return postList
}

func PostSqlSelectId(id int64) []PostInfo {
	var post dbconfig.PostTable
	postList := make([]PostInfo, 0)

	sqlStatement := fmt.Sprintf("SELECT postid, title, content FROM %s where postid = $1", dbconfig.TablePost)
	Err := database.Db.QueryRow(sqlStatement, id).Scan(&post.IdUser, &post.IdPost, &post.Title, &post.Content)

	database.CheckErr(Err)

	createPost := PostInfo{
		UserId:  post.IdUser,
		PostId:  post.IdPost,
		Title:   post.Title,
		Content: post.Content,
	}

	postList = append(postList, createPost)

	return postList
}

func PostSqlInsert(postId int64, userId int64, title string, content string) int64 {

	sqlStatement := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2, $3, $4)", dbconfig.TablePost)

	insert, Err := database.Db.Prepare(sqlStatement)
	database.CheckErr(Err)

	result, Err := insert.Exec(postId, userId, title, content)
	database.CheckErr(Err)

	affect, Err := result.RowsAffected()
	database.CheckErr(Err)

	return affect
}

func PostSqlUpdateContent(postId int64, content string) int64 {
	sqlStatement := fmt.Sprintf("UPDATE %s SET content=$1 where postid=$2", dbconfig.TablePost)

	update, Err := database.Db.Prepare(sqlStatement)
	database.CheckErr(Err)

	result, Err := update.Exec(content, postId)
	database.CheckErr(Err)

	affect, Err := result.RowsAffected()
	database.CheckErr(Err)

	return affect
}

func PostSqlUpdateTitle(postId int64, title string) int64 {
	sqlStatement := fmt.Sprintf("UPDATE %s SET title=$1 where postid=$2", dbconfig.TablePost)

	update, Err := database.Db.Prepare(sqlStatement)
	database.CheckErr(Err)

	result, Err := update.Exec(title, postId)
	database.CheckErr(Err)

	affect, Err := result.RowsAffected()
	database.CheckErr(Err)

	return affect
}

func PostSqlDelete(postId int64) int64 {

	sqlStatement := fmt.Sprintf("delete from %s where postid=$1", dbconfig.TablePost)

	delete, Err := database.Db.Prepare(sqlStatement)
	database.CheckErr(Err)

	result, Err := delete.Exec(postId)
	database.CheckErr(Err)

	affect, Err := result.RowsAffected()
	database.CheckErr(Err)

	return affect
}
