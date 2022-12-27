package test

import (
	"chat_socket/server/define"
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt"
)

// 生成Token
func GenerateToken() (string, error) {

	uc := define.Userclaim{
		Id:       7,
		Identity: "b73eb6b6-6067-4604-9a01-e7cf13ef6dfd",
		Name:     "346614213",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Token解析
func TestAnalyzeToken(t *testing.T) {
	uc := new(define.Userclaim)
	token, err := GenerateToken()
	if err != nil {
		fmt.Println("token")
	}
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {

		fmt.Println("token is invalid")
	}
	if !claims.Valid {
		fmt.Println("token is invalid1")
	}
	fmt.Println("token:", token)
	fmt.Println("uc:", uc)
	fmt.Println("claims:", claims)
}
