package websockets

import (
	"github.com/gorilla/websocket"
)

type Member struct {
	Connection *websocket.Conn
	Message    chan *Message
	UserID     string `json:"user_id"`
	RoomID     string `json:"room_id"`
	Username   string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
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
