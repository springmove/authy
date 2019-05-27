package jwt

import "github.com/dgrijalva/jwt-go"

const (
	JWT_ERR_REQUEST_PAYLOAD = "0001"
	JWT_ERR_SIGN_FAIL       = "0002"
)

type JWTResponse struct {
	Token string `json:"token"`
}

type JwtAutherRequest struct {
	Token  string        `json:"token"`
	Claims jwt.MapClaims `json:"claims"`
}
