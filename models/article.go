package models

//Article is how article structured
type Article struct {
	ID         string
	Date       string
	Title      string
	Preview    string
	Body       string
	Tag        []string
	ImageURL   string
	WriterInfo WriterInfo
}

//ArticleFromClient is how article from client structured
type ArticleFromClient struct {
	ID                string
	Page              string
	Filter            string
	Key               string
	ArticleFromClient Article
}

//ArticleFromServer is how article from server structured
type ArticleFromServer struct {
	ArticlesFromServer []Article
	ArticleFromServer  Article
}
