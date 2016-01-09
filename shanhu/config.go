package shanhu

import (
	"encoding/json"
	"io/ioutil"
)

// Config contains the configuration of the website
type Config struct {
	GitHubAppID     string
	GitHubAppSecret string
	StateKey        string
	Users           []string
	SessionKey      string
}

// LoadConfig loads the config from a json file.
func LoadConfig(file string) (*Config, error) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	c := new(Config)
	err = json.Unmarshal(bs, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
