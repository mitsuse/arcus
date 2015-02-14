/*
Package "pushes" provides an interface and structs to represent types of "push".

See the API documentation for the details: https://docs.pushbullet.com/#pushes
*/
package pushes

import (
	"io"
)

type Push interface {
	Dump(writer io.Writer) error
	Type() string
}
