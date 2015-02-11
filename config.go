package bullet

import (
	"encoding/json"
	"io"
)

type Config struct {
	*configJson
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

func (c *Config) Token() string {
	return c.configJson.Token
}

type configJson struct {
	Token string `json:"token"`
}
