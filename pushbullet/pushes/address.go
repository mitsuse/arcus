package pushes

const (
	TYPE_ADDRESS = "address"
)

type Address struct {
	*Push
	Name    string `json:"name"`
	Address string `json:"address"`
}

func NewAddress() *Address {
	p := &Push{
		Type: TYPE_ADDRESS,
	}

	return &Address{Push: p}
}
