package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

// 使用 jwt-go 库

// 定义自己的JWT

// TokenExpireDuration 设置时间
const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("夏天的回忆")

// MyClaims 存储的数据不那么敏感就行了
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(),
			Issuer: "bluebell",
		},
	}
	// 使用指定的进行
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串 token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析 token
	var mc = new(MyClaims) // 手动初始内存，否则会出现一些问题
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
