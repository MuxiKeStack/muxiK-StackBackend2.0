package socket

import "github.com/pkg/errors"

// Message contains content and object
type Message struct {
	AddresseeSid string
	Raw          []byte
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	sender chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	broadcaster chan Message
}

func NewHub() *Hub {
	return &Hub{
		sender:     make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.sid] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.sid]; ok {
				delete(h.clients, client.sid)
				close(client.send)
			}
		case message := <-h.sender:
			if client, ok := h.clients[message.AddresseeSid]; ok {
				client.send <- message
			}
		case message := <-h.broadcaster:
			for _, client := range h.clients {
				client.send <- message
			}
		}
	}
}

func (h *Hub) Send(addresseeSid string, raw []byte) error {
	message := Message{AddresseeSid: addresseeSid, Raw: raw}
	select {
	case h.sender <- message:
		return nil
	default:
		return errors.New("hub is fault")
	}
}

func (h *Hub) Broadcast(raw []byte) error {
	message := Message{Raw: raw}
	h.broadcaster <- message
	select {
	case h.broadcaster <- message:
		return nil
	default:
		return errors.New("hub is fault")
	}
}
