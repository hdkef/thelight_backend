package models

//AuthFromServer is how auth from client structured
type AuthFromClient struct {
	ID    string
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

type Settings struct {
	ID        string
	Name      string
	Bio       string
	AvatarURL string
}
