package user

import (
	"github.com/OmeChat/server/internal/storage"
	"github.com/OmeChat/server/internal/webserver/models"
	"github.com/gofiber/fiber/v2"
)

type addClientRequest struct {
	Username      string `json:"username"`
	AccountSecret string `json:"account_secret"`
}

type addClientResponse struct {
	Message     string `json:"message"`
	UserHash    string `json:"user_hash"`
	ClientHash  string `json:"client_hash"`
	AccessToken string `json:"access_token"`
	Status      int    `json:"status"`
}

func AddClient(ctx *fiber.Ctx) error {

	req := new(addClientRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(400).JSON(models.RequestError{
			Error:  err.Error(),
			Status: 400,
		})
	}
	if !(storage.Service{}).CheckUserAccess(req.Username, req.AccountSecret) {
		return ctx.Status(400).JSON(models.RequestError{
			Error:  "The given account secret is wrong",
			Status: 400,
		})
	}
	usrHash := storage.UserModel{}.GetHashByUsername(req.Username)
	clientHash, client := storage.ClientModel{}.AddClient(usrHash)
	storage.UserModel{}.AddClientToUser(usrHash, clientHash)

	return ctx.JSON(addClientResponse{
		"Successfully added client to server",
		usrHash,
		clientHash,
		client.AccessToken,
		200,
	})
}
