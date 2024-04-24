package handler

import (
	"fmt"
	"net/http"

	"github.com/IsaelVVI/warezapback.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Succcess 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /teste [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	// Delete Opening -- Unscoped remove permanently data
	if err := db.Unscoped().Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("error deleleting opening with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)

}
