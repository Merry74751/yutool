package encrypt

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func PassEncoding(s string) string {
	bytes := encode(s)
	return hex.EncodeToString(bytes)
}

func encode(s string) []byte {
	hash := hmac.New(sha512.New, []byte("key"))
	hash.Write([]byte(s))
	sum := hash.Sum(nil)
	return sum
}

func MatchPass(src, encoding string) bool {
	b1 := encode(src)
	b2, _ := hex.DecodeString(encoding)
	return hmac.Equal(b1, b2)
}
