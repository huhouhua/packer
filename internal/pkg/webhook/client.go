package webhook

import (
	"encoding/json"
	"net/http"
	"strings"
)

type IClient interface {
	Notification(url string, request NotificationRequest) error
}

type Client struct {
	client *http.Client
}

func NewClient() IClient {
	return &Client{
		&http.Client{Transport: http.DefaultTransport},
	}
}

func (h *Client) Notification(url string, request NotificationRequest) error {
	data, err := json.Marshal(request)
	if err != nil {
		return err
	}
	payload := strings.NewReader(string(data))
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := h.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}
