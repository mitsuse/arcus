package pushes

import (
	"encoding/json"
	"io"
)

const (
	TYPE_LIST = "list"
)

type List struct {
	jsonObj *listJson
}

func NewList(title string, itemSeq ...string) *List {
	copySeq := make([]string, len(itemSeq))
	copy(copySeq, itemSeq)

	jsonObj := &listJson{
		Type:    TYPE_LIST,
		Title:   title,
		ItemSeq: copySeq,
	}

	l := &List{
		jsonObj: jsonObj,
	}

	return l
}

func (l *List) Dump(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(l.jsonObj)
}

func (l *List) Type() string {
	return l.jsonObj.Type
}

func (l *List) Title() string {
	return l.jsonObj.Title
}

func (l *List) Items() []string {
	copySeq := make([]string, len(l.jsonObj.ItemSeq))
	copy(copySeq, l.jsonObj.ItemSeq)

	return copySeq
}

type listJson struct {
	Type    string   `json:"type"`
	Title   string   `json:"title"`
	ItemSeq []string `json:"items"`
}
