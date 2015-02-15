package pushes

const (
	TYPE_LINK = "link"
)

type Link struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`
}
