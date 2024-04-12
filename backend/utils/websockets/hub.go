package websockets

import "sync"

type Room struct {
	RoomID   string             `json:"room_id"`
	RoomName string             `json:"room_name"`
	Members  map[string]*Member `json:"members"`
	mutex    sync.RWMutex
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Member
	Unregister chan *Member
	Broadcast  chan *Message
	mutex      sync.RWMutex
}

// create constructor
func New() *Hub {
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
			h.mutex.Lock()
			if room, ok := h.Rooms[cl.RoomID]; ok {
				room.mutex.Lock()
				room.Members[cl.UserID] = cl
				room.mutex.Unlock()
			}
			h.mutex.Unlock()
		case cl := <-h.Unregister:
			h.mutex.Lock()
			if room, ok := h.Rooms[cl.RoomID]; ok {
				room.mutex.Lock()
				if _, ok := room.Members[cl.UserID]; ok {
					delete(room.Members, cl.UserID)
					close(cl.Message)

					if len(room.Members) == 0 {
						delete(h.Rooms, cl.RoomID)
						room.mutex.Unlock()
						leaveMsg := &Message{
							Content:  "User left the room",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
						h.Broadcast <- leaveMsg
						break
					}
				}
				room.mutex.Unlock()
			}
			h.mutex.Unlock()
		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {
				go func(m *Message) {
					for _, cl := range h.Rooms[m.RoomID].Members {
						cl.Message <- m
					}
				}(m)
			}
		}
	}
}
