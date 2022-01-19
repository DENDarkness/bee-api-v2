package handler

import (
	"bee-api-v2/internal/bee"
)

type Handler struct {
	svc bee.ServiceApp
}

func NewHandler(svc bee.ServiceApp) *Handler {
	return &Handler{
		svc: svc,
	}
}
