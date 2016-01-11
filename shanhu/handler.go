package shanhu

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Handler is the http handler for shanhu.io website
type Handler struct {
	c        *Config
	gh       *gitHub
	sessions *sessionStore
	users    map[string]bool

	fs http.Handler
}

// NewHandler creates a new http handler.
func NewHandler(c *Config) (*Handler, error) {
	gh := newGitHub(c.GitHubAppID, c.GitHubAppSecret, []byte(c.StateKey))

	users := make(map[string]bool)
	for _, u := range c.Users {
		users[u] = true
	}

	return &Handler{
		c:        c,
		gh:       gh,
		sessions: newSessionStore([]byte(c.SessionKey)),
		users:    users,
		fs:       http.FileServer(http.Dir("_")),
	}, nil
}

func (h *Handler) serveFile(c *context, p string) {
	c.req.URL.Path = p
	h.fs.ServeHTTP(c.w, c.req)
}

func (h *Handler) hasUser(u string) bool { return h.users[u] }

func (h *Handler) serve(c *context, path string) {
	if strings.HasPrefix(path, "/assets/") {
		h.serveFile(c, path)
		return
	}

	switch path {
	case "/style.css":
		h.serveFile(c, path)
	case "/favicon.ico":
		h.serveFile(c, "/assets/favicon.ico")

	case "/github/signin":
		c.redirect(h.gh.signInURL())
	case "/github/callback":
		user, err := h.gh.callback(c.req)
		if err != nil {
			// print a better error message page
			fmt.Fprintln(c.w, err)
			return
		}
		if !h.hasUser(user) {
			fmt.Fprintf(c.w, "User %q is not authorized.", user)
			log.Printf("Unauthorized user %q tried to sign in.", user)
			return
		}

		log.Printf("User %q signed in.", user)
		session, expires := h.sessions.New(user)
		c.writeCookie("session", session, expires)
		c.redirect("/")
	case "/github/signout":
		c.clearCookie("session")
		c.redirect("/")

	default:
		sessionStr := c.readCookie("session")
		ok, session := h.sessions.Check(sessionStr)
		if !ok {
			c.clearCookie("session")
			h.serveFile(c, "/signin.html")
			return
		}

		user := session.User
		if !h.hasUser(user) {
			c.clearCookie("session")
			msg := fmt.Sprintf("user %q not authorized, " +
				"please contact liulonnie@gmail.com.")
			http.Error(c.w, msg, 403)
			return
		}

		h.serveUser(c, user, path)
	}
}

func (h *Handler) serveUser(c *context, user, path string) {
	log.Printf("[%s] %s", user, path)
	if strings.HasPrefix(path, "/data/") {
		h.serveData(c, user, path)
		return
	}

	switch path {
	case "/proj.html", "/":
		h.serveFile(c, "/proj.html")
	case "/file.html":
		h.serveFile(c, path)

	case "/js/proj.js":
		h.serveFile(c, "/proj.js")
	case "/js/file.js":
		h.serveFile(c, "/file.js")
	case "/shanhu.js":
		h.serveFile(c, "/shanhu.js")

	default:
		http.Error(c.w, "File not found.", 404)
	}
}

func serveJsVar(w http.ResponseWriter, v interface{}) {
	fmt.Fprintf(w, "var shanhuData = ")
	bs, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(bs)
	fmt.Fprintf(w, ";")
}

func (h *Handler) serveData(c *context, user, path string) {
	obj := make(map[string]interface{})
	obj["session"] = map[string]string{
		"user": user,
	}

	switch path {
	case "/data/proj.js":
		obj["proj"] = nil
	default:
		http.Error(c.w, "Invalid data request.", 404)
		return
	}

	serveJsVar(c.w, obj)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	c := &context{w: w, req: req}
	path := req.URL.Path
	h.serve(c, path)
}
