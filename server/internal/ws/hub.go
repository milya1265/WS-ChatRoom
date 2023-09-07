package ws

import "log"

type Room struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Clients map[string]*Client
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}

}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			// V my work
			log.Println("cl <- Register \n", cl)
			if _, ok := h.Rooms[cl.RoomID]; ok == true {

				room := h.Rooms[cl.RoomID]
				if _, ok := room.Clients[cl.ID]; ok == false {
					room.Clients[cl.ID] = cl
				}
			}
		case cl := <-h.Unregister:

			//V my work
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Body:     "user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}

					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
				}
			}

		case message := <-h.Broadcast:

			//V my work
			if _, ok := h.Rooms[message.RoomID]; ok == true {
				for _, cl := range h.Rooms[message.RoomID].Clients {
					cl.Message <- message
				}
			}
		}
	}
}
