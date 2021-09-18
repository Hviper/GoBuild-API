package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID           uint
	jwt.StandardClaims
}
func GenerateToken(id uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	issuer := "frank"
	claims := Claims{
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}