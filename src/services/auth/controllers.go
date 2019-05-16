package auth

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/linshenqi/authy/src/core"
	"github.com/linshenqi/authy/src/services/wechat"
)

type AuthControllers struct {
	service *AuthService
}

// 认证接口
func (s *AuthControllers) Auth(ctx iris.Context) {
	auth := AuthBody{}
	err := ctx.ReadJSON(&auth)
	ctx.Header("content-type", "application/json")

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		b, _ := json.Marshal(core.RequestError{
			Code: AUTHY_ERR_REQUEST_BODY,
			Msg:  err.Error(),
		})
		ctx.Write(b)
		return
	}

	switch auth.Type {
	case AUTH_TYPE_WECHAT:
		// 微信认证
		wechatAuthData := auth.Data.(AuthWeChat)
		wechatAuth := s.service.core.GetService(AUTH_TYPE_WECHAT).(*wechat.WeCharService)
		wxresp, err := wechatAuth.Auth(wechatAuthData.Code)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			b, _ := json.Marshal(core.RequestError{
				Code: AUTHY_ERR_FAIL,
				Msg:  err.Error(),
			})
			ctx.Write(b)
			return
		}

		ctx.StatusCode(iris.StatusOK)
		sBody, _ := json.Marshal(AuthBody{
			Type: auth.Type,
			Data: wxresp,
		})

		ctx.Write(sBody)
		return

	default:
		ctx.StatusCode(iris.StatusBadRequest)
		b, _ := json.Marshal(core.RequestError{
			Code: AUTHY_ERR_TYPE,
			Msg:  "type error",
		})
		ctx.Write(b)
		return
	}

}
