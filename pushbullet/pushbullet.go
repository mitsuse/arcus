/*
Package "pushbullet" provides interfaces for Pushbullet HTTP API.

Pushbullet is a web service,
which makes your devices work better together by allowing you to move things between them easily.

The official url: https://www.pushbullet.com/

Currently, this package supports only "pushes" except file.

See the API documentation for the details: https://docs.pushbullet.com/#http
*/
package pushbullet

import (
	"bytes"
	"errors"
	"net/http"

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

func (pb *Pushbullet) PostPushesNote(n *pushes.Note) error {
	return pb.postPushes(n)
}

func (pb *Pushbullet) PostPushesLink(l *pushes.Link) error {
	return pb.postPushes(l)
}

func (pb *Pushbullet) PostPushesAddress(a *pushes.Address) error {
	return pb.postPushes(a)
}

func (pb *Pushbullet) PostPushesCheck(c *pushes.Check) error {
	return pb.postPushes(c)
}

func (pb *Pushbullet) postPushes(p pushes.Push) error {
	buffer := &bytes.Buffer{}

	if err := p.Dump(buffer); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", ENDPOINT_PUSHES, buffer)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(pb.token, "")

	// TODO: Set the timeout.
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// TODO: Return an error value with human friendly message.
	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}

	return nil
}
