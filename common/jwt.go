package common

import (
	"Common/SQL"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//jwt加密密钥

var jwtKey = []byte("cbjcbsjcb")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// ReleaseToken 生成颁发Token
func ReleaseToken(user *SQL.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "lizhang",
			Subject:   "user token"},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
func JwtIsOk(tokenString string) bool {
	_, _, err := ParseToken(tokenString)
	if err != nil {
		return false
	}
	return true
}
