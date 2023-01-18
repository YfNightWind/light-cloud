package define

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.RegisteredClaims
}

// JwtSignature jwt签名
var JwtSignature = "你的JWT签名"

var (
	MailUsername = "你发送的邮箱"
	MailPassword = "你邮箱的授权码"
)
