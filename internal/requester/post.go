package requester

import (
	"bytes"
	"fmt"
	"net/http"
)

func (r *Request) Post(url string, body string) (*http.Response, error) {
	token, err := r.getToken()
	if err != nil {
		return nil, fmt.Errorf("post: %w", err)
	}

	buffer := bytes.NewBufferString(body)

	req, err := http.NewRequest("POST", url, buffer)
	req.Header.Set("Cookie", token.SesInfo)
	req.Header.Set("__RequestVerificationToken", token.TokInfo)

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post: %w", err)
	}

	return resp, nil
}
