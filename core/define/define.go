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

// 邮箱的账号和授权码
var (
	MailUsername = "你发送的邮箱"
	MailPassword = "你邮箱的授权码"
)

// 验证码相关
var (
	CodeLength = 6
	ExpireTime = 300 // 秒
)

// redis
var (
	RedisAddress  = "你的redis地址"
	RedisPassword = "你的redis密码，没有就空着"
)
