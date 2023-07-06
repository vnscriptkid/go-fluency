package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// generate RSA keys
var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

func createToken(userId int64) (string, error) {
	// Set custom and standard claims
	claims := &CustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Sign the token with our private key
	tokenString, err := token.SignedString(privateKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func parseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func main() {
	// read the private key from a file
	privateKeyBytes, err := ioutil.ReadFile("./private_key.pem")
	if err != nil {
		log.Fatalf("Error reading private key: %v", err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}

	// read the public key from a file
	publicKeyBytes, err := ioutil.ReadFile("public_key.pem")
	if err != nil {
		log.Fatalf("Error reading public key: %v", err)
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		log.Fatalf("Error parsing public key: %v", err)
	}

	token, _ := createToken(123)
	fmt.Println("Token:", token)

	claims, _ := parseToken(token)
	fmt.Println("Claims:", claims)
}
