package pushes

import (
	"encoding/json"
	"io"
)

const (
	TYPE_FILE = "file"
)

type File struct {
	jsonObj *fileJson
}

func NewFile(title, body, fileName, fileType, fileUrl string) *File {
	jsonObj := &fileJson{
		Type:     TYPE_FILE,
		Title:    title,
		Body:     body,
		FileName: fileName,
		FileType: fileType,
		FileUrl:  fileUrl,
	}

	f := &File{
		jsonObj: jsonObj,
	}

	return f
}

func (f *File) Dump(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(f.jsonObj)
}

func (f *File) Type() string {
	return f.jsonObj.Type
}

func (f *File) Title() string {
	return f.jsonObj.Title
}

func (f *File) Body() string {
	return f.jsonObj.Body
}

func (f *File) FileName() string {
	return f.jsonObj.FileName
}

func (f *File) FileUrl() string {
	return f.jsonObj.FileUrl
}

func (f *File) FileType() string {
	return f.jsonObj.FileType
}

type fileJson struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	FileName string `json:"file_name"`
	FileUrl  string `json:"file_url"`
	FileType string `json:"file_type"`
}
