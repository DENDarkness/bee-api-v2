package handler

import (
	"bee-api-v2/internal/bee"
)

//var _ bee.ServiceApp = &Handler{}

type Handler struct {
	svc *bee.Service
}

func NewHandler(s *bee.Service) *Handler {
	return &Handler{
		svc: s,
	}
}