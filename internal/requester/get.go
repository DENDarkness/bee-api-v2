package requester

import (
	"fmt"
	"net/http"
)

func (r *Request) Get(url string) (*http.Response, error) {
	token, err := r.getToken()
	if err != nil {
		return nil, fmt.Errorf("post: %w", err)
	}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", token.SesInfo)
	req.Header.Set("__RequestVerificationToken", token.TokInfo)

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post: %w", err)
	}
	defer resp.Body.Close()

	return resp, nil
}
