package websocket

import (
	"github.com/gofiber/websocket/v2"
)

type exchangeKeyResponse struct {
	Action     string `json:"action"`
	Key        string `json:"key"`
	SenderHash string `json:"sender_hash"`
}

// exchangeKey executes the process of sending the public key for
// the end to end encryption to all clients of the given target
// hash.
func exchangeKey(c *websocket.Conn, payload interface{}, userHash string) {
	data, ok := PayloadParser("exchange-key", payload)
	if !ok {
		_ = c.WriteJSON(ErrorResponse{
			Message: "Invalid payload",
			Error:   "Cannot parse payload",
			Status:  200,
		})
		return
	}
	hash, _ := data["target_hash"].(string)
	conns := WS_CONNECTIONS[hash]
	key, _ := data["key"].(string)
	for _, el := range conns {
		_ = el.Connection.WriteJSON(exchangeKeyResponse{
			Action:     "exchange-key",
			Key:        key,
			SenderHash: userHash,
		})
	}
}
