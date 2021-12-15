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

	svc.isReboot = false

	var body string = "<request><Control>1</Control></request>"
	var url string = "http://192.168.8.1/api/device/control"

	resp, err := svc.req.Post(url, body)
	if err != nil {
		fmt.Errorf("ModemReboot: %w", err)
	}
	defer resp.Body.Close()

	go rebootControl(&svc.isReboot)

	return nil
}

func rebootControl(isReboot *bool) {
	time.Sleep(25 * time.Second)

	status := rebootCheck()
	if !status {
		var urlModem = "http://192.168.8.1/html/home.html"
		code, err := healthCheck(urlModem)
		if err != nil || code != 200 {
			log.Println(fmt.Errorf("modemRebootState: StatusCode(%d): %w", code, err))
		} else {
			var urlInt = "https://yandex.ru"
			code, err := healthCheck(urlInt)
			if err != nil || code != 200 {
				log.Println(fmt.Errorf("modemRebootState: StatusCode(%d): %w", code, err))
			}
		}
	}

	*isReboot = false
}

func rebootCheck() bool {
	var status = false
	var checkURL = "https://yandex.ru"

	for i := 0; i < 10; i++ {
		_, err := healthCheck(checkURL)
		if err != nil {
			time.Sleep(time.Second * 3)
			continue
		}
		status = true
		break
	}

	return status
}