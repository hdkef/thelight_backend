package controller

import "net/http"

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

	}
}

//GetArticle give one article. Probably because user view Article without going to landing-page first
func (x *ArticleHandler) GetArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}

//SearchArticles give all articles filtered by something and paginated by ID
func (x *ArticleHandler) SearchArticles() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}

//SaveArticle will save the article as a draft
func (x *ArticleHandler) SaveArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}

//PublishArticle will publish article to public
func (x *ArticleHandler) PublishArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}

//DeleteArticle will destroy article from existence
func (x *ArticleHandler) DeleteArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}

//EditArticle will edit/update existed article
func (x *ArticleHandler) EditArticle() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}
