package launcher

import (
	"bee-api-v2/internal/api/router"
	"bee-api-v2/internal/api/server"
	"bee-api-v2/internal/bee"
	"bee-api-v2/internal/config"
	"bee-api-v2/internal/logger"
	"bee-api-v2/internal/repository/memcache"
	"bee-api-v2/internal/requester"
	"context"
	"os"
	"os/signal"

	"github.com/patrickmn/go-cache"
)

type App struct {
	cfg *config.Cfg
}

func NewApp(cfg *config.Cfg) *App {
	return &App{
		cfg: cfg,
	}
}

func (app *App) Launch() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Created logger
	l := logger.NewLogger()

	// Create memory cache
	c := cache.New(app.cfg.Cache.DefaultExpiration, app.cfg.Cache.CleanupInterval)

	ms := memcache.NewMemStore(c)
	// Created requester
	req := requester.NewRequest(app.cfg)
	// Created core
	a := bee.New(req, ms, app.cfg, l)
	// Created router
	h := router.NewRouter(a, app.cfg, l)

	s := server.NewServer(h, app.cfg, l)

	s.Start()
	<-ctx.Done()
	cancel()

	l.Info("shutting down gracefully, press Ctrl+C again to force")
	s.Stop()

}
