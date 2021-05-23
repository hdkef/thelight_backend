package models

//Comment is how comment structured
type Comment struct {
	ID   uint
	Name string
	Text string
}

//CommentFromClient is how comment from client structured
type CommentFromClient struct {
	ID                uint
	CommentFromClient Comment
}

//CommentFromServer is how comment from server structured
type CommentFromServer struct {
	CommentsFromServer []Comment
}
