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
	var url = "http://192.168.8.1/api/webserver/SesTokInfo"

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
