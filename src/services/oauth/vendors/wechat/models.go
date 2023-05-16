package wechat

import (
	"fmt"

	"github.com/springmove/authy/src/base"
	"gopkg.in/resty.v1"
)

const (
	WxOK = 0
)

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type MiniProgramAuthResponse struct {
	Response
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}

func (s *MiniProgramAuthResponse) toAuthResponseData() *base.Response {
	return &base.Response{
		Type:    base.WeChatMiniProgram,
		OpenID:  s.OpenID,
		UnionID: s.UnionID,
	}
}

type OAuthResponse struct {
	Response
	AccessToken  string `json:"access_token"`
	Expires      int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
	UnionID      string `json:"unionid"`
}

type UserInfo struct {
	OpenID     string   `json:"openid"`
	UnionID    string   `json:"unionid"`
	NickName   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privileges []string `json:"privilege"`
}

type UserInfoResponse struct {
	Response
	UserInfo
}

func (s *UserInfoResponse) toAuthResponseData() *base.Response {
	return &base.Response{
		Type:     base.WeChat,
		OpenID:   s.OpenID,
		UnionID:  s.UnionID,
		Name:     s.NickName,
		Gender:   s.Sex,
		Province: s.Province,
		City:     s.City,
		Country:  s.Country,
		Avatar:   s.HeadImgUrl,
	}
}

type ReqAccessToken struct {
	http   *resty.Client
	AppID  string
	Secret string
}

type RespAccessToken struct {
	Response
	Token  string `json:"access_token"`
	Expiry int    `json:"expires_in"`
}

type ReqMiniProgramAuthMobile struct {
	Code string `json:"code"`
}

type RespMiniProgramAuthMobile struct {
	Response

	Mobile *MiniProgramAuthMobile `json:"phone_info"`
}

type MiniProgramAuthMobile struct {
	FullMobile  string `json:"phoneNumber"`
	Mobile      string `json:"purePhoneNumber"`
	CountryCode string `json:"countryCode"`
}

func (s *MiniProgramAuthMobile) ToValidMobile() string {
	return fmt.Sprintf("+%s-%s", fmt.Sprintf("%03s", s.CountryCode), s.Mobile)
}
