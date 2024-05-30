package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/golang-jwt/jwt"
)

// CreateToken creates a JSON Web Token (JWT) with the provided user information, private key, and expiry duration.
// The user information is embedded in the token claims. The token is signed with the private key using the RS256 signing method.
// The expiry duration determines the validity period of the token.
// The function returns the generated token string and any error encountered during the process.
func CreateToken(userInfo models.Payload, privateKey *rsa.PrivateKey, expiryDuration time.Duration) (string, error) {
	// Embed user info in token claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"userInfo": userInfo,
		"exp":      time.Now().Add(expiryDuration).Unix(),
	})

	// Sign the token with the private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies the authenticity of a JWT token using the provided public key.
// It parses the token string and returns a *jwt.Token if the token is valid, otherwise it returns an error.
func VerifyToken(tokenString string, publicKey *rsa.PublicKey) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
}

// RandomKeyPair generates a random RSA key pair.
// It returns a pointer to the generated private key, a pointer to the corresponding public key, and any error encountered.
func RandomKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

// EncodePublicKeyToPem encodes the given RSA public key to PEM format.
// It takes a pointer to an rsa.PublicKey as input and returns the PEM-encoded public key as a string.
func EncodePublicKeyToPem(publicKey *rsa.PublicKey) string {
	publicKeyDer, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return ""
	}

	publicKeyBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDer,
	}

	publicKeyPem := pem.EncodeToMemory(&publicKeyBlock)
	return string(publicKeyPem)
}

// DecodePublicKeyFromPem decodes a PEM-encoded public key and returns an *rsa.PublicKey.
// It takes a string parameter publicKeyPem, which is the PEM-encoded public key.
// It returns the decoded *rsa.PublicKey and an error if decoding fails.
func DecodePublicKeyFromPem(publicKeyPem string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPem))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("unable to cast public key to RSA public key")
	}

	return rsaPub, nil
}
