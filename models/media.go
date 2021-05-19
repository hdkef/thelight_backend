package models

import "github.com/gorilla/websocket"

//how media is structured
type Media struct {
	ID       string
	ImageURL string
}

//how client and server communication JSON is structured
type MediaPayload struct {
	ID     string
	Type   string
	Conn   *websocket.Conn
	Token  string
	Medias []Media
	Media  Media
}
