package shanhu

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Handler is the http handler for shanhu.io website
type Handler struct {
	c        *Config
	gh       *gitHub
	sessions *states
}

const sessionTTL = time.Hour * 24 * 7

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
	return &Handler{
		c:        c,
		gh:       gh,
		sessions: newStates(nil, sessionTTL),
	}, nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	path := req.URL.Path
	log.Println("serving: ", path)
	c := &context{w: w, req: req}

	switch path {
	case "/github/signin":
		c.redirect(h.gh.signInURL())
	case "/github/callback":
		user, err := h.gh.callback(req)
		if err != nil {
			// print a better error message page
			fmt.Fprintln(w, err)
			return
		}
		if user != "h8liu" {
			fmt.Fprintln(w, err)
			return
		}
		if user == "h8liu" {
			session := h.sessions.New()
			expires := time.Now().Add(sessionTTL)
			c.writeCookie("user", user, expires)
			c.writeCookie("session", session, expires)
			c.redirect("/")
		}
	default:
		// TODO: this is wrong
		// this trusts the user field when the session passes,
		// this will allow the user pretent to be anyone by using
		// a valid session token
		user := c.readCookie("user")
		session := c.readCookie("session")
		if !h.sessions.Check(session) {
			fmt.Fprintln(w, `<a href="/github/signin">please sign in</a>`)
			return
		}

		fmt.Fprintf(w, "You are %s.", user)
	}
}
