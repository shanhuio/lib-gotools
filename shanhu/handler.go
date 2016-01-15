package shanhu

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"e8vm.io/tools/repodb"
)

// Handler is the http handler for shanhu.io website
type Handler struct {
	c        *Config
	gh       *gitHub
	sessions *sessionStore
	users    map[string]bool
	db       *repodb.RepoDB
	hook     *hook

	fs http.Handler
}

// NewHandler creates a new http handler.
func NewHandler(c *Config) (*Handler, error) {
	dbPath := c.DB
	if dbPath == "" {
		dbPath = "shanhu.db"
	}

	db, err := repodb.Open(dbPath)
	if err != nil {
		return nil, err
	}

	gh := newGitHub(c.GitHubAppID, c.GitHubAppSecret, []byte(c.StateKey))

	users := make(map[string]bool)
	for _, u := range c.Users {
		users[u] = true
	}

	sessions := newSessionStore([]byte(c.SessionKey))

	return &Handler{
		c:        c,
		gh:       gh,
		sessions: sessions,
		users:    users,
		db:       db,
		hook:     newHook(db),
		fs:       http.FileServer(http.Dir("_")),
	}, nil
}

func (h *Handler) serveFile(c *context, p string) {
	c.req.URL.Path = p
	h.fs.ServeHTTP(c.w, c.req)
}

func (h *Handler) servePage(c *context, p string, dat interface{}) {
	t, err := template.ParseFiles(p)
	if err != nil {
		log.Println(err)
		http.Error(c.w, "page not found", 404)
		return
	}
	if err := t.Execute(c.w, dat); err != nil {
		log.Println(err)
		return
	}
}

func (h *Handler) hasUser(u string) bool { return h.users[u] }

func (h *Handler) serve(c *context, path string) {
	if strings.HasPrefix(path, "/assets/") {
		h.serveFile(c, path)
		return
	}
	if strings.HasPrefix(path, "/js/") {
		h.serveFile(c, path)
		return
	}

	log.Println(path)

	switch path {
	case "/style.css":
		h.serveFile(c, path)
	case "/favicon.ico":
		h.serveFile(c, "/assets/favicon.ico")
	case "/signin.html":
		h.serveFile(c, "/signin.html")

	case "/github/hook":
		h.hook.ServeHTTP(c.w, c.req)
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
			log.Println("session check failed")
			c.clearCookie("session")
			h.servePage(c, "_/signin.html", nil)
			return
		}

		user := session.User
		log.Printf("[%s] session check passed", user)

		if !h.hasUser(user) {
			c.clearCookie("session")
			msg := fmt.Sprintf("user %q not authorized, " +
				"please contact liulonnie@gmail.com.")
			log.Println(msg)
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

	path = strings.TrimPrefix(path, "/")
	_, dirs := pathSplit(path)

	if len(dirs) <= 1 {
		dat, err := projDat(h.db, user, path)
		if err != nil {
			log.Println(err)
			http.Error(c.w, "project loading failed", 404)
			return
		}
		h.servePage(c, "_/proj.html", dat)
	} else {
		dat, err := fileDat(h.db, user, path)
		if err != nil {
			log.Println(err)
			http.Error(c.w, "file loading failed", 404)
			return
		}
		h.servePage(c, "_/file.html", dat)
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

	switch path {
	case "/data/proj.js":

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
