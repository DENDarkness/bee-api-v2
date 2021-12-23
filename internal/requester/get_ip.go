package requester

import (
	"fmt"
	"net/http"
)

func (r *Request) GetIP(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating a new request failed [GetIP method]: %w", err)
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed [GetIP method]: %w", err)
	}
	// defer resp.Body.Close()

	return resp, nil
}
