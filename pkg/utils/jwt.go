package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtKey = []byte("AllYourBase")

type MyCustomClaims struct {
	UserID uint `json:"user_id"`
	//Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(userID uint) (string, error) {
	expiresAt := time.Now().Add(2 * time.Hour) // 2小时过期一次
	claims := MyCustomClaims{
		userID,
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(time.Unix(expiresAt.Unix(), 0)),
			Issuer:    "lemuzhi",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(jwtKey)
	return stringToken, err
}

// ParseToken 验证token
func ParseToken(stringToken string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(stringToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if token != nil {
		claims, ok := token.Claims.(*MyCustomClaims)
		if ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
