package websockets

type Room struct {
	RoomID   string             `json:"room_id"`
	RoomName string             `json:"room_name"`
	Members  map[string]*Member `json:"members"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Member
	Unregister chan *Member
	Broadcast  chan *Message
}

// create constructor
func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Member),
		Unregister: make(chan *Member),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				if _, ok := r.Members[cl.UserID]; ok {
					r.Members[cl.UserID] = cl
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Members[cl.UserID]; ok {
					if len(h.Rooms[cl.RoomID].Members) != 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}

					delete(h.Rooms[cl.RoomID].Members, cl.UserID)
					close(cl.Message)
				}
			}
		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range h.Rooms[m.RoomID].Members {
					cl.Message <- m
				}
			}
		}
	}
}
