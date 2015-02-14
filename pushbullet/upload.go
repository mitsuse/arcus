package pushbullet

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

// Post the information of file and get the information of upload destination.
func (pb *Pushbullet) PostUploadRequest(name, mime string) (*UploadReqRes, error) {
	res, err := pb.postUploadRequest(name, mime)
	if err != nil {
		return nil, err
	}

	// TODO: Return an error value with human friendly message.
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	decoder := json.NewDecoder(res.Body)

	var uploadReqRes *UploadReqRes
	if err := decoder.Decode(&uploadReqRes); err != nil {
		return nil, err
	}

	return uploadReqRes, nil
}

func (pb *Pushbullet) postUploadRequest(name, mime string) (*http.Response, error) {
	values := url.Values{
		"file_name": []string{name},
		"file_type": []string{mime},
	}
	reader := strings.NewReader(values.Encode())

	req, err := http.NewRequest("POST", ENDPOINT_UPLOADREQ, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(pb.token, "")

	// TODO: Set the timeout.
	client := &http.Client{}

	return client.Do(req)
}
