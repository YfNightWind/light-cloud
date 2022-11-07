package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"light-cloud/src/core/define"
)

// Md5 将string转换为MD5(string->byte->md5)
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GenerateToken 生成token
func GenerateToken(id int, identity string, name string) (string, error) {
	// token使用用户的id、identity和用户名来加密
	userClaim := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)

	signedToken, err := token.SignedString([]byte(define.JwtSignature))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
