package models

//Comment is how comment structured
type Comment struct {
	ID   string
	Name string
	Text string
}

//CommentFromClient is how comment from client structured
type CommentFromClient struct {
	ID                string
	CommentFromClient Comment
}

//CommentFromServer is how comment from server structured
type CommentFromServer struct {
	CommentsFromServer []Comment
}
