package bullet

import (
	"encoding/json"
	"io"
)

type Auth struct {
	*authJson
}

func LoadAuth(reader io.Reader) (*Auth, error) {
	decoder := json.NewDecoder(reader)

	var jsonObj *authJson
	if err := decoder.Decode(&jsonObj); err != nil {
		return nil, err
	}

	auth := &Auth{
		authJson: jsonObj,
	}

	return auth, nil
}

func (a *Auth) Token() string {
	return a.authJson.Token
}

type authJson struct {
	Token string `json:"token"`
}
