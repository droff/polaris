package models

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
)

// Session struct
type Session struct {
	Data map[string]bool
}

// Generate returns session key
func (s *Session) Generate() string {
	h := md5.New()
	key := fmt.Sprintf("%x", h.Sum(randToken()))
	s.Data[key] = true

	return key
}

// Find returns true if an entry is found
func (s *Session) Find(str string) bool {
	if _, ok := s.Data[str]; ok {
		return true
	}

	return false
}

// randToken returns a random set of bytes
func randToken() []byte {
	b := make([]byte, 8)
	rand.Read(b)

	return b
}
