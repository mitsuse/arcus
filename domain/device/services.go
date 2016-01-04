package device

import (
	"fmt"

	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/responses"
)

// `ListPushable` lists the pushable devices.
func ListPushable(token string) ([]*Type, error) {
	f := func(r *responses.Device) bool {
		return r.Pushable
	}

	return find(token, f)
}

// `GetPushableByName` returns the pushable device which has the same name as the given one.
func GetPushableByName(token, name string) (*Type, error) {
	devices, err := ListPushable(token)
	if err != nil {
		return nil, err
	}

	for _, d := range devices {
		if d.Name() != name {
			continue
		}

		return d, nil
	}

	err = &NotFound{
		name: name,
	}

	return nil, err
}

type filter func(*responses.Device) bool

func find(token string, f filter) ([]*Type, error) {
	response, err := pushbullet.New(token).GetDevices()
	if err != nil {
		return nil, err
	}

	devices := make([]*Type, 0, len(response))
	for _, r := range response {
		if f != nil && !f(r) {
			continue
		}

		devices = append(devices, covertFrom(r))
	}

	return devices, nil
}

// `NotFound` is an error type to represent that no device exists which has same `name` as the given one.
type NotFound struct {
	name string
}

func (err *NotFound) Name() string {
	return err.name
}

func (err *NotFound) Error() string {
	return fmt.Sprintf("%#v", err)
}
