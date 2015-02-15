package pushes

const (
	TYPE_LINK = "link"
)

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
