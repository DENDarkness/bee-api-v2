package launcher

import (
	"bee-api-v2/internal/api/router"
	"bee-api-v2/internal/api/server"
	"bee-api-v2/internal/bee"
	"bee-api-v2/internal/config"
	"bee-api-v2/internal/repository/memcache"
	"bee-api-v2/internal/requester"
	"context"
	"os"
	"os/signal"

	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
)

type App struct {
	cfg      *config.Cfg
	logger   *zap.SugaredLogger
	memCache *cache.Cache
}

func NewApp(cfg *config.Cfg, logger *zap.SugaredLogger, cache *cache.Cache) *App {
	return &App{
		cfg:      cfg,
		logger:   logger,
		memCache: cache,
	}
}

func (app *App) Launch() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	ms := memcache.NewMemStore(app.memCache)
	// Created requester
	req := requester.NewRequest(app.cfg)
	// Created core
	a := bee.New(req, ms, app.cfg, app.logger)
	// Created router
	h := router.NewRouter(a)

	s := server.NewServer(h, app.cfg, app.logger)

	s.Start()
	<-ctx.Done()
	cancel()

	app.logger.Info("shutting down gracefully, press Ctrl+C again to force")
	s.Stop()

}
