package requester

import (
	"bee-api-v2/internal/bee"
	"bee-api-v2/internal/config"
	"net"
	"net/http"
	"time"
)

var _ bee.Requester = &Request{}

type Request struct {
	client *http.Client
	cfg    *config.Cfg
}

func NewRequest(cfg *config.Cfg) *Request {
	// TODO: Нужно привести в порядок клиент
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			Proxy:             http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
			IdleConnTimeout:       120 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
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
