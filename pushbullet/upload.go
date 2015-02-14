package pushbullet

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
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

// Upload a file to S3 specified with the response of PostUploadRequest.
func Upload(upload *UploadReqRes, reader io.Reader) error {
	buffer, err := createMultipart(upload, reader)
	if err != nil {
		return err
	}

	// TODO: Post the upload request.
	_ = buffer

	return nil
}

func createMultipart(upload *UploadReqRes, reader io.Reader) (*bytes.Buffer, error) {
	dest := upload.Data

	buffer := &bytes.Buffer{}

	writer := newMultipartWriter(buffer)
	defer writer.Close()

	writer.WriteField("awsaccesskeyid", dest.AwsAccessKey)
	writer.WriteField("acl", dest.Acl)
	writer.WriteField("key", dest.Key)
	writer.WriteField("signature", dest.Signature)
	writer.WriteField("policy", dest.Policy)
	writer.WriteField("content-type", dest.ContentType)

	if err := writer.Error(); err != nil {
		return nil, err
	}

	fw, err := writer.CreateFormFile("file", upload.FileName)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(fw, reader); err != nil {
		return nil, err
	}

	return buffer, nil
}
