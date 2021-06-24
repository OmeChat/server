package actions

import (
	ws "github.com/OmeChat/server/internal/websocket"
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

// ExchangeKey executes the process of sending the public key for
// the end to end encryption to all clients of the given target
// hash. The state of this process is being returned as a
// boolean value
func ExchangeKey(c *websocket.Conn, payload interface{}, userHash string) {
	data, ok := payload.(exchangeKeyPayload)
	if !ok {
		_ = c.WriteJSON(ws.ErrorResponse{
			Message: "Invalid payload",
			Error:   "Cannot parse payload",
			Status:  200,
		})
		return
	}
	conns := ws.WS_CONNECTIONS[data.TargetHash]
	for _, el := range conns {
		_ = el.WriteJSON(exchangeKeyResponse{
			Action:     "exchange-key",
			Key:        data.key,
			SenderHash: userHash,
		})
	}
}
