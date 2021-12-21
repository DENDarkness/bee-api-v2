package requester

import (
	"bee-api-v2/internal/bee"
	"bee-api-v2/internal/config"
	"net/http"
	"time"
)

var _ bee.Requester = &Request{}

type Request struct {
	client *http.Client
	cfg    *config.Cfg
}

func NewRequest(cfg *config.Cfg) *Request {
	client := &http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar:     nil,
		Timeout: time.Second * cfg.HTTP.Client.Timeout,
	}

	return &Request{
		client: client,
		cfg:    cfg,
	}
}
