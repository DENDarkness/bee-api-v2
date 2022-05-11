package bee

import "fmt"

func (svc *Service) InternetTurnOff() error {

	url := svc.cfg.Modem.Host + svc.cfg.Modem.PathInternetSwitch
	body := "<request><dataswitch>0</dataswitch></request>"

	resp, err := svc.req.Post(url, body)
	if err != nil {
		return fmt.Errorf("InternetOff: PostRequest: %w", err)
	}
	defer resp.Body.Close()

	return nil
}
