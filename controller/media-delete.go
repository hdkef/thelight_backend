package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"thelight/driver"
	"thelight/models"
	"thelight/utils"
)

//MediaDelete handle request that want to delete media (from database and from harddrive)
func (x *MediaHandler) MediaDelete() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		Token := getTokenHeader(req)
		claims, err := checkTokenStringClaims(&Token)
		if err != nil {
			handleTokenErrClearBearer(&res, &err)
			return
		}

		var payload models.MediaPayload
		err = json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		imgdir, err := driver.DBMediaGetImageURL(x.db, &payload, &claims)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		workingdir, err := os.Getwd()
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}
		staticpath := os.Getenv("STATICPATH")
		filedir := filepath.Join(workingdir, staticpath, imgdir)

		err = os.Remove(filedir)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		err = driver.DBMediaDelete(x.db, &payload, &claims)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		utils.ResOK(&res, "IMAGE DELETED")
	}
}
