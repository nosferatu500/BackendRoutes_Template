package model

import (
	"github.com/dgrijalva/jwt-go"
	"math/big"
)

var (
	Secret          = "@&^4286dsk_834AVHHG"
	Gaslimit        *big.Int
	ContractAddress string
)

type ApiCustomClaims struct {
	AccessToken string `json:"access_token"`
	jwt.StandardClaims
}
