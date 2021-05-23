package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thelight/driver"
	"thelight/models"
	"thelight/utils"
)

//SaveArticle will save the article as a draft
func (x *ArticleHandler) SaveArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("SaveArticle")

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

		//TOBE IMPLEMENTED STORE / UPDATE ARTICLE TO DRAFT ARTICLE

		////////////////////////////////////////////////////

		utils.ResOK(&res, "OK")
	}
}

//PublishArticle will publish article to public
func (x *ArticleHandler) PublishArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("PublishArticle")

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

		err = driver.DBPublishArticle(x.db, &payload, claims.ID)

		utils.ResOK(&res, "OK")
	}
}

//DeleteArticle will destroy article from existence
func (x *ArticleHandler) DeleteArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("DeleteArticle")

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

		//TOBE IMPLEMENTED DELETE ARTICLE FROM RELEASED

		///////////////////////////////////////////////

		utils.ResOK(&res, "OK")
	}
}

//EditArticle will edit/update existed article
func (x *ArticleHandler) EditArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("EditArticle")

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

		//TOBE IMPLEMENTED UPDATE ARTICLE IN RELEASED

		/////////////////////////////////////////////

		utils.ResOK(&res, "OK")
	}
}
