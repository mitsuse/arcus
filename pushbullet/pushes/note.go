package pushes

const (
	TYPE_NOTE = "note"
)

type Note struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
