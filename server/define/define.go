package define

import "github.com/golang-jwt/jwt"

type Userclaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"
var TokenExpire = 3600
var RefreshTokenExpire = 7200
