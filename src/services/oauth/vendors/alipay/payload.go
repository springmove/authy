package alipay

import (
	"github.com/springmove/authy/src/base"
	v3 "github.com/smartwalle/alipay/v3"
)

type UserInfoResponse struct {
	v3.UserInfoShareRsp
}

func (s *UserInfoResponse) toAuthResponseData() *base.Response {
	gender := 0
	if s.Content.Gender == "M" {
		gender = 1
	} else if s.Content.Gender == "F" {
		gender = 2
	}

	return &base.Response{
		Type:     base.AliPay,
		UnionID:  s.Content.UserId,
		Name:     s.Content.NickName,
		Gender:   gender,
		Province: s.Content.Province,
		City:     s.Content.City,
		Avatar:   s.Content.Avatar,
	}
}
