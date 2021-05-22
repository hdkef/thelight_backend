package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thelight/mock"
	"thelight/models"
	"thelight/utils"
)

//ArticleHandler is a type that contain article handlefunc
type ArticleHandler struct {
}

//NewArticleHandler return new pointer of article handler
func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}

//GetArticles give all articles filtered and paginated by ID
func (x *ArticleHandler) GetArticles() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("GetArticles")

		var payload models.ArticleFromClient

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		//TO BE IMPLEMENTED GET ARTICLES AND PAGINATING FROM DB

		var articles []models.Article = mock.Articles

		///////////////////////////////////////////////

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
		fmt.Println("GetArticle")

		var payload models.ArticleFromClient

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		//TO BE IMPLEMENTED GET ARTICLE BY ID

		var article models.Article = mock.Onearticle

		///////////////////////////////////////////////

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
		fmt.Println("SearchArticles")

		var payload models.ArticleFromClient

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		//TO BE IMPLEMENTED GET ARTICLES AND PAGINATING FROM DB FILTERED BY KEY

		var articles []models.Article = mock.Articles

		///////////////////////////////////////////////

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
