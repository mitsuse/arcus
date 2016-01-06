package resource

import (
	"mime"
	"path"
)

// `Local` represents the information of a resource stored at local.
type Local struct {
	name string
	t    string
	path string
}

// `Read` returns the information of local resource located at the given path `p`.
func Read(p string) (*Local, error) {
	base := path.Base(p)
	extension := path.Ext(base)
	mimeType := mime.TypeByExtension(extension)

	l := &Local{
		name: base,
		t:    mimeType,
		path: p,
	}

	return l, nil
}

func (l *Local) Name() string {
	return l.name
}

func (l *Local) Type() string {
	return l.t
}

func (l *Local) Path() string {
	return l.path
}
