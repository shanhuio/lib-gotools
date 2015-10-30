package main

import (
	"log"
	"net/http"
	"os/exec"
	"time"
)

type hook struct {
	p *project

	req chan bool
}

func newHook(p *project) *hook {
	ret := new(hook)
	ret.p = p
	ret.req = make(chan bool)

	go ret.service()

	return ret
}

func (h *hook) service() {
	ticker := time.Tick(time.Second)
	needUpdate := false

	for {
		select {
		case <-h.req:
			needUpdate = true
		case <-ticker:
			if needUpdate {
				needUpdate = false
				h.update()
			}
		}
	}
}

func (h *hook) update() {
	log.Print("(updating repo)")
	cmd := exec.Command("git", "pull")
	out, e := cmd.CombinedOutput()
	if e != nil {
		log.Print(string(out))
		log.Print(e)
		return
	}
	h.p.build()
}

func (h *hook) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.req <- true
}
