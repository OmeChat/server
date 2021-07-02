package user

import (
	"github.com/OmeChat/server/internal/storage"
	"github.com/OmeChat/server/internal/webserver/models"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type getUsernamesRequest struct {
	UserHash    string `json:"user_hash"`
	ClientHash  string `json:"client_hash"`
	AccessToken string `json:"access_token"`
	Usernames   string `json:"usernames"`
}

type getUsernamesResponse struct {
	Usernames map[string]string `json:"usernames"`
}

func GetUsernames(ctx *fiber.Ctx) error {

	obj := getUsernamesRequest{
		UserHash:    ctx.Query("user_hash", ""),
		ClientHash:  ctx.Query("client_hash", ""),
		AccessToken: ctx.Query("access_token", ""),
		Usernames:   ctx.Query("usernames", ""),
	}
	if !validateGetUsernamesRequest(obj) {
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
	split := strings.Split(obj.Usernames, ";")

	var usernames map[string]string
	for _, el := range split {
		usernames[el] = storage.UserModel{}.GetUserByHash(el).Username
	}
	return ctx.JSON(getUsernamesResponse{
		Usernames: usernames,
	})
}

func validateGetUsernamesRequest(obj getUsernamesRequest) bool {
	return obj.UserHash != "" && obj.ClientHash != "" &&
		obj.AccessToken != "" && obj.Usernames != ""
}
