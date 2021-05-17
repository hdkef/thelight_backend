package controller

import "net/http"

//AuthHandler is a type that contain article handlefunc
type AuthHandler struct {
}

//NewAuthHandler return new pointer of auth handler
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

//Login will give jwt and claims to authenticated user
func (x *AuthHandler) Login() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}

//AutoLogin will check the validity of jwt and check jwt's expiratedAt. Will return new jwt and claims if
//jwt's expiredAt within time range, will return error if jwt is not valid, and will return claims if
//jwt valid and jwt's expiredAt not within time range.
func (x *AuthHandler) AutoLogin() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}
