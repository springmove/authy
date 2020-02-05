package oauth

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/linshenqi/sptty"
)

// 认证接口
func (s *Service) postAuth(ctx iris.Context) {
	req := Request{}
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.Write(sptty.NewRequestError(AuthFailed, err.Error()))
		return
	}

	resp, err := s.doAuth(req)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.Write(sptty.NewRequestError(AuthFailed, err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	body, _ := json.Marshal(resp)
	_, _ = ctx.Write(body)
}
