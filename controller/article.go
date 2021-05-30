package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"thelight/driver"
	"thelight/models"
	"thelight/utils"
)

//ArticleHandler is a type that contain article handlefunc
type ArticleHandler struct {
	db *sql.DB
}

//NewArticleHandler return new pointer of article handler
func NewArticleHandler(db *sql.DB) *ArticleHandler {
	return &ArticleHandler{db}
}

//GetArticles give all articles filtered and paginated by ID
func (x *ArticleHandler) GetArticles() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		var payload models.ArticleFromClient

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		articles, err := driver.DBArticleGetAll(x.db, &payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		response := models.ArticleFromServer{
			ArticlesFromServer: articles,
		}

		err = json.NewEncoder(res).Encode(&response)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}
	}
}

//GetArticle give one article. Probably because user view Article without going to landing-page first
func (x *ArticleHandler) GetArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		var payload models.ArticleFromClient

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		article, err := driver.DBArticleGetOne(x.db, &payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		response := models.ArticleFromServer{
			ArticleFromServer: article,
		}

		err = json.NewEncoder(res).Encode(&response)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}
	}
}

//SearchArticles give all articles filtered by something and paginated by ID
func (x *ArticleHandler) SearchArticles() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		var payload models.ArticleFromClient

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		articles, err := driver.DBArticleSearch(x.db, &payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		response := models.ArticleFromServer{
			ArticlesFromServer: articles,
		}

		err = json.NewEncoder(res).Encode(&response)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}
	}
}
