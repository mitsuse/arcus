package file

import (
	"mime"
	"os"
	"path"

	"github.com/mitsuse/pushbullet-go"
)

// `Upload` uploads the given file and return the information of uploaded file such as the URL.
func Upload(token string, f *os.File) (*Type, error) {
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

	t := &Type{
		name: response.FileName,
		t:    response.FileType,
		url:  response.FileUrl,
	}

	return t, nil
}
