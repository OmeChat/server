package hashing

import (
	"crypto/sha512"
	"encoding/hex"
)

// SHA512 hashes the given string
// and encodes it into hex
func SHA512(val string) string {
	hasher := sha512.New()
	hasher.Write([]byte(val))
	return hex.EncodeToString(hasher.Sum(nil))
}
