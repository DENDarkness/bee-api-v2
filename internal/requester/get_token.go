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
		return nil, fmt.Errorf("getTokenModem: %w", err)
	}
	defer resp.Body.Close()

	if err = xml.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, fmt.Errorf("getTokenModem: %w", err)
	}

	return &token, nil
}
