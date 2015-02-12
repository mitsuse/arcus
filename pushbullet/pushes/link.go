package pushes

import (
	"encoding/json"
	"io"
)

const (
	TYPE_LINK = "link"
)

type Link struct {
	jsonObj *linkJson
}

func NewLink(title, body, url string) *Link {
	jsonObj := &linkJson{
		Type:  TYPE_LINK,
		Title: title,
		Body:  body,
		Url:   url,
	}

	l := &Link{
		jsonObj: jsonObj,
	}

	return l
}

func (l *Link) Dump(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(l.jsonObj)
}

func (l *Link) Type() string {
	return l.jsonObj.Type
}

func (l *Link) Title() string {
	return l.jsonObj.Title
}

func (l *Link) Body() string {
	return l.jsonObj.Body
}

func (l *Link) Url() string {
	return l.jsonObj.Url
}

type linkJson struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`
}
