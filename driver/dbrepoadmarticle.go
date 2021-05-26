package driver

import (
	"context"
	"database/sql"
	"strings"
	"thelight/models"
	"time"
)

//DBArticlePublish
func DBArticlePublish(db *sql.DB, payload *models.ArticleFromClient, claims *models.WriterInfo) (int64, error) {
	ctx := context.Background()

	date := time.Now()
	tag := strings.Join(payload.ArticleFromClient.Tag, ",")

	var insertedID int64

	err := db.QueryRowContext(
		ctx,
		"INSERT INTO articles (Title, Date, Body, Tag, ImageURL, User_Ref) VALUES ($1,$2,$3,$4,$5,$6) RETURNING ID",
		payload.ArticleFromClient.Title, date, payload.ArticleFromClient.Body, tag, payload.ArticleFromClient.ImageURL, claims.ID,
	).Scan(&insertedID)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

//DBArticleSaveAs
func DBArticleSaveAs(db *sql.DB, payload *models.ArticleFromClient, claims *models.WriterInfo) (int64, error) {
	ctx := context.Background()

	date := time.Now()
	tag := strings.Join(payload.ArticleFromClient.Tag, ",")

	var insertedID int64

	err := db.QueryRowContext(
		ctx,
		"INSERT INTO drafts (Title, Date, Body, Tag, ImageURL, User_Ref) VALUES ($1,$2,$3,$4,$5,$6) RETURNING ID",
		payload.ArticleFromClient.Title, date, payload.ArticleFromClient.Body, tag, payload.ArticleFromClient.ImageURL, claims.ID,
	).Scan(&insertedID)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

//DBArticleEdit
func DBArticleEdit(db *sql.DB, payload *models.ArticleFromClient) (int64, error) {
	ctx := context.Background()

	var insertedID int64
	tag := strings.Join(payload.ArticleFromClient.Tag, ",")

	err := db.QueryRowContext(
		ctx,
		"UPDATE articles SET Title=$1,Body=$2,Tag=$3,ImageURL=$4 WHERE ID=$5 RETURNING ID",
		payload.ArticleFromClient.Title, payload.ArticleFromClient.Body, tag, payload.ArticleFromClient.ImageURL, payload.ArticleFromClient.ID,
	).Scan(&insertedID)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

//DBArticleDelete
func DBArticleDelete(db *sql.DB, payload *models.ArticleFromClient) error {
	ctx := context.Background()

	_, err := db.ExecContext(ctx, "DELETE FROM articles WHERE id=$1", payload.ID)
	if err != nil {
		return err
	}

	return nil
}

//DBArticleSave
func DBArticleSave(db *sql.DB, payload *models.ArticleFromClient, claims *models.WriterInfo) (int64, error) {
	ctx := context.Background()

	var insertedID int64
	tag := strings.Join(payload.ArticleFromClient.Tag, ",")

	err := db.QueryRowContext(
		ctx,
		"UPDATE drafts SET Title=$1,Body=$2,Tag=$3,ImageURL=$4 WHERE ID=$5 AND USER_REF=$6 RETURNING ID",
		payload.ArticleFromClient.Title, payload.ArticleFromClient.Body, tag, payload.ArticleFromClient.ImageURL, payload.ArticleFromClient.ID, claims.ID,
	).Scan(&insertedID)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

//DBArticleDraftGetAll
func DBArticleDraftGetAll(db *sql.DB, payload *models.ArticleFromClient, claims *models.WriterInfo) ([]models.Article, error) {
	ctx := context.Background()

	var limit int64 = 6
	offset := (payload.Page - 1) * limit

	var articles []models.Article

	rows, err := db.QueryContext(
		ctx,
		"SELECT ID, Title, Date, Body, Tag, ImageURL FROM articles WHERE USER_REF=$1 LIMIT $2 OFFSET $3",
		claims.ID, limit, offset,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmp models.Article
		var tagstring string
		err = rows.Scan(&tmp.ID, &tmp.Title, &tmp.Date, &tmp.Body, &tagstring, &tmp.ImageURL)
		if err != nil {
			return nil, err
		}
		tmp.Tag = strings.Split(tagstring, ",")
		articles = append(articles, tmp)
	}

	return articles, nil
}

//DBArticleDraftGetOne
func DBArticleDraftGetOne(db *sql.DB, payload *models.ArticleFromClient, claims *models.WriterInfo) (models.Article, error) {
	ctx := context.Background()

	var article models.Article
	var tagstring string

	err := db.QueryRowContext(
		ctx,
		"SELECT ID, Title, Date, Body, Tag, ImageURL FROM drafts WHERE ID=$1 AND USER_REF=$2",
		payload.ID, claims.ID,
	).Scan(&article.ID, &article.Title, &article.Date, &article.Body, &tagstring, &article.ImageURL)
	if err != nil {
		return article, err
	}

	article.Tag = strings.Split(tagstring, ",")

	return article, nil
}
