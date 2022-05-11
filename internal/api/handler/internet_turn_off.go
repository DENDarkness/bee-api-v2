package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InternetTurnOff(ctx *gin.Context) {

	if err := h.svc.InternetTurnOff(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"msg":   nil,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "internet turn off",
		"error": nil,
	})
}
