package pushes

const (
	TYPE_NOTE = "note"
)

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
