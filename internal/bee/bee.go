package bee

import (
	"bee-api-v2/internal/config"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Service struct {
	req      Requester
	cache    Repository
	cfg      *config.Cfg
	logger   *zap.SugaredLogger
	isReboot bool
}

func New(r Requester, c Repository, cfg *config.Cfg, logger *zap.SugaredLogger) *Service {
	return &Service{
		req:      r,
		cache:    c,
		cfg:      cfg,
		logger:   logger,
		isReboot: false,
	}
}

func healthCheck(url string) (int, error) {
	client := &http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar:     nil,
		Timeout: time.Second,
	}

	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return 0, fmt.Errorf("healthCheck: %w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
