package controllers

import (
	"github.com/ceibe/ceibe-server/core/messages"
	"github.com/ceibe/ceibe-server/core/services"
	"github.com/labstack/echo"
	"net/http"
)

func SetupUserController(e *echo.Echo) {
	e.POST("/users", createUser)
}

func createUser(c echo.Context) (err error) {
	newUserRequest := new(messages.NewUserRequest)
	if err = c.Bind(newUserRequest); err != nil {
		return
	}

	userCreated, err := services.GetUserService().CreateUser(newUserRequest)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, userCreated)
}