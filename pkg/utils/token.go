package utils

import (
	"encoding/base64"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/ssh"
)

func CreateToken(alg jwt.SigningMethod, did, domain string, notes interface{}, privateKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}

	key, err := ssh.ParseRawPrivateKey(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	token := jwt.New(alg)
	claims := make(jwt.MapClaims)
	claims["Did"] = did
	claims["Domain"] = domain
	claims["Notes"] = notes
	token.Claims = claims
	return token.SignedString(key)
}
