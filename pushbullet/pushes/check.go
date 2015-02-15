package pushes

const (
	TYPE_LIST = "list"
)

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
