package server

import (
	"bee-api-v2/internal/config"
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(h http.Handler, cfg *config.Cfg) *Server {
	var s = &Server{}

	s.httpServer = &http.Server{
		Addr:              cfg.Server.Addr,
		Handler:           h,
		ReadTimeout:       time.Second * cfg.Server.ReadTimeout,
		WriteTimeout:      time.Second * cfg.Server.WriteTimeout,
		ReadHeaderTimeout: time.Second * cfg.Server.ReadHeaderTimeout,
	}

	return s
}

func (srv *Server) Start() {
	if err := srv.httpServer.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			srv.Stop()
			log.Fatalf("Http Server stopped unexpected: %v", err)
		} else {
			log.Println("Http Server stopped")
		}
	}
}

func (srv *Server) Stop() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := srv.httpServer.Shutdown(ctx); err != nil {
		// TODO: Обработать ошибку
		log.Fatalf("Failed to shutdown http server gracefully: %v", err)
	}
}
