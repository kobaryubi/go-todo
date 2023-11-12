package jwt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": name,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	privateKey, err := GetTokenKey()
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenKey() error {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}

	x509Encoded, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}

	pemEncoded := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509Encoded,
	})

	file, err := os.Create("es256_private_key.pem")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(pemEncoded)
	if err != nil {
		return err
	}

	return nil
}

func GetTokenKey() (*ecdsa.PrivateKey, error) {
	bytes, err := os.ReadFile("es256_private_key.pem")
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, errors.New("")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
