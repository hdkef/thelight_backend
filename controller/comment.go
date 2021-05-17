package controller

import (
	"fmt"
	"net/http"
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
	}
}

//InsertComment will insert one comment
func (x *CommentHandler) InsertComment() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("InsertComment")
	}
}
