package bullet

import (
	"encoding/json"
	"io"
)

type Config struct {
	*configJson
}

func NewConfig() *Config {

	c := &Config{
		configJson: &configJson{},
	}

	return c
}

func LoadConfig(reader io.Reader) (*Config, error) {
	decoder := json.NewDecoder(reader)

	var jsonObj *configJson
	if err := decoder.Decode(&jsonObj); err != nil {
		return nil, err
	}

	config := &Config{
		configJson: jsonObj,
	}

	return config, nil
}

func DumpConfig(c *Config, writer io.Writer) error {
	return json.NewEncoder(writer).Encode(c.configJson)
}

func (c *Config) Token() string {
	return c.configJson.Token
}

func (c *Config) SetToken(token string) {
	c.configJson.Token = token
}

type configJson struct {
	Token string `json:"token"`
}
