/*
This package contains the blog's models and holds the database connection globally but as
an unexported value.
*/
package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

const (
	SQL_POST_BY_ID = `
		SELECT idPost, slug, title, abstract, body, date, idUser, users.name, users.email, draft, tag
		FROM posts 
		LEFT JOIN post_tags USING(idPost) 
		INNER JOIN users USING(idUser)
		WHERE idPost=?`

	SQL_POST_BY_SLUG = `
		SELECT idPost, slug, title, abstract, body, date, idUser, users.name, users.email, draft, tag 
		FROM posts 
		LEFT JOIN post_tags USING(idPost)
		INNER JOIN users USING(idUser)
		WHERE slug=?`

	SQL_POSTS_BY_USER = `
		SELECT idPost, title, slug, date, draft
		FROM posts
		WHERE idUser=?
		ORDER BY draft DESC, date DESC`

	SQL_POSTS_BY_TAG = `
		SELECT slug, title, abstract, date, idUser, users.name
		FROM posts
		INNER JOIN users USING(idUser)
		LEFT JOIN post_tags USING(idPost)
		WHERE draft=false AND post_tags.tag=?
		ORDER BY date DESC`

	SQL_ALL_POSTS = `
		SELECT slug, title, abstract, date, idUser, users.name
		FROM posts
		INNER JOIN users USING(idUser)
		WHERE draft=false
		ORDER BY date DESC LIMIT ?`

	SQL_ALL_TAGS = `SELECT DISTINCT tag FROM post_tags INNER JOIN posts USING(idPost) WHERE posts.draft=false`

	SQL_INSERT_POST = `
		INSERT INTO posts (slug, title, abstract, body, idUser, draft)
		VALUES (?, ?, ?, ?, ?, ?)`

	SQL_INSERT_TAGS = `
		INSERT IGNORE INTO post_tags (idPost, tag)
		VALUES (?, ?)`

	SQL_REMOVE_TAGS = `DELETE from post_tags WHERE idPost=?`

	SQL_DELETE_POST = `DELETE from posts WHERE idPost=?`

	SQL_UPDATE_POST = `
		UPDATE posts SET slug=?, title=?, abstract=?, body=?, idUser=?, draft=?
		WHERE idPost=?`

	SQL_USER_BY_ID = `
		SELECT name, email 
		FROM users 
		WHERE idUser=?`

	SQL_USER_AUTH = `
		SELECT name, idUser
		FROM users 
		WHERE email=? AND password=?`
)

// Creates and tests database connection
func ConnectDb(address string) {
	var err error

	db, err = sql.Open("mysql", address)
	if err != nil {
		log.Fatal("Error opening DB")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to DB")
	}
}

// Closes database connection
func CloseDb() {
	db.Close()
}
