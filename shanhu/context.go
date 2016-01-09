package shanhu

import (
	"net/http"
	"time"
)

type context struct {
	w   http.ResponseWriter
	req *http.Request
}

func (c *context) redirect(url string) {
	http.Redirect(c.w, c.req, url, http.StatusFound)
}

func (c *context) readCookie(name string) string {
	cookie, err := c.req.Cookie(name)
	if err != nil || cookie == nil {
		return ""
	}
	return cookie.Value
}

func (c *context) writeCookie(name, value string, expires time.Time) {
	cookie := &http.Cookie{
		Name:    name,
		Value:   value,
		Path:    "/",
		Expires: expires,
	}
	http.SetCookie(c.w, cookie)
}

func (c *context) clearCookie(name string) {
	cookie := &http.Cookie{
		Name:  name,
		Value: "",
		Path:  "/",
	}
	http.SetCookie(c.w, cookie)
}
