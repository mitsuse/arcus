package pushbullet

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Get the devices thath can be pushed to.
func (pb *Pushbullet) GetDevices() (*DevicesRes, error) {
	// TODO: Implement this.
	req, err := http.NewRequest("GET", ENDPOINT_DEVICES, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(pb.token, "")

	// TODO: Set the timeout.
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// TODO: Return an error value with human friendly message.
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	var devicesRes *DevicesRes

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&devicesRes); err != nil {
		return nil, err
	}

	return devicesRes, nil
}

type DevicesRes struct {
	Devices []*Device `json:"devices"`
}

type Device struct {
	Iden         string  `json:"iden"`
	PushToken    string  `json:"push_token"`
	FinderPrint  string  `jsonL"fingerprint"`
	Nickname     string  `json:"nickname"`
	Manufacturer string  `json:"manufacturer"`
	Type         string  `json:"type"`
	Model        string  `json:"model"`
	AppVersion   int     `json:"app_version"`
	Created      float64 `json:"created"`
	Modified     float64 `json:"modified"`
	Active       bool    `json:"active"`
	Pushable     bool    `json:"pushable"`
}
