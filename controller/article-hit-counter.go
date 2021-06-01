package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"thelight/driver"
	"thelight/models"
	"thelight/utils"
)

//ArticleHitCounterHandler is to handle request
func (x *ArticleHandler) ArticleHitCounterHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		go articleHitCounter(x.db, req)
		utils.ResOK(&res, "Hit article hit counter")
	}
}

//articleHitCounter is a function that will increment value each article get hit
func articleHitCounter(db *sql.DB, req *http.Request) {
	var payload models.AnalyticPayload
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		return
	}

	//COUNTER
	_, err = driver.DBArticleHitCounter(db, &payload)
	if err != nil {
		return
	}
	//COUNTER
}
