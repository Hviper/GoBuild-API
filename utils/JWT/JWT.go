package JWT

import (
	"awesomeProject/dto"
	"github.com/dgrijalva/jwt-go"
	"time"
)
const (
	SecretKey = "welcome to hdc's gtihub"
)
func GetToken() dto.Token {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, _ := token.SignedString([]byte(SecretKey))
	val := dto.Token{tokenString}
	return val
}
func ValidateTokenMiddleware() {
	//TODO
}
