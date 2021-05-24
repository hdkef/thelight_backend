package driver

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"thelight/models"
)

//DBArticleGetAll
func DBArticleGetAll(db *sql.DB, payload *models.ArticleFromClient) ([]models.Article, error) {
	ctx := context.Background()

	var limit int64 = 6
	offset := (payload.Page - 1) * limit

	var articles []models.Article

	rows, err := db.QueryContext(
		ctx,
		"SELECT articles.ID, articles.Title, articles.Date, articles.Body, articles.Tag, articles.ImageURL, users.ID, users.Name, users.AvatarURL, users.Bio FROM articles FULL JOIN users ON users.ID = articles.User_Ref LIMIT $1 OFFSET $2",
		limit, offset,
	)
	if err != nil {
		return articles, err
	}

	for rows.Next() {
		var tmp models.Article
		var tagstring string
		rows.Scan(&tmp.ID, &tmp.Title, &tmp.Date, &tmp.Body, &tagstring, &tmp.ImageURL, &tmp.WriterInfo.ID, &tmp.WriterInfo.Name, &tmp.WriterInfo.AvatarURL, &tmp.WriterInfo.Bio)
		tmp.Tag = strings.Split(tagstring, ",")
		articles = append(articles, tmp)
	}

	return articles, nil
}

//DBArticleGetOne
func DBArticleGetOne(db *sql.DB, payload *models.ArticleFromClient) (models.Article, error) {
	ctx := context.Background()

	var article models.Article
	var tagstring string

	err := db.QueryRowContext(
		ctx,
		"SELECT articles.ID, articles.Title, articles.Date, articles.Body, articles.Tag, articles.ImageURL, users.ID, users.Name, users.AvatarURL, users.Bio FROM articles JOIN users ON users.ID = articles.User_Ref WHERE articles.ID=$1",
		payload.ID,
	).Scan(&article.ID, &article.Title, &article.Date, &article.Body, &tagstring, &article.ImageURL, &article.WriterInfo.ID, &article.WriterInfo.Name, &article.WriterInfo.AvatarURL, &article.WriterInfo.Bio)
	if err != nil {
		return article, err
	}

	article.Tag = strings.Split(tagstring, ",")

	return article, nil
}

//DBArticleSearch
func DBArticleSearch(db *sql.DB, payload *models.ArticleFromClient) ([]models.Article, error) {
	ctx := context.Background()

	var limit int64 = 6
	offset := (payload.Page - 1) * limit

	var articles []models.Article

	var (
		rows *sql.Rows
		err  error
	)

	if payload.Filter == "Tag" {
		rows, err = db.QueryContext(
			ctx,
			"SELECT articles.ID, articles.Title, articles.Date, articles.Body, articles.Tag, articles.ImageURL, users.ID, users.Name, users.AvatarURL, users.Bio FROM articles JOIN users ON users.ID = articles.User_Ref WHERE articles.Tag LIKE '%' || $1 || '%' LIMIT $2 OFFSET $3",
			payload.Key, limit, offset,
		)
	} else if payload.Filter == "Title" {
		rows, err = db.QueryContext(
			ctx,
			"SELECT articles.ID, articles.Title, articles.Date, articles.Body, articles.Tag, articles.ImageURL, users.ID, users.Name, users.AvatarURL, users.Bio FROM articles JOIN users ON users.ID = articles.User_Ref WHERE articles.Title LIKE '%' || $1 || '%' LIMIT $2 OFFSET $3",
			payload.Key, limit, offset,
		)
	} else {
		return []models.Article{}, errors.New("NO FILTER METHOD FOUND")
	}

	if err != nil {
		return articles, err
	}

	for rows.Next() {
		var tmp models.Article
		var tagstring string
		rows.Scan(&tmp.ID, &tmp.Title, &tmp.Date, &tmp.Body, &tagstring, &tmp.ImageURL, &tmp.WriterInfo.ID, &tmp.WriterInfo.Name, &tmp.WriterInfo.AvatarURL, &tmp.WriterInfo.Bio)
		tmp.Tag = strings.Split(tagstring, ",")
		articles = append(articles, tmp)
	}

	return articles, nil
}
