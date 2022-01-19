package bee

import (
	"bytes"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func (svc *Service) ModemReboot() error {
	if svc.isReboot {
		return nil
	}

	svc.isReboot = true

	bodyReboot := svc.cfg.Modem.BodyReboot
	urlReboot := svc.cfg.Modem.Host + svc.cfg.Modem.PathReboot

	resp, err := svc.req.Post(urlReboot, bodyReboot)
	if err != nil {
		msgErr := fmt.Errorf("ModemReboot: %w", err)
		svc.logger.Warn(msgErr.Error())
		svc.isReboot = false
		return msgErr
	}
	defer resp.Body.Close()

	go svc.rebootControl()

	return nil
}

func (svc *Service) rebootControl() {
	time.Sleep(25 * time.Second)

	homeModem := svc.cfg.Modem.Host + svc.cfg.Modem.PathHome
	checkHost := svc.cfg.Modem.CheckHost

	status := rebootCheck(checkHost)
	if !status {
		code, err := healthCheck(homeModem)
		if err != nil || code != 200 {
			svc.logger.Warnf("rebootControl : no connection to the modem [StatusCode: %d]: %v", code, err)
		} else {
			code, err := healthCheck(checkHost)
			if err != nil || code != 200 {
				svc.logger.Warnf("rebootControl : no internet access [StatusCode: %d]: %v", code, err)
			}
		}
	}

	url := svc.cfg.URL.GetIP
	go svc.getIP(url)

	svc.isReboot = false
}

func rebootCheck(checkHost string) bool {
	var status = false

	for i := 0; i < 10; i++ {
		_, err := healthCheck(checkHost)
		if err != nil {
			time.Sleep(time.Second * 3)
			continue
		}
		status = true
		break
	}

	return status
}

func (svc *Service) getIP(url string) {
	resp, err := svc.req.GetIP(url)
	if err != nil {
		svc.logger.Warnf("getIP: %v", err)
		svc.cache.Delete("ip")
		return
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		svc.logger.Warnf("getIP: %v", err)
		svc.cache.Delete("ip")
		return
	}
	ip := buf.String()

	svc.cache.Set("ip", ip, cache.DefaultExpiration)
}
