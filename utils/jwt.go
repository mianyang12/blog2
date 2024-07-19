package utils

import (
	gojwt "github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// 将 JWT 密钥存储为字节切片
var jwtKey []byte

// init 函数在包加载时执行，此处用于初始化密钥
func init() {
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
}

// 定义一个结构体代表 JWT 声明，包含用户 ID 和标准声明
type Claims struct {
	Uid int
	gojwt.StandardClaims
}

// Award 函数用于为指定的用户 ID 生成 JWT
func Award(uid *int) (string, error) {
	// 设置过期时间为当前时间的 7 天后
	expireTime := time.Now().Add(7 * 24 * time.Hour)

	// 创建一个新的 JWT 声明实例
	claims := &Claims{
		Uid: *uid,
		StandardClaims: gojwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// 使用 HS256 签名算法创建 JWT
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)

	// 使用密钥对 JWT 进行签名，生成字符串形式的 token
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// ParseToken 函数用于解析给定的 JWT 字符串
func ParseToken(tokenStr string) (*gojwt.Token, *Claims, error) {
	// 创建一个空的 Claims 实例，用于存储解析后的声明信息
	claims := &Claims{}

	// 调用 jwt.ParseWithClaims 来解析 token，传入声明实例和一个回调函数，用于获取密钥
	token, err := gojwt.ParseWithClaims(tokenStr, claims, func(t *gojwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, nil, err
	}
	// 返回解析后的 token、声明实例和可能发生的任何错误
	return token, claims, err
}
