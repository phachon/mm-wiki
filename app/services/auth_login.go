package services

import (
	"fmt"
	"github.com/phachon/mm-wiki/app/models"
	"net/url"
)

// AuthLoginResponse auth login response result
type AuthLoginResponse struct {
	GivenName  string `json:"given_name"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	Phone      string `json:"phone"`
	Department string `json:"department"`
	Position   string `json:"position"`
	Location   string `json:"location"`
	Im         string `json:"im"`
}

const (
	AuthLoginProtocolHttp = "http"
	AuthLoginProtocolHttps = "https"
	AuthLoginProtocolLdap = "ldap"
	AuthLoginProtocolLdaps = "ldaps"
)

// AuthLoginService auth login service
type AuthLoginService interface {

	// Init init login config
	InitConf(url string, conf string)

	// GetServiceName get auth login service name
	GetServiceName() string

	// AuthLogin auth login request
	AuthLogin(username string, password string) (*AuthLoginResponse, error)
}

var AuthLogin = NewAuthLoginManager()

// AuthLoginManager auth login manager
type AuthLoginManager struct {
	AuthLoginServices map[string]AuthLoginService
}

// NewAuthLoginManager new a auth login manager
func NewAuthLoginManager() *AuthLoginManager {
	authLoginManager := &AuthLoginManager{
		AuthLoginServices: make(map[string]AuthLoginService),
	}
	return authLoginManager
}

// RegisterService register a auth login service
func (am *AuthLoginManager) RegisterService(serviceName string, authLoginService AuthLoginService) {
	if authLoginService == nil {
		return
	}
	if _, ok := am.AuthLoginServices[serviceName]; ok {
		panic(fmt.Sprintf("[AuthLoginManager] RegisterService '%s' already exists", serviceName))
	}
	am.AuthLoginServices[serviceName] = authLoginService
}

func (am *AuthLoginManager) UrlIsSupport(serviceName string) bool {
	if _, ok := am.AuthLoginServices[serviceName]; ok {
		return true
	}
	return false
}

// AuthLogin start auth login
func (am *AuthLoginManager) AuthLogin(username, password string) (*AuthLoginResponse, error) {

	// get auth login config
	authLogin, err := models.AuthModel.GetUsedAuth()
	if err != nil {
		return nil, fmt.Errorf("查找登录配置失败")
	}
	if len(authLogin) == 0 {
		return nil, fmt.Errorf("没有可用的统一登录配置")
	}
	// parse url protocol
	authUrl, ok := authLogin["url"]
	if !ok || authUrl == "" {
		return nil, fmt.Errorf("登录配置 url 无效")
	}
	u, err := url.Parse(authUrl)
	if err != nil {
		return nil, fmt.Errorf("登录配置 url 不合法：%s", err.Error())
	}
	serviceName := u.Scheme
	authLoginService, ok := am.AuthLoginServices[serviceName]
	if !ok {
		return nil, fmt.Errorf("登录配置 url 协议不支持")
	}
	serviceConf := ""
	if extData, ok := authLogin["ext_data"]; ok {
		serviceConf = extData
	}
	// init auth login service config
	authLoginService.InitConf(authUrl, serviceConf)
	// start auth login
	return authLoginService.AuthLogin(username, password)
}
