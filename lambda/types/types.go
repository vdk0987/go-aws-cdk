package types

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username     string `json:"username" dynamodbav:"username"`
	PasswordHash string `json:"password" dynamodbav:"password"`
}

func NewUser(registerUser RegisterUser) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), 10)

	if err != nil {
		return User{}, err
	}

	return User{
		Username:     registerUser.Username,
		PasswordHash: string(hashedPassword),
	}, nil
}

func ValidatePassword(hashedPassword, plainPassword string) bool {
	res := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return res == nil
}

func CreateToken(user User) (string, error) {
	now := time.Now()
	validUntil := now.Add(time.Hour * 24)

	claims := jwt.MapClaims{
		"user":    user.Username,
		"expires": validUntil.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)
	jwtKey := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return tokenString,
			fmt.Errorf("signingString error %w", err)
	}
	return tokenString, nil
}
