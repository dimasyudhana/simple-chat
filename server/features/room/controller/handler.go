package controller

import (
	"net/http"

	"github.com/dimasyudhana/simple-chat/app/middlewares"
	"github.com/dimasyudhana/simple-chat/features/room"
	"github.com/dimasyudhana/simple-chat/utils/response"
	"github.com/dimasyudhana/simple-chat/utils/websockets"
	"github.com/gin-gonic/gin"
)

var log = middlewares.Log()

type Controller struct {
	service room.UseCase
	hub     *websockets.Hub
}

func New(us room.UseCase, h *websockets.Hub) room.Controller {
	return &Controller{
		service: us,
		hub:     h,
	}
}

func (h *Controller) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := RegisterRequest{}
		err := c.ShouldBind(&request)
		if err != nil {
			log.Error("error on bind input")
			response.BadRequestError(c, "Bad request")
			return
		}

		h.hub.Rooms[request.RoomID] = &websockets.Room{
			RoomID:   request.RoomID,
			RoomName: request.RoomName,
			Members:  make(map[string]*websockets.Member),
		}

		c.JSON(http.StatusOK, response.ResponseFormat(http.StatusOK, "Successfully operation", request, nil))
	}
}

func (h *Controller) Join() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := websockets.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		roomId := c.Param("room_id")
		userId := c.Query("user_id")
		username := c.Query("username")

		cl := &websockets.Member{
			Connection: conn,
			Message:    make(chan *websockets.Message, 10),
			UserID:     userId,
			RoomID:     roomId,
			Username:   username,
		}

		m := &websockets.Message{
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
