package websocket

import (
	"fmt"
	"github.com/gofiber/websocket/v2"
)

// The Router handles the action-requests
// of the connected clients and executes them
func Router(c *websocket.Conn) {
	req := new(WebsocketRequest)
	for {
		if err := c.ReadJSON(req); err != nil {
			_ = c.WriteJSON(ErrorResponse{
				Message: "invalid request payload",
				Error:   err.Error(),
				Status:  200,
			})
			break
		}
		WS_DATAFLOW_CHANNEL <- ConnectionPair{req.UserHash, c}
		switch req.Action {
		case "exchange-key":
			fmt.Println("dass du out bist")
			break
		default:
			err := c.WriteJSON(ErrorResponse{
				Message: "unknown action",
				Error:   "unknown action",
				Status:  200,
			})
			if err != nil {
				_ = c.Close()
			}
			break
		}
	}
}
