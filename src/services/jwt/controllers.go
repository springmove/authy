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
	_, _ = ctx.Write(sbody)
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

	claims, err := s.service.Validate(req.Claims, req.Token)
	if err != nil {
		ctx.StatusCode(iris.StatusConflict)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	sBody, _ := json.Marshal(claims)
	_, _ = ctx.Write(sBody)
}

func (s *JwtControllers) Parser(ctx iris.Context) {
	ctx.Header("content-type", "application/json")
	req := JwtAutherRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Write(services.NewRequestError(JWT_ERR_REQUEST_PAYLOAD, err.Error()))
		return
	}

	claims, err := s.service.Parse(req.Token)
	if err != nil {
		ctx.StatusCode(iris.StatusConflict)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	sBody, _ := json.Marshal(claims)
	_, _ = ctx.Write(sBody)
}
