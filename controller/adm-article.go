package controller

import (
	"encoding/json"
	"net/http"
	"thelight/driver"
	"thelight/models"
	"thelight/utils"
)

//SaveArticleAs will save the article as a draft
func (x *ArticleHandler) SaveArticleAs() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		Token := getTokenHeader(req)
		claims, err := checkTokenStringClaims(&Token)
		if err != nil {
			handleTokenErrClearBearer(&res, &err)
			return
		}

		var payload models.ArticleFromClient

		err = json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		insertedID, err := driver.DBArticleSaveAs(x.db, &payload, &claims)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		json.NewEncoder(res).Encode(struct {
			ID int64
		}{
			insertedID,
		})
	}
}

//PublishArticle will publish article to public
func (x *ArticleHandler) PublishArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		Token := getTokenHeader(req)
		claims, err := checkTokenStringClaims(&Token)
		if err != nil {
			handleTokenErrClearBearer(&res, &err)
			return
		}

		var payload models.ArticleFromClient

		err = json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		_, err = driver.DBArticlePublish(x.db, &payload, &claims)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		utils.ResOK(&res, "OK")
	}
}

//DeleteArticle will destroy article from existence
func (x *ArticleHandler) DeleteArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		Token := getTokenHeader(req)
		err := checkTokenStringErr(&Token)
		if err != nil {
			handleTokenErrClearBearer(&res, &err)
			return
		}

		var payload models.ArticleFromClient

		err = json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		err = driver.DBArticleDelete(x.db, &payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		utils.ResOK(&res, "OK")
	}
}

//EditArticle will edit/update existed article
func (x *ArticleHandler) EditArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		Token := getTokenHeader(req)
		err := checkTokenStringErr(&Token)
		if err != nil {
			handleTokenErrClearBearer(&res, &err)
			return
		}

		var payload models.ArticleFromClient

		err = json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		_, err = driver.DBArticleEdit(x.db, &payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		utils.ResOK(&res, "OK")
	}
}

//SaveArticle will save the article in a draft
func (x *ArticleHandler) SaveArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		Token := getTokenHeader(req)
		claims, err := checkTokenStringClaims(&Token)
		if err != nil {
			handleTokenErrClearBearer(&res, &err)
			return
		}

		var payload models.ArticleFromClient

		err = json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		_, err = driver.DBArticleSave(x.db, &payload, &claims)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		utils.ResOK(&res, "OK")
	}
}

//GetDraftArticles will return all articles paginated by LastID
func (x *ArticleHandler) GetDraftArticles() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		Token := getTokenHeader(req)
		claims, err := checkTokenStringClaims(&Token)
		if err != nil {
			handleTokenErrClearBearer(&res, &err)
			return
		}

		var payload models.ArticleFromClient

		err = json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		article, err := driver.DBArticleDraftGetAll(x.db, &payload, &claims)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		response := models.ArticleFromServer{
			ArticlesFromServer: article,
		}

		err = json.NewEncoder(res).Encode(&response)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}
	}
}

//GetDraftArticle will get ID specified draft
func (x *ArticleHandler) GetDraftArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		Token := getTokenHeader(req)
		claims, err := checkTokenStringClaims(&Token)
		if err != nil {
			handleTokenErrClearBearer(&res, &err)
			return
		}

		var payload models.ArticleFromClient

		err = json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		article, err := driver.DBArticleDraftGetOne(x.db, &payload, &claims)
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

//DeleteDraft will destroy draft from existence
func (x *ArticleHandler) DeleteDraft() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		Token := getTokenHeader(req)
		err := checkTokenStringErr(&Token)
		if err != nil {
			handleTokenErrClearBearer(&res, &err)
			return
		}

		var payload models.ArticleFromClient

		err = json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		err = driver.DBArticleDraftDelete(x.db, &payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		utils.ResOK(&res, "OK")
	}
}
