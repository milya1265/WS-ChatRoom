package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{hub: h}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusCreated, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomID := c.Param("roomID")
	clientID := c.Query("clientID")
	username := c.Query("username")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	log.Println("Trying to join the room")
	mess := &Message{
		Body:     "A new user has joined the room",
		RoomID:   cl.RoomID,
		Username: cl.Username,
	}

	log.Println(mess)

	// Register a new client throw the register chanel
	h.hub.Register <- cl
	//Broadcast that message
	h.hub.Broadcast <- mess

	go cl.WriteMessage()
	cl.ReadMessage(h.hub)

}

type RoomRes struct {
	ID   string
	Name string
}

func (h *Handler) GetRooms(c *gin.Context) {

	rooms := make([]*RoomRes, 0)

	for _, room := range h.hub.Rooms {
		roomRes := &RoomRes{
			ID:   room.ID,
			Name: room.Name,
		}
		rooms = append(rooms, roomRes)
	}

	c.JSON(http.StatusOK, rooms)
}

type ClientRes struct {
	ID       string
	Username string
}

func (h *Handler) GetClients(c *gin.Context) {
	roomID := c.Param("roomID")

	clients := make([]*ClientRes, 0)

	for _, cl := range h.hub.Rooms[roomID].Clients {
		clRes := &ClientRes{
			ID:       cl.ID,
			Username: cl.Username,
		}
		clients = append(clients, clRes)
	}

	c.JSON(http.StatusOK, clients)
}
