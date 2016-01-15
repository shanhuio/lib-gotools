package shanhu

import (
	"encoding/json"
	"io/ioutil"
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
	log.Printf("updating repo %s (fake)", p)
}

func mapRepoName(fullName string) string {
	switch fullName {
	case "e8vm/e8vm":
		return "e8vm.io/e8vm"
	case "e8vm/tools":
		return "e8vm.io/tools"
	}
	return ""
}

func (h *hook) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("read error in hook", err)
		return
	}

	var dat struct {
		Repo struct {
			FullName string `json:"full_name"`
		} `json:"repository"`
	}

	err = json.Unmarshal(payload, &dat)
	if err != nil {
		log.Println(err)
		return
	}

	path := mapRepoName(dat.Repo.FullName)
	if path == "" {
		log.Printf("unrecognized repo: %s", dat.Repo.FullName)
		return
	}

	h.req <- path
}
