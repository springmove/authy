package alipay

import (
	v3 "github.com/smartwalle/alipay/v3"
	"github.com/springmove/authy/src/base"
)

type UserInfoResponse struct {
	v3.UserInfoShareRsp
}

func (s *UserInfoResponse) toAuthResponseData() *base.Response {
	gender := 0
	if s.Gender == "M" {
		gender = 1
	} else if s.Gender == "F" {
		gender = 2
	}

	return &base.Response{
		Type:     base.AliPay,
		UnionID:  s.UserId,
		Name:     s.NickName,
		Gender:   gender,
		Province: s.Province,
		City:     s.City,
		Avatar:   s.Avatar,
	}
}
