package router

import (
	"bee-api-v2/internal/api/handler"
	"bee-api-v2/internal/api/middleware"
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
	r.Use(middleware.Auth(cfg.Bee.Token))

	h := handler.NewHandler(s)

	old := r.Group("/v2")
	old.GET("/reset/new", h.ModemReboot)

	api := r.Group("/api/v2")
	api.GET("/node/new", h.ModemReboot)
	api.GET("/node/ip", h.GetIP)
	// api.GET("/node/device/information", h.GetDeviceInformation)
	// api.POST("node/ussd/send", h.USSDSend)
	// api.GET("node/ussd/get", h.USSDGet)

	api.POST("/node/internet/off", h.InternetTurnOff)
	api.POST("/node/internet/on", h.InternetTurnOn)

	return r
}
