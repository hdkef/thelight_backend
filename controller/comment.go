package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thelight/driver"
	"thelight/models"
	"thelight/utils"

	"gorm.io/gorm"
)

//CommentHandler is a type that contain comment handlefunc
type CommentHandler struct {
	db *gorm.DB
}

//NewCommentHandler return new pointer of comment handler
func NewCommentHandler(db *gorm.DB) *CommentHandler {
	return &CommentHandler{db}
}

//GetComments will get all comments
func (x *CommentHandler) GetComments() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("GetComments")

		var payload models.CommentFromClient

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		comments, err := driver.DBReadComments(x.db, &payload)

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
		fmt.Println("InsertComment")

		var payload models.CommentFromClient

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		err = driver.DBInsertComment(x.db, &payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		utils.ResOK(&res, "OK")

	}
}
