package user

import (
	"github.com/OmeChat/server/internal/storage"
	"github.com/OmeChat/server/internal/webserver/models"
	"github.com/gofiber/fiber/v2"
)

type createAccountRequest struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type createAccountResponse struct {
	Message       string `json:"message"`
	AccountSecret string `json:"account_secret"`
	Status        int    `json:"status"`
}

func CreateAccount(ctx *fiber.Ctx) error {

	req := new(createAccountRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(400).JSON(models.RequestError{
			Error:  err.Error(),
			Status: 400,
		})
	}
	if (storage.Service{}).CheckIfUserExists(req.Username) {
		return ctx.Status(400).JSON(models.RequestError{
			Error:  "This username is already taken",
			Status: 400,
		})
	}
	user := storage.UserModel{}.CreateUserAccount(req.Username, req.Age)
	return ctx.JSON(createAccountResponse{
		Message:       "Successfully created new user account",
		AccountSecret: user.Secret,
		Status:        200,
	})
}
