package shanhu

import (
	"testing"
)

func TestStates(t *testing.T) {
	s := newStates()
	state := s.New()
	t.Log("state: ", state)
	if !s.Check(state) {
		t.Errorf("check on state %q failed", state)
	}
}
