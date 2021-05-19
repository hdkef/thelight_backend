package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"thelight/controller"
	"thelight/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {

	var auth = controller.NewAuthHandler()
	var article = controller.NewArticleHandler()
	var comment = controller.NewCommentHandler()
	var media = controller.NewMediaHandler()

	router := mux.NewRouter()

	router.HandleFunc("/auth/login", utils.Cors(auth.Login()))
	router.HandleFunc("/auth/autologin", utils.Cors(auth.AutoLogin()))

	router.HandleFunc("/article/getall", utils.Cors(article.GetArticles()))
	router.HandleFunc("/article/getone", utils.Cors(article.GetArticle()))
	router.HandleFunc("/article/search", utils.Cors(article.SearchArticles()))
	router.HandleFunc("/article/save", utils.Cors(article.SaveArticle()))
	router.HandleFunc("/article/publish", utils.Cors(article.PublishArticle()))
	router.HandleFunc("/article/delete", utils.Cors(article.DeleteArticle()))

	router.HandleFunc("/comment/getall", utils.Cors(comment.GetComments()))
	router.HandleFunc("/comment/insert", utils.Cors(comment.InsertComment()))

	router.HandleFunc("/media/ws", utils.Cors(media.Media()))

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	fmt.Println("about to serving and listening")

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal(err)
	}

}
