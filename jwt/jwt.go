package jwt

import (
	"github.com/nosferatu500/BackendRoutes_Template/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var signingKey = []byte(model.Secret)

func CreateToken(key string) (string, error) {
	claims := model.ApiCustomClaims{
		AccessToken: key,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)

	return tokenString, err
}

func ParseToken(tokenString string) (*model.ApiCustomClaims, error, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &model.ApiCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if claims, ok := token.Claims.(*model.ApiCustomClaims); ok && token.Valid {
		return claims, nil, true
	} else {
		return nil, err, false
	}
}
