package main

import (
	"encoding/json"
)

type Hub struct {
	clients map[*Client]*Client
	broadcast chan *BroatMsg
	register chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*Client]*Client),
		broadcast: make(chan *BroatMsg),
		register: make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) getUser() []byte {
	msg := UserMsg{Type: "user", Users: make([]User, 0)}
	for client := range h.clients {
		msg.Users = append(msg.Users, *client.user)
	}
	data, err := json.Marshal(msg)
	checkErr(err)
	return data
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case b := <-h.broadcast:
			var message []byte
			chat := b.chat
			if chat.Type == "user" {
				message = h.getUser()
			} else if chat.Type == "addUser" {
				message = resMsgFormat("message", chat.Date, "", chat.Message, false, false)
			}
			for client := range h.clients {
				if chat.Type == "message" {
					message = resMsgFormat(chat.Type, chat.Date, chat.Name, chat.Message, false, client == b.client)
				}
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}