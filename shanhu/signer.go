package shanhu

import (
	"bytes"

	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type signer struct {
	key []byte
}

func randBytes(n int) []byte {
	ret := make([]byte, n)
	_, err := rand.Read(ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func newSigner(key []byte) *signer {
	if key == nil {
		key = randBytes(32)
	}
	return &signer{key: key}
}

func (s *signer) hash(dat []byte) []byte {
	m := hmac.New(sha256.New, s.key)
	m.Write(dat)
	return m.Sum(nil)
}

func (s *signer) sign(dat []byte) []byte {
	buf := new(bytes.Buffer)
	buf.Write(dat)

	h := s.hash(buf.Bytes())
	buf.Write(h)

	return buf.Bytes()
}

func (s *signer) signHex(dat []byte) string {
	return hex.EncodeToString(s.sign(dat))
}

func (s *signer) check(bs []byte) (bool, []byte) {
	n := len(bs)
	if n < sha256.Size {
		return false, nil
	}

	if n < sha256.Size {
		return false, nil
	}

	dat := bs[:n-sha256.Size]
	hashGot := bs[n-sha256.Size:]
	hashWant := s.hash(dat)
	if !hmac.Equal(hashGot, hashWant) {
		return false, nil
	}
	return true, dat
}

func (s *signer) checkHex(str string) (bool, []byte) {
	bs, err := hex.DecodeString(str)
	if err != nil {
		return false, nil
	}
	return s.check(bs)
}
