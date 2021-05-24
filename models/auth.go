package models

//AuthFromServer is how auth from client structured
type AuthFromClient struct {
	ID    int64
	Name  string
	Pass  string
	Token string
}

//AuthFromServer is how auth from server structured
type AuthFromServer struct {
	WriterInfo WriterInfo
	Token      string
	NewToken   string
}

//Settings is how settings format be sent or received
type Settings struct {
	ID        int64
	Name      string
	Bio       string
	AvatarURL string
}
