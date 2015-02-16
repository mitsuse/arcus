package pushbullet

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mitsuse/bullet/pushbullet/responses"
)

// Get the devices thath can be pushed to.
func (pb *Pushbullet) GetDevices() ([]*responses.Device, error) {
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

	var devices *devicesResponse

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&devices); err != nil {
		return nil, err
	}

	return devices.Devices, nil
}

type devicesResponse struct {
	Devices []*responses.Device `json:"devices"`
}
