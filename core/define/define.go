package define

import (
	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

// DataSourceName MySQL datasource name
var DataSourceName = ""

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

// 腾讯云对象存储
var (
	TencentSecretKey = ""
	TencentSecretID  = ""
	CosBucket        = ""
)

// PageSize 分页的默认参数
var PageSize = 20

// Datetime 时间格式化
var Datetime = "2006-01-02 15:04:05"

var TokenExpire = 60 * 60 * 24 * 3        // 3 days
var RefreshTokenExpire = 60 * 60 * 24 * 7 // 7 days
