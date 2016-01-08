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
	log.Println("get: ", path)

	switch path {
	case "/github/signin":
		http.Redirect(w, req, h.gh.signInURL(), http.StatusFound)
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
			addCookie := func(name, value string) {
				cookie := &http.Cookie{
					Name:    name,
					Value:   value,
					Path:    "/",
					Expires: expires,
				}
				http.SetCookie(w, cookie)
			}

			addCookie("user", user)
			addCookie("session", session)
			http.Redirect(w, req, "/", http.StatusFound)
		}
	default:
		readCookie := func(name string) string {
			cookie, err := req.Cookie(name)
			if err != nil || cookie == nil {
				return ""
			}
			return cookie.Value
		}

		user := readCookie("user")
		session := readCookie("session")
		if !h.sessions.Check(session) {
			fmt.Fprintln(w, `<a href="/github/signin">please sign in</a>`)
			return
		}

		fmt.Fprintf(w, "You are %s.", user)
	}
}
