package helper

import (
	"chat_socket/server/define"
	"crypto/md5"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// 生成Token
func GenerateToken(id int, identity, name string, second int) (string, error) {

	uc := define.Userclaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Token解析
func AnalyzeToken(token string) (*define.Userclaim, error) {
	uc := new(define.Userclaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}

	return uc, err
}
