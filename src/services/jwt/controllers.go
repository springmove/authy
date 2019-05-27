package jwt

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/linshenqi/authy/src/services"
)

type JwtControllers struct {
	service *JwtService
}

func (s *JwtControllers) Signer(ctx iris.Context) {
	ctx.Header("content-type", "application/json")

	claims := jwt.MapClaims{}
	err := ctx.ReadJSON(&claims)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Write(services.NewRequestError(JWT_ERR_REQUEST_PAYLOAD, err.Error()))
		return
	}

	token, err := s.service.Sign(claims)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Write(services.NewRequestError(JWT_ERR_SIGN_FAIL, err.Error()))
		return
	}

	sbody, _ := json.Marshal(JWTResponse{
		Token: token,
	})

	ctx.StatusCode(iris.StatusOK)
	ctx.Write(sbody)
}

func (s *JwtControllers) Auther(ctx iris.Context) {
	ctx.Header("content-type", "application/json")
	req := JwtAutherRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Write(services.NewRequestError(JWT_ERR_REQUEST_PAYLOAD, err.Error()))
		return
	}

	err = s.service.Validate(req.Claims, req.Token)
	if err != nil {
		ctx.StatusCode(iris.StatusConflict)
		return
	}

	ctx.StatusCode(iris.StatusOK)
}
