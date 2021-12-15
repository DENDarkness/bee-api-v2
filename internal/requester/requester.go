package requester

import (
	"bee-api-v2/internal/bee"
	"net/http"
	"time"
)

var _ bee.Requester = &Request{}

type Request struct {
	client *http.Client
}

func NewRequest() *Request {
	client := &http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar:     nil,
		Timeout: time.Second,
	}

	return &Request{
		client: client,
	}
}