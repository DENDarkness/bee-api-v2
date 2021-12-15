package bee

import "net/http"

type ServiceApp interface {
	ModemReboot() error
}

type Requester interface {
	Get(url string) (*http.Response, error)
	Post(url, body string) (*http.Response, error)
}
