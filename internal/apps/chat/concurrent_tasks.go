package chat

import (
	"fmt"
	"log"
)

var wsChan = make(chan WsPayload)

func ListenForWs(conn *WebsocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {

		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func ListenToWsChannel() {
	var response WsJsonResponse

	for {
		e := <-wsChan
		switch e.Action {
		case "username":
			users := getUserList()
			response.ConnectedUsers = users
			response.Action = "list_users"
			broadCastToAll(response)
		case "left":
			delete(clients, e.Conn)
			users := getUserList()
			response.ConnectedUsers = users
			response.Action = "list_users"
			broadCastToAll(response)

		case "send_message":
			users := getUserList()
			response.ConnectedUsers = users
			response.Message = e.Message
			response.Username = clients[e.Conn]
			response.Action = "send_message"
			broadCastToAll(response)

		}

		// response.Action = "Got here"
		// response.Message = fmt.Sprintf("Some message, and action was %s", e.Action)
		// broadCastToAll(response)

	}
}

func broadCastToAll(response WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("Websocket error")
			_ = client.Close()
			delete(clients, client)
		}
	}
}
