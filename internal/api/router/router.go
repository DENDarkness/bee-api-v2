package router

import (
	"bee-api-v2/internal/api/handler"
	"bee-api-v2/internal/bee"
	"bee-api-v2/internal/config"

	"github.com/gin-gonic/gin"
)

func NewRouter(s bee.ServiceApp, cfg *config.Cfg) *gin.Engine {
	mode := cfg.Bee.Mode
	switch mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	gin.DisableConsoleColor()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	h := handler.NewHandler(s)

	api := r.Group("/api/v2")
	api.GET("/node/new", h.ModemReboot)
	api.GET("/node/ip", h.GetIP)

	return r
}
