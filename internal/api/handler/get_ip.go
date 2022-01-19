package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetIP(ctx *gin.Context) {
	ip, err := h.svc.GetIP()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"msg": nil,
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg": ip,
		"error": nil,
	})
}
