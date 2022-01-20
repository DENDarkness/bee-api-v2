package bee

import (
	"fmt"
	"time"
)

func (svc *Service) ModemReboot() error {
	// if variable isReboot is true then error nil is returned
	if svc.isReboot {
		return nil
	}
	// set the value of the isReboot variable to true
	svc.isReboot = true

	bodyReboot := svc.cfg.Modem.BodyReboot
	urlReboot := svc.cfg.Modem.Host + svc.cfg.Modem.PathReboot
	//
	resp, err := svc.req.Post(urlReboot, bodyReboot)
	if err != nil {
		svc.isReboot = false
		return fmt.Errorf("ModemReboot: PostRequest: %w", err)
	}
	defer resp.Body.Close()

	go svc.rebootControl()

	return nil
}

func (svc *Service) rebootControl() {
	time.Sleep(25 * time.Second)
	// clear cache
	svc.cache.Delete("ip")

	homeModem := svc.cfg.Modem.Host + svc.cfg.Modem.PathHome
	checkHost := svc.cfg.Modem.CheckHost

	status := rebootCheck(checkHost)
	if !status {
		code, err := healthCheck(homeModem)
		if err != nil || code != 200 {
			svc.logger.Sugar().Warnf("rebootControl : no connection to the modem [StatusCode: %d]: %v", code, err)
		} else {
			code, err := healthCheck(checkHost)
			if err != nil || code != 200 {
				svc.logger.Sugar().Warnf("rebootControl : no internet access [StatusCode: %d]: %v", code, err)
			}
		}
	}

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

