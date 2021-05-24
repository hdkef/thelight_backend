package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thelight/driver"
	"thelight/models"
	"thelight/utils"
)

//Settings will handle settings from client
func (x *AuthHandler) Settings() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Settings")

		var payload models.Settings

		err := req.ParseMultipartForm(1024)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		Token := getTokenHeader(req)
		claims, err := checkTokenStringClaims(&Token)
		if err != nil {
			handleTokenErrClearBearer(&res, &err)
			return
		}

		payload.ID = claims.ID
		payload.Name = req.FormValue("Name")
		payload.Bio = req.FormValue("Bio")

		imgurl, err := storeImage(req, "Avatar", "Avatar")

		if err != nil && err != http.ErrMissingFile {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		} else if err != nil && err == http.ErrMissingFile {
			//DO NOTHING
		} else {
			payload.AvatarURL = imgurl
		}

		err = driver.DBAuthSettings(x.db, &payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		err = json.NewEncoder(res).Encode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}
	}
}
