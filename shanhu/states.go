package shanhu

import (
	"encoding/binary"
	"time"
)

// states serves as a signed timestamp service for checking
// oauth2 states.
type states struct {
	s   *signer
	ttl time.Duration

	// timeNow grants ability to overwrite the time reading
	// from time package
	timeNow func() time.Time
}

func newStates(key []byte, ttl time.Duration) *states {
	return &states{
		s:   newSigner(key),
		ttl: ttl,
	}
}

func (s *states) now() time.Time {
	if s.timeNow == nil {
		return time.Now()
	}
	return s.timeNow()
}

const tsLen = 8

func (s *states) New() string {
	ts := make([]byte, tsLen)
	now := s.now().UnixNano()
	binary.LittleEndian.PutUint64(ts, uint64(now))
	return s.s.signHex(ts)
}

func (s *states) Check(state string) bool {
	ok, ts := s.s.checkHex(state)
	if !ok {
		return false
	}

	if len(ts) != tsLen {
		return false
	}

	t := int64(binary.LittleEndian.Uint64(ts))
	if t < 0 {
		return false
	}
	return s.now().Before(time.Unix(0, t).Add(s.ttl))
}
