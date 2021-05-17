package controller

import "net/http"

//MediaHandler is a type that contain media handlefunc
type MediaHandler struct {
}

//NewMediaHandler return new pointer of comment handler
func NewMediaHandler() *MediaHandler {
	return &MediaHandler{}
}

//UploadImage will store image to folder and insert IMGDir to database
func (x *MediaHandler) UploadImage() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}

//GetImageDirs will give all image directories
func (x *MediaHandler) GetImageDirs() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}
