package router

import (
	"bee-api-v2/internal/api/handler"
	"bee-api-v2/internal/bee"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(s bee.Service) http.Handler {
	var r = gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v2")
	h := handler.NewHandler(&s)

	api.GET("/node/new", h.ModemReboot)
	return r
}
