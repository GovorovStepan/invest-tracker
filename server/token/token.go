package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("gsptravelsecret") //TODO: move to env
type AccessTokenClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

func GenerateAccessToken(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour) //TODO: move to env
	claims := &AccessTokenClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
func GenerateRefreshToken(userID string) (tokenString string, err error) {
	refreshExpirationTime := time.Now().Add(30 * 24 * time.Hour) // TODO: move to env
	refreshClaims := &RefreshTokenClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateAccessToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&AccessTokenClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*AccessTokenClaim)
	if !ok {
		err = errors.New("claims parse problem")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func ValidateRefreshToken(signedToken string) (claims *RefreshTokenClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&RefreshTokenClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	refreshClaims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok {
		return nil, errors.New("claims parse problem")
	}

	if refreshClaims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}

	return refreshClaims, nil
}
