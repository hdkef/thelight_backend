package models

//Article is how article structured
type Article struct {
	ID         int64
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
	ID                int64
	Page              int64 //Stop using offset with limit to paginate,  where ID > and LIMIT instead
	LastID            int64 //but still temporarily use that method for search article, drafts articles, and media until caching logic done
	Filter            string
	Key               string
	ArticleFromClient Article
}

//ArticleFromServer is how article from server structured
type ArticleFromServer struct {
	ArticlesFromServer []Article
	ArticleFromServer  Article
}
