package pushes

const (
	TYPE_LIST = "list"
)

type Check struct {
	Type    string   `json:"type"`
	Title   string   `json:"title"`
	ItemSeq []string `json:"items"`
}
