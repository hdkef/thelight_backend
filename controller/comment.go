package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"thelight/driver"
	"thelight/models"
	"thelight/utils"
)

//CommentHandler is a type that contain comment handlefunc
type CommentHandler struct {
	db *sql.DB
}

//NewCommentHandler return new pointer of comment handler
func NewCommentHandler(db *sql.DB) *CommentHandler {
	return &CommentHandler{db}
}

//GetComments will get all comments
func (x *CommentHandler) GetComments() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		var payload models.CommentFromClient

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		comments, err := driver.DBCommentGetAll(x.db, &payload)

		response := models.CommentFromServer{
			CommentsFromServer: comments,
		}

		err = json.NewEncoder(res).Encode(&response)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}
	}
}

//InsertComment will insert one comment
func (x *CommentHandler) InsertComment() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		var payload models.CommentFromClient

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		_, err = driver.DBCommentInsert(x.db, &payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		utils.ResOK(&res, "OK")

	}
}

//DeleteComment will insert one comment
func (x *CommentHandler) DeleteComment() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		Token := getTokenHeader(req)
		err := checkTokenStringErr(&Token)
		if err != nil {
			handleTokenErrClearBearer(&res, &err)
			return
		}

		var payload models.CommentFromClient

		err = json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		err = driver.DBCommentDelete(x.db, &payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		utils.ResOK(&res, "OK")

	}
}
