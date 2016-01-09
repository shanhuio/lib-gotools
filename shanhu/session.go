package shanhu

import (
	"encoding/json"
	"time"
)

type sessionStore struct {
	s   *signer
	ttl time.Duration
}

type session struct {
	User    string
	Expires int64
}

const sessionTTL = time.Hour * 24 * 7

func newSessionStore(key []byte) *sessionStore {
	return &sessionStore{
		s:   newSigner(key),
		ttl: sessionTTL,
	}
}

func (s *sessionStore) New(user string) (string, time.Time) {
	expires := time.Now().Add(s.ttl)
	sess := &session{user, expires.UnixNano()}
	bs, err := json.Marshal(sess)
	if err != nil {
		panic(err)
	}
	return s.s.signHex(bs), expires
}

func (s *sessionStore) Check(str string) (bool, *session) {
	ok, bs := s.s.checkHex(str)
	if !ok {
		return false, nil
	}

	ret := new(session)
	err := json.Unmarshal(bs, ret)
	if err != nil {
		return false, nil
	}

	return true, ret
}
