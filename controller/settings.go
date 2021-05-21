package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
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

		payload.ID = req.FormValue("ID")
		Name := req.FormValue("Name")
		Bio := req.FormValue("Bio")

		if Name != "" {
			payload.Name = Name
		}

		if Bio != "" {
			payload.Bio = Bio
		}

		imgurl, err := storeImage(req, "Avatar", "Avatar")

		if err != nil && err != http.ErrMissingFile {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		} else if err != nil && err == http.ErrMissingFile {
			//DO NOTHING
		} else {
			payload.AvatarURL = imgurl
			//TOBE IMPLEMENTED STORE IMGURL TO DB

			////////////////////////////////////
		}

		err = json.NewEncoder(res).Encode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}
	}
}
