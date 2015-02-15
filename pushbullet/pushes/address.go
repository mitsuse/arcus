package pushes

const (
	TYPE_ADDRESS = "address"
)

type Address struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
