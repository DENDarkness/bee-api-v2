package server

import (
	"bee-api-v2/internal/config"
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	httpServer *http.Server
	logger *zap.SugaredLogger
}

func NewServer(h http.Handler, cfg *config.Cfg, logger *zap.SugaredLogger) *Server {
	var s = &Server{}

	s.httpServer = &http.Server{
		Addr:              cfg.HTTP.Server.Addr,
		Handler:           h,
		ReadTimeout:       time.Second * cfg.HTTP.Server.ReadTimeout,
		WriteTimeout:      time.Second * cfg.HTTP.Server.WriteTimeout,
		ReadHeaderTimeout: time.Second * cfg.HTTP.Server.ReadHeaderTimeout,
	}

	s.logger = logger

	return s
}

func (srv *Server) Start() {
	go func() {
		if err := srv.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.logger.Fatalf("listen: %s\n", err)
		}
	}()
}

func (srv *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.httpServer.Shutdown(ctx); err != nil {
		srv.logger.Fatalf("server forced to shutdown: %v", err)
	}
	srv.logger.Info("Server exiting")
}
