package shanhu

import (
	"fmt"
	"log"
	"net/http"
)

// Handler is the http handler for shanhu.io website
type Handler struct {
	c        *Config
	gh       *gitHub
	sessions *sessionStore
}

// NewHandler creates a new http handler.
func NewHandler(c *Config) (*Handler, error) {
	gh := newGitHub(c.GitHubAppID, c.GitHubAppSecret, []byte(c.StateKey))
	return &Handler{
		c:        c,
		gh:       gh,
		sessions: newSessionStore([]byte(c.SessionKey)),
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
			session, expires := h.sessions.New(user)
			c.writeCookie("session", session, expires)
			c.redirect("/")
		}
	default:
		sessionStr := c.readCookie("session")
		ok, session := h.sessions.Check(sessionStr)

		if !ok {
			c.clearCookie("session")
			fmt.Fprintln(w, `<a href="/github/signin">please sign in</a>`)
			return
		}

		fmt.Fprintf(w, "You are %s.", session.User)
	}
}
