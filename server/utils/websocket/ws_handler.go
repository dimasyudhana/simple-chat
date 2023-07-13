package websockets

import (
	"net/http"

	"github.com/dimasyudhana/simple-chat/app/middlewares"
	"github.com/dimasyudhana/simple-chat/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var log = middlewares.Log()
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type RegisterRoomRequest struct {
	RoomID   string `json:"room_id"`
	RoomName string `json:"room_name"`
}

func (h *Handler) RegisterRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request RegisterRoomRequest
		err := c.ShouldBind(&request)
		if err != nil {
			log.Error("error on bind input")
			response.BadRequestError(c, "Bad request")
			return
		}

		h.hub.Rooms[request.RoomID] = &Room{
			RoomID:   request.RoomID,
			RoomName: request.RoomName,
			Members:  make(map[string]*Member),
		}

		c.JSON(http.StatusOK, response.ResponseFormat(http.StatusOK, "Successfully operation", request, nil))
	}
}

func (h *Handler) JoinRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		roomId := c.Param("room_id")
		userId := c.Query("user_id")
		username := c.Query("username")

		cl := &Member{
			Connection: conn,
			Message:    make(chan *Message, 10),
			UserID:     userId,
			RoomID:     roomId,
			Username:   username,
		}

		m := &Message{
			Content:  "A new user has joined the room",
			RoomID:   roomId,
			Username: username,
		}

		// Register a new member through the register channel
		h.hub.Register <- cl

		// Broadcast that message
		h.hub.Broadcast <- m

		// Write message
		go cl.WriteMessage()
		// Read message
		cl.ReadMessage(h.hub)
	}
}
