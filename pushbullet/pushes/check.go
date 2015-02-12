package pushes

import (
	"encoding/json"
	"io"
)

const (
	TYPE_LIST = "list"
)

type Check struct {
	jsonObj *checkJson
}

func NewCheck(title string, itemSeq ...string) *Check {
	copySeq := make([]string, len(itemSeq))
	copy(copySeq, itemSeq)

	jsonObj := &checkJson{
		Type:    TYPE_LIST,
		Title:   title,
		ItemSeq: copySeq,
	}

	c := &Check{
		jsonObj: jsonObj,
	}

	return c
}

func (c *Check) Dump(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(c.jsonObj)
}

func (c *Check) Type() string {
	return c.jsonObj.Type
}

func (c *Check) Title() string {
	return c.jsonObj.Title
}

func (c *Check) Items() []string {
	copySeq := make([]string, len(c.jsonObj.ItemSeq))
	copy(copySeq, c.jsonObj.ItemSeq)

	return copySeq
}

type checkJson struct {
	Type    string   `json:"type"`
	Title   string   `json:"title"`
	ItemSeq []string `json:"items"`
}
