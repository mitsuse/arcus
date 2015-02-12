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

func (pb *Pushbullet) PostNote(n *pushes.Note) error {
	buffer := &bytes.Buffer{}

	if err := n.Dump(buffer); err != nil {
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

func (pb *Pushbullet) PostLink(n *pushes.Link) error {
	buffer := &bytes.Buffer{}

	if err := n.Dump(buffer); err != nil {
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
