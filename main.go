package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"thelight/controller"
	"thelight/driver"
	"thelight/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {

	db, _ := driver.InitiateDB()
	defer db.Close()

	var auth = controller.NewAuthHandler(db)
	var article = controller.NewArticleHandler(db)
	var comment = controller.NewCommentHandler(db)
	var media = controller.NewMediaHandler(db)

	router := mux.NewRouter()

	router.HandleFunc("/auth/login", utils.Cors(auth.Login()))
	router.HandleFunc("/auth/autologin", utils.Cors(auth.AutoLogin()))
	router.HandleFunc("/auth/settings", utils.Cors(auth.Settings()))
	router.HandleFunc("/auth/reg", utils.Cors(auth.Register()))

	router.HandleFunc("/article/getall", utils.Cors(article.GetArticles()))
	router.HandleFunc("/article/getone", utils.Cors(article.GetArticle()))
	router.HandleFunc("/article/search", utils.Cors(article.SearchArticles()))
	router.HandleFunc("/article/saveas", utils.Cors(article.SaveArticleAs()))
	router.HandleFunc("/article/save", utils.Cors(article.SaveArticle()))
	router.HandleFunc("/article/publish", utils.Cors(article.PublishArticle()))
	router.HandleFunc("/article/delete", utils.Cors(article.DeleteArticle()))
	router.HandleFunc("/article/edit", utils.Cors(article.EditArticle()))

	router.HandleFunc("/draft/getall", utils.Cors(article.GetDraftAticles()))
	router.HandleFunc("/draft/getone", utils.Cors(article.GetDraftArticle()))

	router.HandleFunc("/comment/getall", utils.Cors(comment.GetComments()))
	router.HandleFunc("/comment/insert", utils.Cors(comment.InsertComment()))
	router.HandleFunc("/comment/delete", utils.Cors(comment.DeleteComment()))

	router.HandleFunc("/media/ws", utils.Cors(media.Media()))
	router.HandleFunc("/media/upload", utils.Cors(media.MediaUpload()))

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	fmt.Println("about to serving and listening")

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal(err)
	}

}
