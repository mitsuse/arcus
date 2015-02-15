/*
Package "pushes" provides an interface and structs to represent types of "push".

See the API documentation for the details: https://docs.pushbullet.com/#pushes
*/
package pushes

const (
	TYPE_NOTE    = "note"
	TYPE_LINK    = "link"
	TYPE_ADDRESS = "address"
	TYPE_LIST    = "list"
	TYPE_FILE    = "file"
)

type Push struct {
	Type string `json:"type"`
}

type Note struct {
	*Push
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewNote() *Note {
	p := &Push{
		Type: TYPE_NOTE,
	}

	return &Note{Push: p}
}

type Link struct {
	*Push
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`
}

func NewLink() *Link {
	p := &Push{
		Type: TYPE_LINK,
	}

	return &Link{Push: p}
}

type Address struct {
	*Push
	Name    string `json:"name"`
	Address string `json:"address"`
}

func NewAddress() *Address {
	p := &Push{
		Type: TYPE_ADDRESS,
	}

	return &Address{Push: p}
}

type Check struct {
	*Push
	Title   string   `json:"title"`
	ItemSeq []string `json:"items"`
}

func NewCheck() *Check {
	p := &Push{
		Type: TYPE_LIST,
	}

	return &Check{Push: p}
}

type File struct {
	*Push
	Title    string `json:"title"`
	Body     string `json:"body"`
	FileName string `json:"file_name"`
	FileUrl  string `json:"file_url"`
	FileType string `json:"file_type"`
}

func NewFile() *File {
	p := &Push{
		Type: TYPE_FILE,
	}

	return &File{Push: p}
}
