package chat

import "github.com/gorilla/websocket"

// Binding from JSON

type WebsocketConnection struct {
	*websocket.Conn
}

type WsPayload struct {
	Action  string              `json:"action"`
	Message string              `json:"message"`
	Conn    WebsocketConnection `json:"-"`
}
