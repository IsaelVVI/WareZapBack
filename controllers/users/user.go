package controllers

import (
	"net/http"

	"github.com/IsaelVVI/warezapback.git/config"
	"github.com/IsaelVVI/warezapback.git/controllers"
	"github.com/IsaelVVI/warezapback.git/schemas"
	"github.com/IsaelVVI/warezapback.git/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

// swagger response

// @BasePath /api/v1/user

// @Summary Create user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Request body"
// @Succcess 200 {object} CreateUserRequest
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /create [post]
func CreatedUser(ctx *gin.Context) {

	db, logger = controllers.GetConfigs()

	request := CreateUserRequest{}

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		controllers.SendError(ctx, http.StatusBadRequest, "error in register user")

		return
	}

	password := services.BcryptEncoder(request.Password)

	if password == "error" {
		logger.Errorf("password error: %v", err.Error())
		controllers.SendError(ctx, http.StatusBadRequest, "password not accepted")
		return
	}

	user := schemas.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
	}

	err = db.Create(&user).Error

	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		controllers.SendError(ctx, http.StatusBadRequest, "error in register user")

		return
	}

	controllers.SendSuccess(ctx, "user", "user created")

}
