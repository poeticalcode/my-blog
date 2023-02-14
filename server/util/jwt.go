package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

// Jwtkey 秘钥 可通过配置文件配置
var jwtSecret = []byte("blog_jwt_key")

// GenerateToken 生成一个 Token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()                    //当前时间
	expireTime := nowTime.Add(3 * time.Hour) //有效时间
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "its me",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 利用密钥进行签名，生成 Token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// GenerateToken 解析 Token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
