package pushbullet

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mitsuse/bullet/pushbullet/pushes"
)

// Push a note, which consists of "title" and "message" strings.
func (pb *Pushbullet) PostPushesNote(n *pushes.Note) error {
	return pb.postPushes(n)
}

// Push a link, which consists of "title", "message" and "url" strings.
func (pb *Pushbullet) PostPushesLink(l *pushes.Link) error {
	return pb.postPushes(l)
}

// Push an address, which consists of the place "name" and "address (searchquery)" for map.
func (pb *Pushbullet) PostPushesAddress(a *pushes.Address) error {
	return pb.postPushes(a)
}

// Push a checklist, which consists of "title" and the list of items.
func (pb *Pushbullet) PostPushesChecklist(c *pushes.Checklist) error {
	return pb.postPushes(c)
}

// Push a file, which consists of "title", "message" and the information of uploaded file.
func (pb *Pushbullet) PostPushesFile(f *pushes.File) error {
	return pb.postPushes(f)
}

// func (pb *Pushbullet) postPushes(p pushes.Push) error {
func (pb *Pushbullet) postPushes(p interface{}) error {
	buffer := &bytes.Buffer{}

	encoder := json.NewEncoder(buffer)
	if err := encoder.Encode(p); err != nil {
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
