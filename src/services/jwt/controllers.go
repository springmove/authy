package jwt

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/linshenqi/sptty"
)

func (s *Service) Signer(ctx iris.Context) {
	ctx.Header("content-type", "application/json")

	claims := jwt.MapClaims{}
	err := ctx.ReadJSON(&claims)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Write(sptty.NewRequestError(JWT_ERR_REQUEST_PAYLOAD, err.Error()))
		return
	}

	token, err := s.Sign(claims)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Write(sptty.NewRequestError(JWT_ERR_SIGN_FAIL, err.Error()))
		return
	}

	sbody, _ := json.Marshal(JWTResponse{
		Token: token,
	})

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.Write(sbody)
}

func (s *Service) Auther(ctx iris.Context) {
	ctx.Header("content-type", "application/json")
	req := JwtAutherRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Write(sptty.NewRequestError(JWT_ERR_REQUEST_PAYLOAD, err.Error()))
		return
	}

	claims, err := s.Validate(req.Claims, req.Token)
	if err != nil {
		ctx.StatusCode(iris.StatusConflict)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	sBody, _ := json.Marshal(claims)
	_, _ = ctx.Write(sBody)
}

func (s *Service) Parser(ctx iris.Context) {
	ctx.Header("content-type", "application/json")
	req := JwtAutherRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Write(sptty.NewRequestError(JWT_ERR_REQUEST_PAYLOAD, err.Error()))
		return
	}

	claims, err := s.Parse(req.Token)
	if err != nil {
		ctx.StatusCode(iris.StatusConflict)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	sBody, _ := json.Marshal(claims)
	_, _ = ctx.Write(sBody)
}
