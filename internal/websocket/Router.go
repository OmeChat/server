package websocket

import (
	"fmt"
	"github.com/OmeChat/server/internal/storage"
	"github.com/gofiber/websocket/v2"
)

// The Router handles the action-requests
// of the connected clients and executes them
func Router(c *websocket.Conn) {
	req := new(WebsocketRequest)
	for {
		if err := c.ReadJSON(req); err != nil {
			err = c.WriteJSON(ErrorResponse{
				Message: "invalid request payload",
				Error:   err.Error(),
				Status:  200,
			})
			if err != nil {
				fmt.Println(err.Error())
			}
			break
		}
		if !storage.CheckAuthStatus(req.UserHash, req.ClientHash, req.AccessToken) {
			err := c.WriteJSON(ErrorResponse{
				Message: "login failed",
				Error:   "The given login credentials are wrong",
				Status:  200,
			})
			if err != nil {
				fmt.Println(err.Error())
			}
			break
		}
		WS_DATAFLOW_CHANNEL <- ConnectionPair{req.UserHash, req.ClientHash, c}
		switch req.Action {
		case "exchange-key":
			exchangeKey(c, req.Payload, req.UserHash)
		case "request-random-people":
			requestRandomPeople(c, req.UserHash, req.Payload)
		case "send-message":
			sendMessage(c, req.UserHash, req.Payload)
		default:
			err := c.WriteJSON(ErrorResponse{
				Message: "unknown action",
				Error:   "unknown action",
				Status:  200,
			})
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
