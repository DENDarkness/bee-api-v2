package router

import (
	"bee-api-v2/internal/api/handler"
	"bee-api-v2/internal/bee"
	"bee-api-v2/internal/config"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter(s bee.ServiceApp, cfg *config.Cfg, logger *zap.Logger) *gin.Engine {
	mode := cfg.Bee.Mode
	switch mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	// Disable Console Color
	gin.DisableConsoleColor()

	r := gin.New()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	h := handler.NewHandler(s)

	api := r.Group("/api/v2")
	api.GET("/node/new", h.ModemReboot)
	api.GET("/node/ip", h.GetIP)

	return r
}
