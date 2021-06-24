package websocket

import (
	"github.com/gofiber/websocket/v2"
)

type exchangeKeyPayload struct {
	TargetHash string
	key        string
}

type exchangeKeyResponse struct {
	Action     string `json:"action"`
	Key        string `json:"key"`
	SenderHash string `json:"sender_hash"`
}

// exchangeKey executes the process of sending the public key for
// the end to end encryption to all clients of the given target
// hash. The state of this process is being returned as a
// boolean value
func exchangeKey(c *websocket.Conn, payload interface{}, userHash string) {
	data, ok := payload.(exchangeKeyPayload)
	if !ok {
		_ = c.WriteJSON(ErrorResponse{
			Message: "Invalid payload",
			Error:   "Cannot parse payload",
			Status:  200,
		})
		return
	}
	conns := WS_CONNECTIONS[data.TargetHash]
	for _, el := range conns {
		_ = el.WriteJSON(exchangeKeyResponse{
			Action:     "exchange-key",
			Key:        data.key,
			SenderHash: userHash,
		})
	}
}
