package helper

import (
	"crypto/md5"
	"fmt"
)

// Md5 将string转换为MD5
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
