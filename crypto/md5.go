package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5(password string) string {
	m := md5.New()
	m.Write([]byte(password))
	return hex.EncodeToString(m.Sum(nil))
}

func MD5Verify(password, hashPassword string) bool {
	m := md5.New()
	m.Write([]byte(hashPassword))
	return strings.EqualFold(password, hex.EncodeToString(m.Sum(nil)))
}
