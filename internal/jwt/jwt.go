package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID   string
	jwt.RegisteredClaims
}

func CreateToken(id, key string) (string, error) {
	claim := Claims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Minute)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", fmt.Errorf("failed to create token: %w", err)
	}
	return tokenString, nil
}

func ValidateToken(tokenStr, key string) (*Claims, error) {
	claim := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}
	return claim, nil
}