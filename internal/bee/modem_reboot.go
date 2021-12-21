package bee

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func (svc *Service) ModemReboot() error {
	if svc.isReboot {
		return nil
	}

	svc.isReboot = true

	var bodyReboot = svc.cfg.Modem.BodyReboot
	var urlReboot = svc.cfg.Modem.Host + svc.cfg.Modem.PathReboot

	resp, err := svc.req.Post(urlReboot, bodyReboot)
	if err != nil {
		return fmt.Errorf("ModemReboot: %w", err)
	}
	defer resp.Body.Close()

	var homeModem = svc.cfg.Modem.Host + svc.cfg.Modem.PathHome
	var checkHost = svc.cfg.Modem.CheckHost

	go rebootControl(&svc.isReboot, homeModem, checkHost, svc.logger)

	return nil
}

func rebootControl(isReboot *bool, homeModem, checkHost string, l *zap.Logger) {
	time.Sleep(25 * time.Second)
	logger := l.Sugar()

	status := rebootCheck(checkHost)
	if !status {
		code, err := healthCheck(homeModem)
		if err != nil || code != 200 {
			logger.Warnf("rebootControl : no connection to the modem [StatusCode: %d]: %v", code, err)
		} else {
			code, err := healthCheck(checkHost)
			if err != nil || code != 200 {
				logger.Warnf("rebootControl : no internet access [StatusCode: %d]: %v", code, err)
			}
		}
	}

	*isReboot = false
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
