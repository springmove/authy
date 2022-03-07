package base

import "github.com/dgrijalva/jwt-go"

const (
	ServiceJwt = "jwt"
	TimeStamp  = "timestamp"

	CodeExpiry = "expiry"
	CodeFailed = "failed"
)

var Secret = "LijwefL(*IJWE)J@309j@#)(I#$@)(*"

type IServiceJwt interface {
	Sign(claims jwt.MapClaims) (string, error)
	Validate(myClaims jwt.MapClaims, tokenStr string) (jwt.MapClaims, error)
	Parse(tokenStr string) (jwt.MapClaims, error)
	Refresh(tokenStr string) (string, error)
	SetSecret(secret string)
}
