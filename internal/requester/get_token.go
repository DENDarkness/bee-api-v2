package requester

import (
	"encoding/xml"
	"fmt"
)

type TokenModem struct {
	SesInfo string `xml:"SesInfo"`
	TokInfo string `xml:"TokInfo"`
}

func (r *Request) getToken() (*TokenModem, error) {
	var token = TokenModem{}
	var url = r.cfg.Modem.Host + r.cfg.Modem.PathToken

	resp, err := r.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request execution failed [getToken method]: %w", err)
	}
	defer resp.Body.Close()

	if err = xml.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, fmt.Errorf("token decode failed [getToken method]: %w", err)
	}

	return &token, nil
}
