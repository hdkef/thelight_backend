package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"thelight/models"
	"thelight/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	EXPINTEGER         = 100000
	TIMEDIFF   float64 = 5000
)

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
		fmt.Println("Login")

		var payload models.AuthFromClient

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		//TOBEIMPLEMENTED GET DATA FROM DB
		user := models.WriterInfo{
			ID:        "1",
			AvatarURL: "avatar",
			Name:      "name",
			Bio:       "bio",
		}
		/////////////////////////////////////

		token, err := createToken(&user)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		//TOBE REPLACED WITH REAL WRITERINFO FROM DB
		response := models.AuthFromServer{
			WriterInfo: models.WriterInfo{
				ID:        user.ID,
				Name:      user.Name,
				AvatarURL: user.AvatarURL,
				Bio:       user.Bio,
			},
			Token: token,
		}

		err = json.NewEncoder(res).Encode(&response)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}
	}
}

//AutoLogin will check the validity of jwt and check jwt's expiratedAt. Will return new jwt and claims if
//jwt's expiredAt within time range, will return error if jwt is not valid, and will return claims if
//jwt valid and jwt's expiredAt not within time range.
func (x *AuthHandler) AutoLogin() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("AutoLogin")

		Token := getTokenHeader(*&req)

		claims, err := checkTokenStringClaims(&Token)
		if err != nil && err.Error() == "NEED NEW TOKEN" {
			sendNewToken(&res, &claims)
			return
		} else if err != nil && err.Error() != "NEED NEW TOKEN" {
			handleTokenErrClearBearer(&res, &err)
			return
		}

		response := models.AuthFromServer{
			WriterInfo: claims,
		}

		err = json.NewEncoder(res).Encode(&response)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}
	}
}

//sendNewToken will refresh old token with the new one
func sendNewToken(res *http.ResponseWriter, claims *models.WriterInfo) {
	token, err := createToken(claims)
	if err != nil {
		utils.ResErr(res, http.StatusInternalServerError, err)
		return
	}

	response := models.AuthFromServer{
		WriterInfo: *claims,
		NewToken:   token,
	}

	err = json.NewEncoder(*res).Encode(&response)
	if err != nil {
		utils.ResErr(res, http.StatusInternalServerError, err)
		return
	}
}

//newClaimsMap create new jwt mapclaims from user struct and return it
func newClaimsMap(user *models.WriterInfo) jwt.MapClaims {
	fmt.Println("newClaimsMap")
	var claims jwt.MapClaims = make(jwt.MapClaims)

	var userval = reflect.ValueOf(*user)
	var usertype = reflect.TypeOf(*user)

	for i := 0; i < userval.NumField(); i++ {
		fieldName := usertype.Field(i).Name
		fieldValue := userval.Field(i).String()
		claims[fieldName] = fieldValue
	}

	claims["exp"] = time.Now().Unix() + EXPINTEGER
	//this code is intended to be place after for loop so that new exp override old exp for refresh token

	return claims
}

//createToken return new token string
func createToken(user *models.WriterInfo) (string, error) {
	fmt.Println("createToken")

	secret := os.Getenv("SECRET")

	claims := newClaimsMap(user)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := newToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}

//getTokenHeader return token from Bearer header
func getTokenHeader(req *http.Request) string {
	return req.Header.Get("Bearer")
}

//handleTokenErrClearBearer set header ClearBearer when authentication fail and send response error
func handleTokenErrClearBearer(res *http.ResponseWriter, err *error) {
	fmt.Println("handleTokenErrClearBearer")
	(*res).Header().Set("Clearbearer", "OK")
	(*res).Header().Set("Access-Control-Expose-Headers", "Clearbearer")
	utils.ResErr(res, http.StatusUnauthorized, *err)
}

//checkTokenStringClaims will validate token and return claims (NOT map claims) and error
//TOBE IMPLEMENTED IF EXP BETWEEN RANGE RETURN ERR "NEED NEW TOKEN"
func checkTokenStringClaims(token *string) (models.WriterInfo, error) {
	fmt.Println("checkTokenStringClaims")

	if *token == "" {
		return models.WriterInfo{}, errors.New("NO TOKEN")
	}

	parsedToken, err := parseTokenString(token)

	if err != nil {
		fmt.Println(err)
		return models.WriterInfo{}, err
	}
	if parsedToken.Valid == false {
		return models.WriterInfo{}, errors.New("INVALID TOKEN")
	}

	var mapClaims jwt.MapClaims = parsedToken.Claims.(jwt.MapClaims)

	claims := models.WriterInfo{
		ID:        mapClaims["ID"].(string),
		Name:      mapClaims["Name"].(string),
		AvatarURL: mapClaims["AvatarURL"].(string),
		Bio:       mapClaims["Bio"].(string),
	}
	return claims, nil
}

//checkTokenStringErr will validate token and return bool // being used in checkOrigin websocket
func checkTokenStringBool(token *string) bool {
	fmt.Println("checkTokenStringBool")
	err := checkTokenStringErr(token)
	if err != nil {
		return false
	}
	return true
}

//checkTokenStringErr will validate token and return error // being used in endpoints that doesn't need claims
func checkTokenStringErr(token *string) error {
	fmt.Println("checkTokenStringErr")

	if *token == "" {
		return errors.New("NO TOKEN")
	}

	parsedToken, err := parseTokenString(token)

	if err != nil && err.Error() != "NEED NEW TOKEN" {
		fmt.Println(err)
		return err
	}
	if parsedToken.Valid == false {
		return errors.New("INVALID TOKEN")
	}
	return nil
}

//parseTokenString will parse token string and return jwt.Token & error
func parseTokenString(token *string) (*jwt.Token, error) {
	fmt.Println("parseTokenString")
	parsedToken, err := jwt.Parse(*token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	renew := checkTokenRenew(parsedToken)
	if renew == true {
		return nil, errors.New("NEED NEW TOKEN")
	}

	return parsedToken, nil
}

//checkTokenRenew will return true if token expiration time between range that need to be renewed
func checkTokenRenew(token *jwt.Token) bool {
	fmt.Println("checkTokenRenew")

	now := time.Now().Unix()
	timeSubNow := (*token).Claims.(jwt.MapClaims)["exp"].(float64) - float64(now)

	if timeSubNow <= TIMEDIFF {
		return true
	}

	return false
}
