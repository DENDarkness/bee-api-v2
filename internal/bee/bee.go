package bee

import (
	"fmt"
	"net/http"
	"time"
)

type Service struct {
	req      Requester
	isReboot bool
}

func New(r Requester) *Service {
	return &Service{
		req:      r,
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