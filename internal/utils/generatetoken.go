package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Jwtkey []byte

func Init(key string) {
	Jwtkey = []byte(key)
}
func GenerateJWT(id, role int, login string) (string, error) {

	claims := jwt.MapClaims{
		"id":    id,
		"login": login,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Jwtkey)
}
