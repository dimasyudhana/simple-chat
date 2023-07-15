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

		h.hub.Register <- cl
		h.hub.Broadcast <- m
		go cl.WriteMessage()
		cl.ReadMessage(h.hub)
	}
}

type RoomResponse struct {
	RoomID   string `json:"room_id"`
	RoomName string `json:"room_name"`
}

func (h *Controller) GetRooms() gin.HandlerFunc {
	return func(c *gin.Context) {
		rooms := make([]RoomResponse, 0)

		for _, r := range h.hub.Rooms {
			rooms = append(rooms, RoomResponse{
				RoomID:   r.RoomID,
				RoomName: r.RoomName,
			})
		}

		c.JSON(http.StatusOK, rooms)
	}
}

type MemberResponse struct {
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

func (h *Controller) GetMembers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var members []MemberResponse
		roomId := c.Param("room_id")

		if _, ok := h.hub.Rooms[roomId]; !ok {
			members = make([]MemberResponse, 0)
			c.JSON(http.StatusOK, members)
		}

		for _, c := range h.hub.Rooms[roomId].Members {
			members = append(members, MemberResponse{
				RoomID:   c.RoomID,
				Username: c.Username,
			})
		}

		c.JSON(http.StatusOK, members)
	}
}
