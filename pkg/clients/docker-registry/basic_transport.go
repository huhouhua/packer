package docker_registry

import (
	"net/http"
	"strings"
)

type BasicTransport struct {
	Transport http.RoundTripper
	URL       string
	Username  string
	Password  string
}

func NewBasicTransport(URL string, username string, password string) *BasicTransport {
	return &BasicTransport{

		URL:      URL,
		Username: username,
		Password: password}
}

func (t *BasicTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if strings.HasPrefix(req.URL.String(), t.URL) {
		if t.Username != "" || t.Password != "" {
			req.SetBasicAuth(t.Username, t.Password)
		}
	}
	resp, err := t.Transport.RoundTrip(req)
	return resp, err
}
