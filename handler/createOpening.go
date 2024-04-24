package handler

import (
	"net/http"

	"github.com/IsaelVVI/warezapback.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Opening
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Succcess 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /teste [post]
func CreateOpeningHandler(ctx *gin.Context) {
	// verify body request with struc creat opening
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "erro creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
