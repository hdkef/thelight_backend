package models

import "github.com/gorilla/websocket"

//how media is structured
type Media struct {
	ID       int64
	ImageURL string
	UserRef  int64
}

//how client and server communication JSON is structured
type MediaPayload struct {
	ID     int64
	Type   string
	Conn   *websocket.Conn
	Token  string
	Medias []Media
	Media  Media
	Page   int64
}
