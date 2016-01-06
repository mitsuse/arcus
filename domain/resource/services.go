package resource

import (
	"os"

	"github.com/mitsuse/pushbullet-go"
)

// `Upload` uploads the given local resource and return the information of uploaded remote resource.
func Upload(token string, l *Local) (*Remote, error) {
	client := pushbullet.New(token)

	f, err := os.Open(l.Path())
	if err != nil {
		return nil, err
	}
	defer f.Close()

	response, err := client.PostUploadRequest(l.Name(), l.Type())
	if err != nil {
		return nil, err
	}

	if err := pushbullet.Upload(client.Client(), response, f); err != nil {
		return nil, err
	}

	r := &Remote{
		name: response.FileName,
		t:    response.FileType,
		url:  response.FileUrl,
	}

	return r, nil
}
