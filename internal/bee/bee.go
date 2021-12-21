package bee

import (
	"bee-api-v2/internal/config"
	"fmt"
	"net/http"
	"time"
)

type Service struct {
	req			Requester
	cfg 		*config.Cfg
	isReboot 	bool
}

func New(r Requester, cfg *config.Cfg) *Service {
	return &Service{
		req:      	r,
		cfg:		cfg,
		isReboot: 	false,
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
		// TODO: Настроить логирование
		return 0, fmt.Errorf("healthCheck: %w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}