package define

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.RegisteredClaims
}

// JwtSignature jwt签名
var JwtSignature = "This is Light-Cloud's unique Signature-By Lin Yuhong"
