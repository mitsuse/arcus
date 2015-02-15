package pushes

const (
	TYPE_FILE = "file"
)

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
