package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// import "github.com/dgrijalva/jwt-go"

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserRole string `json:"user_role"`
	UserId   int    `json:"id"`
	jwt.StandardClaims
}

var jwtKey = []byte("secret")

func GenerateJWT(email string, username string, userRole string, id int) (token string, err error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &JWTClaim{
		Username: username,
		Email:    email,
		UserRole: userRole,
		UserId:   id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = generateToken.SignedString(jwtKey)

	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
