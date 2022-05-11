package bee

import "fmt"

func (svc *Service) InternetTurnOn() error {

	url := svc.cfg.Modem.Host + svc.cfg.Modem.PathInternetSwitch
	body := "<request><dataswitch>1</dataswitch></request>"

	resp, err := svc.req.Post(url, body)
	if err != nil {
		return fmt.Errorf("InternetOn: PostRequest: %w", err)
	}
	defer resp.Body.Close()

	return nil
}
