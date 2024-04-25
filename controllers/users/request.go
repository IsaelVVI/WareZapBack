package controllers

import (
	"fmt"

	"github.com/IsaelVVI/warezapback.git/models"
)

// login struc
type LoginRequest struct {
	models.Login
}

// Create User
type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request *CreateUserRequest) Validate() error {
	if request.Name == "" &&
		request.Email == "" &&
		request.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if request.Name == "" {
		return errParamIsRequired("name", "string")
	}

	if request.Email == "" {
		return errParamIsRequired("email", "string")
	}

	if request.Password == "" {
		return errParamIsRequired("password", "string")
	}

	return nil
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
