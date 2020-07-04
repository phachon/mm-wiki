package services

import (
	"encoding/json"
	"fmt"
	"github.com/phachon/mm-wiki/app/utils"
)

type AuthLoginHttpRes struct {
	Message string             `json:"message"`
	Data    *AuthLoginResponse `json:"data"`
}

// AuthLoinHttpService http auth login
type AuthLoinHttpService struct {
	url     string
	extData string
}

// NewAuthLoginHttpService
func NewAuthLoginHttpService() AuthLoginService {
	return &AuthLoinHttpService{}
}

// InitConf init http auth login config
func (ah *AuthLoinHttpService) InitConf(url string, conf string) error {
	ah.url = url
	ah.extData = conf
	return nil
}

// AuthLogin send http request
func (ah *AuthLoinHttpService) AuthLogin(username string, password string) (*AuthLoginResponse, error) {
	if ah.url == "" {
		return nil, fmt.Errorf("authLogin url is empty")
	}
	queryValue := map[string]string{
		"username": username,
		"password": password,
		"ext_data": ah.extData,
	}
	// request auth login api
	body, code, err := utils.Request.HttpPost(ah.url, queryValue, nil)
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, fmt.Errorf("登录认证失败, httpCode=%d", code)
	}
	v := &AuthLoginHttpRes{}
	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("登录认证失败，返回结果不合法 err=%s", err.Error())
	}
	if v.Message != "" {
		return nil, fmt.Errorf("登录认证失败, message=%s", v.Message)
	}
	return v.Data, nil
}

// GetServiceName http
func (ah *AuthLoinHttpService) GetServiceName() string {
	return AuthLoginProtocolHttp
}

func init() {
	// http://xxx
	AuthLogin.RegisterService(AuthLoginProtocolHttp, NewAuthLoginHttpService())
	// https://xxx
	AuthLogin.RegisterService(AuthLoginProtocolHttps, NewAuthLoginHttpService())
}
