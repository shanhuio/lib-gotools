package shanhu

import (
	"log"
	"net/http"
	"time"
)

type hook struct {
	req chan string
}

func newHook() *hook {
	ret := &hook{
		req: make(chan string),
	}

	go ret.serve()
	return ret
}

func (h *hook) serve() {
	ticker := time.Tick(time.Second)
	pending := make(map[string]bool)

	for {
		select {
		case p := <-h.req:
			pending[p] = true
		case <-ticker:
			if len(pending) > 0 {
				for p := range pending {
					h.update(p)
				}
			}
			pending = make(map[string]bool)
		}
	}
}

func (h *hook) update(p string) {
	log.Printf("updating repo %s", p)
}

func (h *hook) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.req <- "x"
}
