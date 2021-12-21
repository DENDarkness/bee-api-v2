package handler

import (
	"bee-api-v2/internal/bee"
)

// var _ bee.ServiceApp = &Handler{}

type Handler struct {
	svc bee.ServiceApp
}

func NewHandler(s bee.ServiceApp) *Handler {
	return &Handler{
		svc: s,
	}
}
