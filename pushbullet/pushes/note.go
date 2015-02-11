package pushes

const (
	TYPE_NOTE = "note"
)

type Note struct {
	title string
	body  string
}

func NewNote(title, body string) *Note {
	n := &Note{
		title: title,
		body:  body,
	}

	return n
}

func (n *Note) Type() string {
	return TYPE_NOTE
}

func (n *Note) Title() string {
	return n.title
}

func (n *Note) Body() string {
	return n.body
}
