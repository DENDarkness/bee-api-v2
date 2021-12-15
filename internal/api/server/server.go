package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(h http.Handler) *Server {
	var s = &Server{}

	s.httpServer = &http.Server{
		Addr: ":7777",
		Handler: h,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 10,
		ReadHeaderTimeout: time.Second * 10,
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
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	if err := srv.httpServer.Shutdown(ctx); err != nil {
		//TODO: Обработать ошибку
		log.Fatalf("Failed to shutdown http server gracefully: %v", err)
	}
}
