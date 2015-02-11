package pushes

import (
	"encoding/json"
	"io"
)

const (
	TYPE_NOTE = "note"
)

type Note struct {
	jsonObj *noteJson
}

func NewNote(title, body string) *Note {
	jsonObj := &noteJson{
		Type:  TYPE_NOTE,
		Title: title,
		Body:  body,
	}

	n := &Note{
		jsonObj: jsonObj,
	}

	return n
}

func (n *Note) Dump(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(n.jsonObj)
}

func (n *Note) Type() string {
	return n.jsonObj.Type
}

func (n *Note) Title() string {
	return n.jsonObj.Title
}

func (n *Note) Body() string {
	return n.jsonObj.Body
}

type noteJson struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
