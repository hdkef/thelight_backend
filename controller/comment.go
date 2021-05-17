package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thelight/mock"
	"thelight/models"
	"thelight/utils"
)

//CommentHandler is a type that contain comment handlefunc
type CommentHandler struct {
}

//NewCommentHandler return new pointer of comment handler
func NewCommentHandler() *CommentHandler {
	return &CommentHandler{}
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

		//TOBE IMPLEMENTED GET ALL COMMENTS FROM DB

		var comments []models.Comment = mock.Comments

		///////////////////////////////////////////

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
	}
}
