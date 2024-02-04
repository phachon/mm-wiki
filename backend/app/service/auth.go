package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/config"
	"github.com/phachon/mm-wiki/gopkg/errors"
)

// Auth 认证登录业务逻辑
type Auth struct {
	ctx context.Context
}

// Claims 声明
type Claims struct {
	AccountName string `json:"account_name"`
	AccountID   int64  `json:"account_id"`
	jwt.RegisteredClaims
}

// NewAuth 创建登录业务逻辑对象
func NewAuth(ctx context.Context) *Auth {
	return &Auth{
		ctx: ctx,
	}
}

// Login 账号登录逻辑
func (a *Auth) Login(accountName string, password string) (loginToken string, accountInfo *entity.AccountEntity, bizErr errors.BizError) {
	// 根据账号名查询账号信息
	accountInfo, bizErr = dao.NewAccount(a.ctx).GetAccountByName(accountName)
	if bizErr != nil {
		return loginToken, accountInfo, bizErr
	}
	if accountInfo == nil {
		return loginToken, accountInfo, errors.Errorf(errors.BusinessRecordNotExistError, "账号不存在")
	}
	if accountInfo.Status == entity.AccountStatusForbid {
		return loginToken, accountInfo, errors.Errorf(errors.BusinessForbiddenError, "账号被禁用")
	}
	// 判断密码是否相等 md5 加密
	if PasswordEncode(password) != accountInfo.Password {
		return loginToken, accountInfo, errors.Errorf(errors.BusinessPasswordError, "密码错误")
	}
	// 登录成功，生成 token
	var err error
	loginToken, err = a.GenerateLoginToken(accountInfo.AccountId, accountInfo.Name)
	if err != nil {
		return loginToken, accountInfo, errors.Errorf(errors.BusinessAuthTokenMakeErr, "token错误")
	}
	return loginToken, accountInfo, nil
}

// GenerateLoginToken 生成登录 token
func (a *Auth) GenerateLoginToken(accountId int64, accountName string) (string, error) {
	nowTime := time.Now()
	expireTime := time.Duration(config.GetAppConf().Auth.ExpireHours) * time.Hour
	claims := Claims{
		accountName,
		accountId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(nowTime.Add(expireTime)),
			Issuer:    accountName,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := config.GetAppConf().GetAuthConf().JwtSecret
	// 获取完整签名之后的 token
	return tokenClaims.SignedString([]byte(jwtSecret))
}

// ParseToken 解析 token
func ParseToken(token string) (*Claims, errors.BizError) {
	tokenClaims, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := config.GetAppConf().GetAuthConf().JwtSecret
		return []byte(jwtSecret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, errors.Errorf(errors.BusinessAuthTokenParseErr, "parse token fail")
}
