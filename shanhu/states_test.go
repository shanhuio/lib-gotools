package shanhu

import (
	"testing"

	"time"
)

func TestStates(t *testing.T) {
	s := newStates(nil)
	state := s.New()
	if !s.Check(state) {
		t.Errorf("check on state %q failed", state)
	}
}

func TestStatesExpire(t *testing.T) {
	s := newStates(nil)
	now := time.Unix(0, 0)
	s.timeNow = func() time.Time { return now }
	state := s.New()
	t.Log("state: ", state)

	now = now.Add(stateTTL).Add(-time.Nanosecond)
	if !s.Check(state) {
		t.Errorf("check on state %q failed", state)
	}

	now = now.Add(time.Nanosecond)
	if s.Check(state) {
		t.Errorf("check passed, should fail because of time out")
	}
}
