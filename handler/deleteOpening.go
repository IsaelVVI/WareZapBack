package handler

import (
	"fmt"
	"net/http"

	"github.com/IsaelVVI/warezapback.git/schemas"
	"github.com/gin-gonic/gin"
)

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
