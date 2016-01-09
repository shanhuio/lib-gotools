package shanhu

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func serveFile(w http.ResponseWriter, file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Println(err)
		return
	}

	defer f.Close()
	if _, err := io.Copy(w, f); err != nil {
		log.Println(err)
	}
}

func (h *Handler) serve(c *context, path string) {
	switch path {
	case "/github/signin":
		c.redirect(h.gh.signInURL())
	case "/github/callback":
		user, err := h.gh.callback(c.req)
		if err != nil {
			// print a better error message page
			fmt.Fprintln(c.w, err)
			return
		}
		if user != "h8liu" {
			fmt.Fprintln(c.w, err)
			return
		}
		if user == "h8liu" {
			session, expires := h.sessions.New(user)
			c.writeCookie("session", session, expires)
			c.redirect("/")
		}
	case "/favicon.ico":
		serveFile(c.w, "_/assets/favicon.ico")
	case "/style.css":
		c.w.Header().Set("Content-Type", "text/css")
		serveFile(c.w, "_/style.css")
	case "/assets/proj.png":
		serveFile(c.w, "_/assets/proj.png")
	case "/assets/pkg.png":
		serveFile(c.w, "_/assets/pkg.png")
	default:
		sessionStr := c.readCookie("session")
		ok, session := h.sessions.Check(sessionStr)

		if !ok {
			c.clearCookie("session")
			serveFile(c.w, "_/signin.html")
			return
		}

		h.serveUser(c, session.User, path)
	}
}

func (h *Handler) serveUser(c *context, user, path string) {
	serveFile(c.w, "_/proj.html")
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	c := &context{w: w, req: req}
	path := req.URL.Path
	log.Println("serving: ", path)
	h.serve(c, path)
}
