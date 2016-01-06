package shanhu

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

// Handler is the http handler for shanhu.io website
type Handler struct {
	c  *Config
	gh *gitHub
}

// NewHandler creates a new http handler.
func NewHandler(c *Config) (*Handler, error) {
	var key []byte
	if c.StateAuthKey != "" {
		var err error
		key, err = base64.StdEncoding.DecodeString(c.StateAuthKey)
		if err != nil {
			return nil, fmt.Errorf("key error: %v", err)
		}
	}

	gh := newGitHub(c.GitHubAppID, c.GitHubAppSecret, key)
	return &Handler{c: c, gh: gh}, nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.Body.Close()
	fmt.Fprintf(w, "It works!")
}
