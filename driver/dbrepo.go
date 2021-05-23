package driver

import (
	"strings"
	"thelight/models"

	"gorm.io/gorm"
)

//DBPublishArticle is a function to store article published. Inside it is a transaction.
func DBPublishArticle(db *gorm.DB, payload *models.ArticleFromClient, ID uint) error {
	if err := db.Create(&Article{
		Title:    payload.ArticleFromClient.Title,
		Body:     payload.ArticleFromClient.Body,
		ImageURL: payload.ArticleFromClient.ImageURL,
		UserID:   ID,
		Tag:      strings.Join(payload.ArticleFromClient.Tag, ","),
	}).Error; err != nil {
		return err
	}
	return nil
}

//DBSaveArticle will save article to draft table and return autoincremented ID
func DBSaveArticle(db *gorm.DB, payload *models.ArticleFromClient, ID uint) (uint, error) {

	draft := Draft{
		UserID:   ID,
		Title:    payload.ArticleFromClient.Title,
		Body:     payload.ArticleFromClient.Body,
		Tag:      strings.Join(payload.ArticleFromClient.Tag, ","),
		ImageURL: payload.ArticleFromClient.ImageURL,
	}

	if err := db.Create(&draft).Error; err != nil {
		return 0, err
	}
	return draft.ID, nil
}

//DBDeleteArticle will delete article
func DBDeleteArticle(db *gorm.DB, ID uint) error {
	if err := db.Unscoped().Delete(&Article{}, ID).Error; err != nil {
		return err
	}
	return nil
}

//DBReadAllArticles will return 6 articles with pagination
func DBReadAllArticles(db *gorm.DB, Page int) ([]models.Article, error) {

	var articles []models.Article

	limit := 6
	offset := (Page - 1) * limit

	rows, err := db.Model(&Article{}).Limit(limit).Offset(offset).Rows()
	defer rows.Close()
	if err != nil {
		return nil, nil
	}

	for rows.Next() {
		var article Article
		db.ScanRows(rows, &article)
		articles = append(articles, models.Article{
			ID:       article.ID,
			Date:     article.CreatedAt.String(),
			Title:    article.Title,
			Body:     article.Body,
			ImageURL: article.ImageURL,
			Tag:      strings.Split(article.Tag, ","),
		})
	}

	return articles, nil
}

//DBReadOneArticle will read specific articles
func DBReadOneArticle(db *gorm.DB, ID uint) (models.Article, error) {

	var article Article

	if err := db.Take(&Article{}, ID).Scan(&article).Error; err != nil {
		return models.Article{}, err
	}
	return models.Article{
		ID:       article.ID,
		Title:    article.Title,
		Date:     article.CreatedAt.String(),
		Body:     article.Body,
		ImageURL: article.ImageURL,
		Tag:      strings.Split(article.Tag, ","),
	}, nil
}

//DBInsertUser will insert user to users table
func DBAuthInsertUser(db *gorm.DB, payload *models.AuthFromClient) error {
	if err := db.Create(&User{
		Name:      payload.Name,
		Pass:      payload.Pass,
		AvatarURL: "https://asset.kompas.com/crops/bzdYfkGm3H7fXaDmBLFTedTaSuU=/65x2:633x381/750x500/data/photo/2021/05/12/609ba9cac54a2.png",
		Bio:       "hi, i am a writer!",
	}).Error; err != nil {
		return err
	}
	return nil
}

//DBReturnPass will return user information
func DBAuthReadUser(db *gorm.DB, Name string) (models.WriterInfo, error) {

	var usr User

	if err := db.Where("Name = ?", Name).Take(&User{}).Scan(&usr).Error; err != nil {
		return models.WriterInfo{}, err
	}
	return models.WriterInfo{
		ID:        usr.ID,
		AvatarURL: usr.AvatarURL,
		Name:      usr.Name,
		Bio:       usr.Bio,
	}, nil
}
