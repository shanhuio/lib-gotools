package front

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	gh "golang.org/x/oauth2/github"
)

type github struct {
	config *oauth2.Config
	states *states
}

func newGithub(id, secret string) *github {
	return &github{
		config: &oauth2.Config{
			ClientID:     id,
			ClientSecret: secret,
			Scopes:       []string{}, // only need public information
			Endpoint:     gh.Endpoint,
		},
		states: new(states),
	}
}

var stateDuration = time.Minute

func (g *github) signInURL() string {
	state := g.states.New()
	return g.config.AuthCodeURL(state)
}

func (g *github) stateCode(req *http.Request) (state, code string) {
	values := req.URL.Query()
	state = values.Get("state")
	if state != "" {
		code = values.Get("code")
	}
	return state, code
}

func (g *github) getLogin(tok *oauth2.Token) (string, error) {
	client := g.config.Client(oauth2.NoContext, tok)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return "", fmt.Errorf("github api get: %v", err)
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("github api read body: %v", err)
	}
	var user struct {
		Login string `json:"login"`
	}
	if err := json.Unmarshal(bs, &user); err != nil {
		return "", err
	}
	ret := user.Login
	if ret == "" {
		return "", fmt.Errorf("empty login")
	}
	return ret, nil
}

func (g *github) callback(req *http.Request) (string, error) {
	state, code := g.stateCode(req)
	if state == "" {
		return "", fmt.Errorf("invalid oauth redirect")
	}

	check := g.states.Check(state)
	if !check {
		return "", fmt.Errorf("state invalid")
	}

	tok, err := g.config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return "", fmt.Errorf("exchange failed: %v", err)
	}
	if !tok.Valid() {
		return "", fmt.Errorf("token is invalid")
	}

	return g.getLogin(tok)
}
