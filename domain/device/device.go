// Package `device` provides an type to represent the device information and the service to find specific devices.
package device

import (
	"github.com/mitsuse/pushbullet-go/responses"
)

// `Type` is a type to represent the device information.
type Type struct {
	id   string
	name string
}

func (d *Type) Id() string {
	return d.id
}

func (d *Type) Name() string {
	return d.name
}

func covertFrom(r *responses.Device) *Type {
	d := &Type{
		id:   r.Iden,
		name: r.Nickname,
	}

	return d
}
