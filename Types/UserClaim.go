package types

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	jwt.RegisteredClaims
	Role  string
	Email string
}
