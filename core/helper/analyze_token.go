package helper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"light-cloud/src/core/define"
)

// AnalyzeToken Token解析
func AnalyzeToken(token string) (*define.UserClaim, error) {
	userClaim := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtSignature), nil
	})
	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		return nil, errors.New("token不合法")
	}
	return userClaim, nil
}
