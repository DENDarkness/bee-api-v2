package handler

import (
	"bee-api-v2/internal/bee"
)

type Handler struct {
	svc bee.ServiceApp
}

func NewHandler(s bee.ServiceApp) *Handler {
	return &Handler{
		svc: s,
	}
}
