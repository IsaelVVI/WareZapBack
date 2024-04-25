package controllers

import (
	"net/http"

	"github.com/IsaelVVI/warezapback.git/controllers"
	"github.com/IsaelVVI/warezapback.git/schemas"
	"github.com/IsaelVVI/warezapback.git/services"
	"github.com/gin-gonic/gin"
)

// swagger response

// @BasePath /api/v1/auth

// @Summary Login
// @Description Create a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Request body"
// @Succcess 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
func Login(ctx *gin.Context) {
	db, logger = controllers.GetConfigs()

	request := LoginRequest{}

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		logger.Errorf("login error: %v", err.Error())
		controllers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{}

	dbError := db.Where("email = ?", request.Email).First(&user).Error

	if dbError != nil {
		logger.Errorf("Cannot find user: %v", dbError.Error())
		controllers.SendError(ctx, http.StatusBadRequest, "User not found")
		return
	}

	// fmt.Printf(request.Password)

	compare_pass := services.BcryptCompare(request.Password, user.Password)

	if !compare_pass {
		logger.Errorf("email or password incorrect")
		controllers.SendError(ctx, http.StatusBadRequest, "Email or Password are incorrect")
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)

	if err != nil {
		logger.Errorf("internal error: %v", err.Error())
		controllers.SendError(ctx, http.StatusInternalServerError, "Server Error")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "operation from handler: login successfull",
		"access_token": token,
	})

}
