package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

const (
	LdapDefaultAccountPattern     = "(&(objectClass=User)(userPrincipalName=%s))"
	LdapDefaultAttributeGivenName = "displayName"
)

var (
	LdapUserSearchNotFoundErr = errors.New("用户不存在或密码错误")
)

type AuthLoginConfig struct {
	BaseDn         string `json:"basedn"`
	BindUsername   string `json:"bind_username"`
	BindPassword   string `json:"bind_password"`
	AccountPattern string `json:"account_pattern"`
	GivenNameKey   string `json:"given_name_key"`
	EmailKey       string `json:"email_key"`
	MobileKey      string `json:"mobile_key"`
	PhoneKey       string `json:"phone_key"`
	DepartmentKey  string `json:"department_key"`
	PositionKey    string `json:"position_key"`
	LocationKey    string `json:"location_key"`
	ImKey          string `json:"im_key"`
}

// AuthLoginLdapService ldap auth login
type AuthLoginLdapService struct {
	url    string
	conf   string
	config *AuthLoginConfig
}

// NewAuthLoginLdapService
func NewAuthLoginLdapService() AuthLoginService {
	return &AuthLoginLdapService{}
}

// InitConf init ldap auth login config
func (al *AuthLoginLdapService) InitConf(url string, conf string) error {
	al.url = url
	al.conf = conf
	authLoginConfig := &AuthLoginConfig{}
	err := json.Unmarshal([]byte(conf), &authLoginConfig)
	if err != nil {
		return err
	}
	al.config = authLoginConfig
	if al.config.AccountPattern == "" {
		al.config.AccountPattern = LdapDefaultAccountPattern
	}
	if al.config.GivenNameKey == "" {
		al.config.GivenNameKey = LdapDefaultAttributeGivenName
	}
	return nil
}

// AuthLogin ldap auth
func (al *AuthLoginLdapService) AuthLogin(username string, password string) (*AuthLoginResponse, error) {

	if al.url == "" {
		return nil, fmt.Errorf("LDAP URL is empty")
	}
	if al.config == nil || al.conf == "" {
		return nil, fmt.Errorf("LDAP 配置数据错误")
	}
	if al.config.GivenNameKey == "" {
		return nil, fmt.Errorf("LDAP 配置 given_name_key 错误")
	}

	lc, err := ldap.DialURL(al.url)
	if err != nil {
		return nil, fmt.Errorf("连接 LDAP 服务失败, err=%s", err.Error())
	}
	defer lc.Close()

	// bind 用户
	if al.config.BindPassword != "" {
		err = lc.Bind(al.config.BindUsername, al.config.BindPassword)
	} else {
		err = lc.UnauthenticatedBind(al.config.BindUsername)
	}
	if err != nil {
		return nil, fmt.Errorf("绑定 LDAP 用户失败, err=%s", err.Error())
	}

	// 搜索下用户信息
	searchRequest := ldap.NewSearchRequest(
		al.config.BaseDn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf(al.config.AccountPattern, username),
		al.GetAttributes(),
		nil,
	)
	searchResult, err := lc.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("查找 LDAP 用户失败, err=%s", err.Error())
	}
	if len(searchResult.Entries) != 1 {
		return nil, LdapUserSearchNotFoundErr
	}

	// 验证下用户密码
	userDN := searchResult.Entries[0].DN
	err = lc.Bind(userDN, password)
	if err != nil {
		return nil, err
	}

	result := &AuthLoginResponse{
		GivenName:  searchResult.Entries[0].GetAttributeValue(al.config.GivenNameKey),
		Email:      searchResult.Entries[0].GetAttributeValue(al.config.EmailKey),
		Mobile:     searchResult.Entries[0].GetAttributeValue(al.config.MobileKey),
		Phone:      searchResult.Entries[0].GetAttributeValue(al.config.PhoneKey),
		Department: searchResult.Entries[0].GetAttributeValue(al.config.DepartmentKey),
		Position:   searchResult.Entries[0].GetAttributeValue(al.config.PositionKey),
		Location:   searchResult.Entries[0].GetAttributeValue(al.config.LocationKey),
		Im:         searchResult.Entries[0].GetAttributeValue(al.config.ImKey),
	}
	return result, nil
}

// GetAttributes get config attribute name
func (al *AuthLoginLdapService) GetAttributes() []string {

	attributes := []string{"dn"}
	confAttributes := []string{
		"dn", al.config.GivenNameKey, al.config.EmailKey,
		al.config.MobileKey, al.config.PhoneKey, al.config.DepartmentKey,
		al.config.PositionKey, al.config.LocationKey, al.config.ImKey,
	}
	for _, confAttribute := range confAttributes {
		if confAttribute == "" {
			continue
		}
		attributes = append(attributes, confAttribute)
	}
	return attributes
}

// GetServiceName ldap
func (al *AuthLoginLdapService) GetServiceName() string {
	return AuthLoginProtocolLdap
}

func init() {
	// ldap://xxx
	AuthLogin.RegisterService(AuthLoginProtocolLdap, NewAuthLoginLdapService())
	// ldaps://xxx
	AuthLogin.RegisterService(AuthLoginProtocolLdaps, NewAuthLoginLdapService())
}
