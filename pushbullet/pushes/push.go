package pushes

import (
	"io"
)

type Push interface {
	Dump(writer io.Writer) error
	Type() string
}
