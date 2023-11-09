package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Get("/", Sample)
}

func Sample(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).
		JSON(GetResponse(fiber.Map{"id": 1, "name": "john doe"}, "", 400, errors.New("asdasd")))
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type DataResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type JSONResponse struct {
	Success bool          `json:"ok"`
	Data    DataResponse  `json:"data"`
	Error   ErrorResponse `json:"error"`
}

func GetResponse(data interface{}, message string, status int, err error) JSONResponse {
	var response JSONResponse
	if status < 400 {
		response.Success = true
	}
	response.Data = DataResponse{
		Message: message,
		Data:    data,
	}
	response.Error = ErrorResponse{
		Code:    status,
		Message: err.Error(),
	}

	return response
}
