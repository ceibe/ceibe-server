package services

import (
	"errors"
	"github.com/ceibe/ceibe-server/core/messages"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	userServiceInstance *UserService
)

type UserService struct {
}

func GetUserService() *UserService {
	if userServiceInstance == nil {
		userServiceInstance = new(UserService)
	}
	return userServiceInstance
}

func (us *UserService) CreateUser(newUser *messages.NewUserRequest) (*messages.UserBasicInformation, error) {
	validate := validator.New()

	if err := validate.Struct(newUser); err != nil {
		return nil, err
	}

	return nil, errors.New("not implemented yet")
}