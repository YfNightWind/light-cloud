package helper

import (
	"github.com/dgrijalva/jwt-go"
	"light-cloud/src/core/define"
	"time"
)

// GenerateToken 生成token
func GenerateToken(id int, identity string, name string, second int) (string, error) {
	// token使用用户的id、identity和用户名来加密
	userClaim := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)

	signedToken, err := token.SignedString([]byte(define.JwtSignature))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
