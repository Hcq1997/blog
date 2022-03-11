package util

import (
	"crypto/md5"
	"encoding/hex"
)

func ToMd5(password string) string {
	ctx := md5.New()
	ctx.Write([]byte(password))
	return hex.EncodeToString(ctx.Sum(nil))
}
