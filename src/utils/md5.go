package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(text string) string {
	hash := md5.Sum([]byte(text))
	md5Text := hex.EncodeToString(hash[:])
	return md5Text
}
