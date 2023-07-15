package websockets

import (
	"net/http"

	"github.com/dimasyudhana/simple-chat/app/middlewares"
	"github.com/gorilla/websocket"
)

var log = middlewares.Log()

type Member struct {
	Connection *websocket.Conn
	Message    chan *Message
	UserID     string `json:"user_id"`
	RoomID     string `json:"room_id"`
	Username   string `json:"username"`
	Quit       chan bool
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c *Member) WriteMessage() {
	defer func() {
		c.Connection.Close()
	}()
	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Connection.WriteJSON(message)
	}
}

func (c *Member) ReadMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Connection.Close()
	}()

	for {
		_, m, err := c.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Sugar().Infof("error: %v", err)
			}
			break
		}

		msg := &Message{
			Content:  string(m),
			RoomID:   c.RoomID,
			Username: c.Username,
		}

		hub.Broadcast <- msg
	}
}

func (c *Member) Close(hub *Hub) {
	hub.Unregister <- c
	c.Connection.Close()
	close(c.Message)
	c.Quit <- true
}
