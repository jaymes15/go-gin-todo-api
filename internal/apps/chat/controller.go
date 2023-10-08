package chat

import (
	"log"
	"net/http"
	"todo/internal/apps/users"
	"todo/pkg/websocketConnection"

	"github.com/gin-gonic/gin"
)

var clients = make(map[WebsocketConnection]string)

type ChatController struct {
}

func NewChatController() *ChatController {
	return &ChatController{}
}

func (chatController *ChatController) Chat(c *gin.Context) {
	user, isValidUser := c.Get("user")
	if !isValidUser {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if user.(users.UserModel).ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	ws, err := websocketConnection.UpgradeConnection.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response WsJsonResponse
	response.Message = "connected to the server"
	err = ws.WriteJSON(response)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn := WebsocketConnection{Conn: ws}

	clients[conn] = user.(users.UserModel).UserName

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	go ListenForWs(&conn)
	go ListenToWsChannel()

}
