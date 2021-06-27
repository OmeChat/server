package websocket

import "github.com/gofiber/websocket/v2"

type ConnectionPair struct {
	UserHash   string
	ClientHash string
	Connection *websocket.Conn
}

type ConnectionIdentifier struct {
	Connection *websocket.Conn
	ClientHash string
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  int    `json:"status"`
}

type WebsocketRequest struct {
	Action      string      `json:"action"`
	UserHash    string      `json:"user_hash"`
	ClientHash  string      `json:"client_hash"`
	AccessToken string      `json:"access_token"`
	Payload     interface{} `json:"payload"`
}
