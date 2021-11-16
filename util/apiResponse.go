package util

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}


func ApiResponse(c *fiber.Ctx, statusCode int, msg string, data interface{}) error {
	resp := &Response{
		Status:  statusCode,
		Message: msg,
		Data:    data,
	}

	c.Status(statusCode)
	c.Set("Content-type","application/json")
	return c.JSON(resp)
}
