package resource

import (
	"mime"
	"os"
	"path"

	"github.com/mitsuse/pushbullet-go"
)

// `Upload` uploads the given file and return the information of uploaded resource such as the URL.
func Upload(token string, f *os.File) (*Remote, error) {
	client := pushbullet.New(token)

	name := f.Name()
	mimeType := mime.TypeByExtension(path.Ext(name))

	response, err := client.PostUploadRequest(name, mimeType)
	if err != nil {
		return nil, err
	}

	if err := pushbullet.Upload(client.Client(), response, f); err != nil {
		return nil, err
	}

	t := &Remote{
		name: response.FileName,
		t:    response.FileType,
		url:  response.FileUrl,
	}

	return t, nil
}
