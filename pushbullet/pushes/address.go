package pushes

import (
	"encoding/json"
	"io"
)

const (
	TYPE_ADDRESS = "address"
)

type Address struct {
	jsonObj *addressJson
}

func NewAddress(name, address string) *Address {
	jsonObj := &addressJson{
		Type:    TYPE_ADDRESS,
		Name:    name,
		Address: address,
	}

	a := &Address{
		jsonObj: jsonObj,
	}

	return a
}

func (a *Address) Dump(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(a.jsonObj)
}

func (a *Address) Type() string {
	return a.jsonObj.Type
}

func (a *Address) Name() string {
	return a.jsonObj.Name
}

func (a *Address) Address() string {
	return a.jsonObj.Address
}

type addressJson struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
