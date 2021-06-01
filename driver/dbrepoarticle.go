package driver

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"thelight/models"
)

//DBArticleGetAll will return all articles paginated by LastID
func DBArticleGetAll(db *sql.DB, payload *models.ArticleFromClient) ([]models.Article, error) {
	ctx := context.Background()

	var limit int64 = 6

	var articles []models.Article

	rows, err := db.QueryContext(
		ctx,
		"SELECT articles.ID, articles.Title, articles.Date, articles.Body, articles.Tag, articles.ImageURL, articles.Preview, users.ID, users.Name, users.AvatarURL, users.Bio FROM articles FULL JOIN users ON users.ID = articles.User_Ref WHERE articles.ID > $1 ORDER BY articles.ID ASC LIMIT $2",
		payload.LastID, limit,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmp models.Article
		var tagstring string
		err = rows.Scan(&tmp.ID, &tmp.Title, &tmp.Date, &tmp.Body, &tagstring, &tmp.ImageURL, &tmp.Preview, &tmp.WriterInfo.ID, &tmp.WriterInfo.Name, &tmp.WriterInfo.AvatarURL, &tmp.WriterInfo.Bio)
		if err != nil {
			return nil, err
		}
		tmp.Tag = strings.Split(tagstring, ",")
		articles = append(articles, tmp)
	}

	if len(articles) == 0 {
		return nil, errors.New("NO RESULT")
	}

	return articles, nil
}

//DBArticleGetOne return ID specified article
func DBArticleGetOne(db *sql.DB, payload *models.ArticleFromClient) (models.Article, error) {
	ctx := context.Background()

	var article models.Article
	var tagstring string

	err := db.QueryRowContext(
		ctx,
		"SELECT articles.ID, articles.Title, articles.Date, articles.Body, articles.Tag, articles.ImageURL, articles.Preview, users.ID, users.Name, users.AvatarURL, users.Bio FROM articles JOIN users ON users.ID = articles.User_Ref WHERE articles.ID=$1",
		payload.ID,
	).Scan(&article.ID, &article.Title, &article.Date, &article.Body, &tagstring, &article.ImageURL, &article.Preview, &article.WriterInfo.ID, &article.WriterInfo.Name, &article.WriterInfo.AvatarURL, &article.WriterInfo.Bio)
	if err != nil {
		return article, err
	}

	article.Tag = strings.Split(tagstring, ",")

	return article, nil
}

//DBArticleSearch will get all articles paginated by LastID and filtered by tag and key
func DBArticleSearch(db *sql.DB, payload *models.ArticleFromClient) ([]models.Article, error) {
	ctx := context.Background()

	var limit int64 = 6

	var articles []models.Article

	var (
		rows *sql.Rows
		err  error
	)

	if payload.Filter == "Tag" {
		rows, err = db.QueryContext(
			ctx,
			"SELECT articles.ID, articles.Title, articles.Date, articles.Body, articles.Tag, articles.ImageURL, articles.Preview ,users.ID, users.Name, users.AvatarURL, users.Bio FROM articles JOIN users ON users.ID = articles.User_Ref WHERE articles.Tag LIKE '%' || $1 || '%' AND articles.ID > $2 ORDER BY articles.ID ASC LIMIT $3",
			payload.Key, payload.LastID, limit,
		)
	} else if payload.Filter == "Title" {
		rows, err = db.QueryContext(
			ctx,
			"SELECT articles.ID, articles.Title, articles.Date, articles.Body, articles.Tag, articles.ImageURL, articles.Preview, users.ID, users.Name, users.AvatarURL, users.Bio FROM articles JOIN users ON users.ID = articles.User_Ref WHERE articles.Title LIKE '%' || $1 || '%' AND articles.ID > $2 ORDER BY articles.ID ASC LIMIT $3",
			payload.Key, payload.LastID, limit,
		)
	} else {
		return nil, errors.New("NO FILTER METHOD FOUND")
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmp models.Article
		var tagstring string
		err = rows.Scan(&tmp.ID, &tmp.Title, &tmp.Date, &tmp.Body, &tagstring, &tmp.ImageURL, &tmp.Preview, &tmp.WriterInfo.ID, &tmp.WriterInfo.Name, &tmp.WriterInfo.AvatarURL, &tmp.WriterInfo.Bio)
		if err != nil {
			return nil, err
		}
		tmp.Tag = strings.Split(tagstring, ",")
		articles = append(articles, tmp)
	}

	if len(articles) == 0 {
		return nil, errors.New("NO RESULT")
	}

	return articles, nil
}
