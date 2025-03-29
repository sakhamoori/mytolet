package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sakhamoori/mytolet/api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type JWTProvider struct {
	secretKey string
}

func NewJWTProvider(secretKey string) *JWTProvider {
	return &JWTProvider{
		secretKey: secretKey,
	}
}

func (p *JWTProvider) GenerateToken(user *models.User) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["user_type"] = user.UserType
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token valid for 72 hours
	
	// Sign the token with our secret
	tokenString, err := token.SignedString([]byte(p.secretKey))
	if err != nil {
		return "", err
	}
	
	return tokenString, nil
}

func (p *JWTProvider) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(p.secretKey), nil
	})
	
	if err != nil {
		return nil, err
	}
	
	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	
	return nil, errors.New("invalid token")
}

func (p *JWTProvider) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p *JWTProvider) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}