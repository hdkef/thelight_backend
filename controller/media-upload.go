package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"thelight/driver"
	"thelight/models"
	"thelight/utils"
)

//MediaUpload is endpoint to handle image upload
func (x *MediaHandler) MediaUpload() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		Token := getTokenHeader(req)
		claims, err := checkTokenStringClaims(&Token)
		if err != nil {
			handleTokenErrClearBearer(&res, &err)
			return
		}

		err = req.ParseMultipartForm(1024)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		imgurl, err := storeImage(req, "Image", "Image")
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		id, err := driver.DBMediaInsert(x.db, imgurl, &claims)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		response := models.MediaPayload{
			ID:   claims.ID,
			Type: "mediaFromServer",
			Media: models.Media{
				ID:       id,
				ImageURL: imgurl,
			},
		}

		go afterStoreImage(&response)

		utils.ResOK(&res, "IMAGE STORED")
	}
}

//afterStoreImage will send websocket message so that client can update the image list
func afterStoreImage(response *models.MediaPayload) {

	ws := onlineMap[response.ID]

	if ws != nil {
		ws.WriteJSON(response)
	}
}

//store image will store image and return image path / dir
func storeImage(req *http.Request, formfilename string, foldername string) (string, error) {

	uploadedFile, handler, err := req.FormFile(formfilename)
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	workingdir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filename := handler.Filename
	folderpath := filepath.Join(workingdir, os.Getenv("STATICPATH"), "assets", foldername)
	fileloc := filepath.Join(folderpath, filename)

	err = createNewFolderIfNotExist(folderpath)
	if err != nil {
		return "", err
	}

	targetFile, err := os.OpenFile(fileloc, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		return "", err
	}

	return fmt.Sprintf("assets/%s/%s", foldername, filename), nil

}

//createNewFolderIfNotExist will check folder directory and create new folder if not exist
func createNewFolderIfNotExist(path string) error {

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return nil
	}
}
