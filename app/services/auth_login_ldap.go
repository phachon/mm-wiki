package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

var (
	LdapUserSearchNotFoundErr = errors.New("用户不存在或密码错误")
)

type AuthLoginConfig struct {
	BaseDn       string `json:"basedn"`
	BindUsername string `json:"bind_username"`
	BindPassword string `json:"bind_password"`
}

// AuthLoinLdapService ldap auth login
type AuthLoinLdapService struct {
	url    string
	conf   string
	config *AuthLoginConfig
}

// NewAuthLoginLdapService
func NewAuthLoginLdapService() AuthLoginService {
	return &AuthLoinLdapService{}
}

// InitConf init ldap auth login config
func (al *AuthLoinLdapService) InitConf(url string, conf string) error {
	al.url = url
	al.conf = conf
	authLoginConfig := &AuthLoginConfig{}
	err := json.Unmarshal([]byte(conf), &authLoginConfig)
	if err != nil {
		return err
	}
	al.config = authLoginConfig
	return nil
}

// AuthLogin ldap auth
func (al *AuthLoinLdapService) AuthLogin(username string, password string) (*AuthLoginResponse, error) {

	if al.url == "" {
		return nil, fmt.Errorf("LDAP URL is empty")
	}
	if al.config == nil || al.conf == "" {
		return nil, fmt.Errorf("LDAP 配置数据错误")
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
		fmt.Sprintf("(&(objectClass=User)(userPrincipalName=%s))", username),
		[]string{"dn", "mail", "displayName", "telephoneNumber", "mobile", "department", "physicalDeliveryOfficeName"},
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
		GivenName:  searchResult.Entries[0].GetAttributeValue("displayName"),
		Email:      searchResult.Entries[0].GetAttributeValue("mail"),
		Mobile:     searchResult.Entries[0].GetAttributeValue("mobile"),
		Phone:      searchResult.Entries[0].GetAttributeValue("telephoneNumber"),
		Department: searchResult.Entries[0].GetAttributeValue("department"),
		Position:   searchResult.Entries[0].GetAttributeValue(""),
		Location:   searchResult.Entries[0].GetAttributeValue("physicalDeliveryOfficeName"),
		Im:         "",
	}
	return result, nil
}

// GetServiceName ldap
func (al *AuthLoinLdapService) GetServiceName() string {
	return AuthLoginProtocolLdap
}

func init() {
	// ldap://xxx
	AuthLogin.RegisterService(AuthLoginProtocolLdap, NewAuthLoginLdapService())
	// ldaps://xxx
	AuthLogin.RegisterService(AuthLoginProtocolLdaps, NewAuthLoginLdapService())
}
