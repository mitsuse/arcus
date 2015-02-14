package pushbullet

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

func (pb *Pushbullet) PostUploadRequest(name, mime string) error {
	values := url.Values{
		"file_name": []string{name},
		"file_type": []string{mime},
	}
	reader := strings.NewReader(values.Encode())

	req, err := http.NewRequest("POST", ENDPOINT_UPLOADREQ, reader)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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

	// TODO: Return the JSON-fromatted response.
	return nil
}
