package pushbullet

import (
	"github.com/mitsuse/bullet/pushbullet/pushes"
)

type Pushbullet struct {
	token string
}

func New(token string) *Pushbullet {
	pb := &Pushbullet{
		token: token,
	}

	return pb
}

func (pb *Pushbullet) PostNote(n *pushes.Note) error {
	// TODO: Implement this.
	return nil
}
