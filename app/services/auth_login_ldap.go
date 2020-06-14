package services

import (
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

var (
	LdapUserSearchNotFoundErr = errors.New("用户不存在或密码错误")
)

// AuthLoinLdapService ldap auth login
type AuthLoinLdapService struct {
	url    string
	config string
}

// NewAuthLoginLdapService
func NewAuthLoginLdapService() AuthLoginService {
	return &AuthLoinLdapService{}
}

// InitConf init ldap auth login config
func (al *AuthLoinLdapService) InitConf(url string, conf string) {
	al.url = url
	al.config = conf
}

// AuthLogin ldap auth
func (al *AuthLoinLdapService) AuthLogin(username string, password string) (*AuthLoginResponse, error) {

	if al.url == "" {
		return nil, fmt.Errorf("LDAP URL is empty")
	}
	if al.config == "" {
		return nil, fmt.Errorf("LDAP 配置数据错误")
	}

	lc, err := ldap.DialURL(al.url)
	if err != nil {
		return nil, fmt.Errorf("连接 LDAP 服务失败, err=%s", err.Error())
	}
	defer lc.Close()

	err = lc.Bind(username, password)
	if err != nil {
		return nil, fmt.Errorf("绑定 LDAP 用户失败, err=%s", err.Error())
	}

	searchRequest := ldap.NewSearchRequest(
		al.config,
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
