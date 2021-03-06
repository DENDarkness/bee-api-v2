package bee

import (
	"bytes"

	"github.com/patrickmn/go-cache"
)

func (svc *Service) GetIP() (interface{}, error) {
	ip, found := svc.cache.Get("ip")
	if found {
		return ip, nil
	}

	url := svc.cfg.URL.GetIP
	resp, err := svc.req.GetIP(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	ip = buf.String()

	svc.logger.Sugar().Infof("IPAddress: %v", ip)

	svc.cache.Set("ip", ip, cache.DefaultExpiration)

	return ip, nil
}
