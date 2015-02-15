package pushes

const (
	TYPE_FILE = "file"
)

type File struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	FileName string `json:"file_name"`
	FileUrl  string `json:"file_url"`
	FileType string `json:"file_type"`
}
