package websocket

import (
	"github.com/gofiber/websocket/v2"
	"time"
)

type sendMessageResponse struct {
	Message string `json:"message"`
	Sender  string `json:"sender"`
	SentAt  string `json:"sent_at"`
}

// sendMessage sends an message to another user, who is identified by his
// userHash. The message is sent to all clients of the user. The message of the sending
// user is not synced with the other client devices
func sendMessage(c *websocket.Conn, userHash string, payload interface{}) {
	data, ok := PayloadParser("send-message", payload)
	if !ok {
		_ = c.WriteJSON(ErrorResponse{
			Message: "Invalid payload",
			Error:   "Cannot parse payload",
			Status:  200,
		})
		return
	}
	clients := WS_CONNECTIONS[data["target_hash"].(string)]
	for _, client := range clients {
		_ = client.Connection.WriteJSON(sendMessageResponse{
			Message: data["message"].(string),
			Sender:  userHash,
			SentAt:  time.Now().Format("YYYY-MM-DD hh:mm:ss"),
		})
	}
}
