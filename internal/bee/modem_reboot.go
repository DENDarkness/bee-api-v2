package bee

import (
	"fmt"
	"log"
	"time"
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
		// TODO: Настроить логирование
		fmt.Errorf("ModemReboot: %w", err)
	}
	defer resp.Body.Close()

	var homeModem = svc.cfg.Modem.Host + svc.cfg.Modem.PathHome
	var checkHost = svc.cfg.Modem.CheckHost

	go rebootControl(&svc.isReboot, homeModem, checkHost)

	return nil
}

func rebootControl(isReboot *bool, homeModem, checkHost string) {
	time.Sleep(25 * time.Second)

	status := rebootCheck(checkHost)
	if !status {
		code, err := healthCheck(homeModem)
		if err != nil || code != 200 {
			// TODO: Настроить логирование
			log.Println(fmt.Errorf("modemRebootState: StatusCode(%d): %w", code, err))
		} else {
			code, err := healthCheck(checkHost)
			if err != nil || code != 200 {
				// TODO: Настроить логирование
				log.Println(fmt.Errorf("modemRebootState: StatusCode(%d): %w", code, err))
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
