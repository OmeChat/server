package websocket

import "github.com/gofiber/websocket/v2"

type ConnectionPair struct {
	Hash       string
	Connection *websocket.Conn
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  int    `json:"status"`
}

type WebsocketRequest struct {
	Action      string      `json:"action"`
	UserHash    string      `json:"user_hash"`
	AccessToken string      `json:"access_token"`
	Payload     interface{} `json:"payload"`
}
