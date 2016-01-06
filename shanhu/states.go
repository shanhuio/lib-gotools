package shanhu

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"time"
)

type states struct {
	key []byte
}

func newStates() *states {
	const keyLen = 256
	key := make([]byte, keyLen)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return &states{key}
}

const saltLen = 8

func (s *states) hash(salt []byte) []byte {
	m := hmac.New(sha256.New, s.key)
	m.Write(salt)
	return m.Sum(nil)
}

func (s *states) New() string {
	buf := make([]byte, saltLen+sha256.Size)
	salt := buf[:saltLen]
	now := time.Now().UnixNano()
	binary.LittleEndian.PutUint64(buf[:saltLen], uint64(now))
	h := s.hash(salt)
	copy(buf[saltLen:], h)
	return base64.URLEncoding.EncodeToString(buf)
}

var stateTTL = time.Minute * 3

func (s *states) Check(state string) bool {
	bs, err := base64.URLEncoding.DecodeString(state)
	if err != nil {
		return false
	}
	salt := bs[:saltLen]
	got := bs[saltLen:]
	want := s.hash(salt)
	if !hmac.Equal(got, want) {
		return false
	}
	t := int64(binary.LittleEndian.Uint64(salt))
	if t < 0 {
		return false
	}
	return time.Now().Before(time.Unix(0, t).Add(stateTTL))
}
