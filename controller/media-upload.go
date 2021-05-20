package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"thelight/models"
	"thelight/utils"
)

//MediaUpload is endpoint to handle image upload
func (x *MediaHandler) MediaUpload() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("MediaUpload")

		var payload models.MediaPayload

		err := req.ParseMultipartForm(1024)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		payload.ID = req.FormValue("ID")

		imgurl, err := storeImage(req, "Image", "Image")
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		//TOBE IMPLEMENTED SAVE IMGURL TO DB

		response := models.MediaPayload{
			ID:   payload.ID,
			Type: "mediaFromServer",
			Media: models.Media{
				ID:       "id from db",
				ImageURL: imgurl,
			},
		}

		////////////////////////////////////

		go afterStoreImage(&response)

		utils.ResOK(&res, "IMAGE STORED")
	}
}

//afterStoreImage will send websocket message so that client can update the image list
func afterStoreImage(response *models.MediaPayload) {
	fmt.Println("afterStoreImage")

	ws := onlineMap[response.ID]

	ws.WriteJSON(response)
}

//store image will store image and return image path / dir
func storeImage(req *http.Request, formfilename string, foldername string) (string, error) {
	fmt.Println("storeImage")

	uploadedFile, handler, err := req.FormFile(formfilename)
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	workingdir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	filename := handler.Filename
	folderpath := filepath.Join(workingdir, os.Getenv("STATICPATH"), "assets", foldername)
	fileloc := filepath.Join(folderpath, filename)

	err = createNewFolderIfNotExist(folderpath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	targetFile, err := os.OpenFile(fileloc, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		fmt.Println(err)
		return "", err
	}

	return fmt.Sprintf("assets/%s/%s", foldername, filename), nil

}

//createNewFolderIfNotExist will check folder directory and create new folder if not exist
func createNewFolderIfNotExist(path string) error {
	fmt.Println("createNewFolderIfNotExist")

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
