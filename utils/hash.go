package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func HashMD5(value string) string {
	sum := md5.Sum([]byte(value))
	return hex.EncodeToString(sum[:])
}
