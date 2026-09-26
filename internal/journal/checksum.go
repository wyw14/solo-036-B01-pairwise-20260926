package journal

import (
	"crypto/sha256"
	"encoding/hex"
)

func Checksum(s *Store) string {
	h := sha256.New()
	for _, e := range s.Events() {
		h.Write([]byte(e.ID))
		h.Write([]byte(e.Payload))
	}
	return hex.EncodeToString(h.Sum(nil))
}
func Healthy(s *Store) bool { return s != nil && Checksum(s) != "" }
