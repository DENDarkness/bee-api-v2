package requester

import (
	"fmt"
	"net/http"
)

func (r *Request) Get(url string) (*http.Response, error) {
	token, err := r.getToken()
	if err != nil {
		return nil, fmt.Errorf("method get: %w", err)
	}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", token.SesInfo)
	req.Header.Set("__RequestVerificationToken", token.TokInfo)

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed [Get method]: %w", err)
	}

	return resp, nil
}
