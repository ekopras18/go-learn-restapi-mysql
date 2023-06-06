package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("learn-restapi-mysql")
var JWT_EXPIRE = 30 // in minutes

type JWTClaim struct {
	Email    string
	Password string
	jwt.RegisteredClaims
}
