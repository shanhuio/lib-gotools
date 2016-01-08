package shanhu

import (
	"testing"

	"time"
)

func TestStates(t *testing.T) {
	s := newStates(nil, time.Second)
	state := s.New()
	if !s.Check(state) {
		t.Errorf("check on state %q failed", state)
	}

	if s.Check("") {
		t.Errorf("check on empty state is passing")
	}
}

func TestStatesExpire(t *testing.T) {
	const ttl = time.Second
	s := newStates(nil, ttl)
	now := time.Unix(0, 0)
	s.timeNow = func() time.Time { return now }
	state := s.New()
	t.Log("state: ", state)

	now = now.Add(ttl).Add(-time.Nanosecond)
	if !s.Check(state) {
		t.Errorf("check on state %q failed", state)
	}

	now = now.Add(time.Nanosecond)
	if s.Check(state) {
		t.Errorf("check passed, should fail because of time out")
	}
}
