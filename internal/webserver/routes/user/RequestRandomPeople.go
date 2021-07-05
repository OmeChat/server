package user

import (
	"github.com/OmeChat/server/internal/storage"
	"github.com/OmeChat/server/internal/webserver/models"
	"github.com/gofiber/fiber/v2"
)

type requestRandomPeopleRequest struct {
	UserHash    string `json:"user_hash"`
	ClientHash  string `json:"client_hash"`
	AccessToken string `json:"access_token"`
	Tolerance   int    `json:"tolerance"`
}

type requestRandomPeopleResponse struct {
	MatchingUser []storage.ExposedUser `json:"matching_user"`
	Status       int                   `json:"status"`
}

func RequestRandomPeople(ctx *fiber.Ctx) error {

	obj := requestRandomPeopleRequest{
		UserHash:    ctx.Query("user_hash", ""),
		ClientHash:  ctx.Query("client_hash", ""),
		AccessToken: ctx.Query("access_token", ""),
	}
	if !validateRequestRandomPeopleRequest(obj) {
		return ctx.JSON(models.RequestError{
			Status: 400,
			Error:  "Invalid request parameters",
		})
	}
	if !storage.CheckAuthStatus(obj.UserHash, obj.ClientHash, obj.AccessToken) {
		return ctx.JSON(models.RequestError{
			Status: 400,
			Error:  "Login credentials are wrong",
		})
	}
	user := storage.UserModel{}.GetUserByHash(obj.UserHash)
	matchingPeople := storage.UserModel{}.GetUserAtAgeWithTolerance(user.Age, obj.Tolerance, obj.UserHash)
	return ctx.JSON(requestRandomPeopleResponse{
		Status:       200,
		MatchingUser: matchingPeople,
	})
}

func validateRequestRandomPeopleRequest(obj requestRandomPeopleRequest) bool {
	return obj.UserHash != "" && obj.ClientHash != "" &&
		obj.AccessToken != "" && obj.Tolerance > 0
}
