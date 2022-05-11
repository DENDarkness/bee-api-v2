package bee

import (
	"net/http"
	"time"
)

type ServiceApp interface {
	ModemReboot() error
	GetIP() (interface{}, error)
	InternetTurnOn() error
	InternetTurnOff() error
	USSDSend() error
	USSDGet() error
	GetDeviceInformation() (*DeviceInformation, error)
}

type Requester interface {
	Get(url string) (*http.Response, error)
	Post(url, body string) (*http.Response, error)
	GetIP(url string) (*http.Response, error)
}

type Repository interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, d time.Duration)
	Delete(key string)
}
