package websocket

import (
	"github.com/OmeChat/server/internal/storage"
	"github.com/gofiber/websocket/v2"
)

type requestRandomPeopleResponse struct {
	Action string                `json:"action"`
	User   []storage.ExposedUser `json:"user"`
}

// requestRandomPeople requests random people at the same age or at an age
// inside the tolerance. After that it returns a list of them and sends
// them trough the websocket to the client
func requestRandomPeople(c *websocket.Conn, userHash string, payload interface{}) {
	data, ok := PayloadParser("request-random-people", payload)
	if !ok {
		_ = c.WriteJSON(ErrorResponse{
			Message: "Invalid payload",
			Error:   "Cannot parse payload",
			Status:  200,
		})
		return
	}
	user := storage.UserModel{}.GetUserByHash(userHash)
	matchingPeople := storage.UserModel{}.GetUserAtAgeWithTolerance(user.Age, int(data["tolerance"].(float64)), userHash)
	_ = c.WriteJSON(requestRandomPeopleResponse{
		Action: "request-random-people",
		User:   matchingPeople,
	})
}
