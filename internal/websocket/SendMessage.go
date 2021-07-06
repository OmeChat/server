package websocket

import (
	"fmt"
	"github.com/gofiber/websocket/v2"
	"time"
)

type sendMessageResponse struct {
	Message string `json:"message"`
	Sender  string `json:"sender"`
	SentAt  int64  `json:"sent_at"`
	Action  string `json:"action"`
}

// sendMessage sends an message to another user, who is identified by his
// userHash. The message is sent to all clients of the user. The message of the sending
// user is not synced with the other client devices
func sendMessage(c *websocket.Conn, userHash string, payload interface{}) {
	data, ok := PayloadParser("send-message", payload)
	if !ok {
		err := c.WriteJSON(ErrorResponse{
			Message: "Invalid payload",
			Error:   "Cannot parse payload",
			Status:  200,
		})
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
	clients := WS_CONNECTIONS[data["target_hash"].(string)]
	for _, client := range clients {
		err := client.Connection.WriteJSON(sendMessageResponse{
			Message: data["message"].(string),
			Sender:  userHash,
			SentAt:  time.Now().Unix(),
			Action:  "send-message",
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
