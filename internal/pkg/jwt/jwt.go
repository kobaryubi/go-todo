package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: 秘密鍵を置き換える
const SecretKey = "secret"

func CreateToken(name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": name,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
