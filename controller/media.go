package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"thelight/mock"
	"thelight/models"
	"time"

	"github.com/gorilla/websocket"
)

//MediaHandler is a type that contain media handlefunc
type MediaHandler struct {
}

//NewMediaHandler return new pointer of comment handler
func NewMediaHandler() *MediaHandler {
	return &MediaHandler{}
}

//upgrader to upgrade http to websocket.Conn
var upgrader websocket.Upgrader = websocket.Upgrader{
	CheckOrigin: func(req *http.Request) bool {
		// Bearer := req.Header.Get("Bearer")
		// return checkTokenStringBool(&Bearer)
		return true
	},
}

//const for pingponger

const (
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

//onlineMap to store websocket.Conn with ID key

var onlineMap map[string]*websocket.Conn = make(map[string]*websocket.Conn)

//various channel to handle various payload type

var initFromClientChan chan models.MediaPayload = make(chan models.MediaPayload)
var pagingFromClientChan chan models.MediaPayload = make(chan models.MediaPayload)

//Media will handle websocket connection to media
func (x *MediaHandler) Media() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("UploadImage")

		ws, err := upgrader.Upgrade(res, req, res.Header())
		if err != nil {
			log.Println(err)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())

		go readAndSend(cancel, ws)
		go receiveAndHandle(ctx)
	}
}

//readAndSend read incoming payload, assign websocket.Conn to payload and send to corresponding channel
func readAndSend(cancel context.CancelFunc, ws *websocket.Conn) {
	fmt.Println("readAndSend")

	var payload models.MediaPayload = models.MediaPayload{
		Conn: ws,
	}
	defer cancel()

	for {
		err := ws.ReadJSON(&payload)
		if err != nil {
			log.Println(err)
			break
		}
		switch payload.Type {
		case "initFromClient":
			initFromClientChan <- payload
		case "pagingFromClient":
			pagingFromClientChan <- payload
		}
	}
}

//receiveAndHandle receive payload from channel and handle to corresponding function
func receiveAndHandle(ctx context.Context) {
	fmt.Println("receiveAndHandle")
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-initFromClientChan:
			initFromClient(msg)
		case msg := <-pagingFromClientChan:
			pagingFromClient(msg)
		}
	}
}

//initFromClient handle initFromClient payload type
func initFromClient(payload models.MediaPayload) {
	fmt.Println("initFromClient")

	onlineMap[payload.ID] = payload.Conn

	fmt.Println("onlineMap : ", onlineMap)

	go pingPonger(payload.ID, payload.Conn)

	//TOBE IMPLEMENTED GET ALL IMAGE DIRS FROM DB AND VERIFY TOKEN
	medias := mock.Medias
	/////////////////////////////////////////////
	response := models.MediaPayload{
		ID:     payload.ID,
		Type:   "initFromServer",
		Medias: medias,
	}

	payload.Conn.WriteJSON(&response)
}

//pingPonger will ping websocket conn and delete onlineMap if return error for defined time range
func pingPonger(ID string, ws *websocket.Conn) {
	fmt.Println("pingPonger")

	ws.SetPongHandler(func(appData string) error {
		ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	timer := time.NewTicker(pingPeriod)

	defer func() {
		timer.Stop()
		if onlineMap[ID] == ws {
			delete(onlineMap, ID)
		}
		fmt.Println("onlineMap : ", onlineMap)
	}()

	for {
		select {
		case <-timer.C:
			fmt.Println("tick")
			if err := ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

//pagingFromClient handle pagingFromClient payload type
func pagingFromClient(payload models.MediaPayload) {
	fmt.Println("pagingFromClient")

	//TOBE IMPLEMENTED GET ALL IMAGE DIRS FROM DB AND VERIFY TOKEN
	medias := mock.Medias
	/////////////////////////////////////////////
	response := models.MediaPayload{
		ID:     payload.ID,
		Type:   "pagingFromServer",
		Medias: medias,
	}

	payload.Conn.WriteJSON(&response)
}
