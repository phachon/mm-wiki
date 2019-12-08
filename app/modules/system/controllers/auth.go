package controllers

import (
	"strings"

	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/utils"

	"github.com/astaxie/beego/validation"
	valid "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)
	limit := (page - 1) * number

	var err error
	var count int64
	var auths []map[string]string
	if keyword != "" {
		count, err = models.AuthModel.CountAuthsByKeyword(keyword)
		auths, err = models.AuthModel.GetAuthsByKeywordAndLimit(keyword, limit, number)
	} else {
		count, err = models.AuthModel.CountAuths()
		auths, err = models.AuthModel.GetAuthsByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("获取登录认证列表失败: " + err.Error())
		this.ViewError("获取登录认证列表失败", "/system/main/index")
	}

	this.Data["auths"] = auths
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("auth/list", "auth")
}

func (this *AuthController) Add() {
	this.viewLayout("auth/form", "auth")
}

func (this *AuthController) Save() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/auth/list")
	}
	name := strings.TrimSpace(this.GetString("name", ""))
	url := strings.TrimSpace(this.GetString("url", ""))
	usernamePrefix := strings.TrimSpace(this.GetString("username_prefix", ""))
	extData := strings.TrimSpace(this.GetString("ext_data", ""))

	v := validation.Validation{}
	if name == "" {
		this.jsonError("登录认证名称不能为空！")
	}
	if usernamePrefix == "" {
		this.jsonError("用户名前缀不能为空！")
	}
	if !v.AlphaNumeric(usernamePrefix, "username_prefix").Ok {
		this.jsonError("用户名前缀格式不正确！")
	}
	if url == "" {
		this.jsonError("认证URL不能为空！")
	}
	if valid.Validate(url, is.URL) != nil {
		this.jsonError("认证URL格式不正确！")
	}

	ok, err := models.AuthModel.HasAuthName(name)
	if err != nil {
		this.ErrorLog("添加登录认证失败：" + err.Error())
		this.jsonError("添加登录认证失败！")
	}
	if ok {
		this.jsonError("登录认证名称已经存在！")
	}

	ok, err = models.AuthModel.HasAuthUsernamePrefix(usernamePrefix)
	if err != nil {
		this.ErrorLog("添加登录认证失败：" + err.Error())
		this.jsonError("添加登录认证失败！")
	}
	if ok {
		this.jsonError("用户名前缀已经存在！")
	}

	authId, err := models.AuthModel.Insert(map[string]interface{}{
		"name":            name,
		"url":             url,
		"username_prefix": usernamePrefix,
		"ext_data":        extData,
	})

	if err != nil {
		this.ErrorLog("添加登录认证失败：" + err.Error())
		this.jsonError("添加登录认证失败")
	}
	this.InfoLog("添加登录认证 " + utils.Convert.IntToString(authId, 10) + " 成功")
	this.jsonSuccess("添加登录认证成功", nil, "/system/auth/list")
}

func (this *AuthController) Edit() {

	authId := this.GetString("login_auth_id", "")
	if authId == "" {
		this.ViewError("登录认证不存在", "/system/auth/list")
	}

	auth, err := models.AuthModel.GetAuthByAuthId(authId)
	if err != nil {
		this.ViewError("登录认证不存在", "/system/auth/list")
	}

	this.Data["auth"] = auth
	this.viewLayout("auth/form", "auth")
}

func (this *AuthController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/auth/list")
	}
	authId := this.GetString("login_auth_id", "")
	name := strings.TrimSpace(this.GetString("name", ""))
	url := strings.TrimSpace(this.GetString("url", ""))
	usernamePrefix := strings.TrimSpace(this.GetString("username_prefix", ""))
	extData := strings.TrimSpace(this.GetString("ext_data", ""))

	v := validation.Validation{}
	if authId == "" {
		this.jsonError("登录认证不存在！")
	}
	if name == "" {
		this.jsonError("登录认证名称不能为空！")
	}
	if usernamePrefix == "" {
		this.jsonError("用户名前缀不能为空！")
	}
	if !v.AlphaNumeric(usernamePrefix, "username_prefix").Ok {
		this.jsonError("用户名前缀格式不正确！")
	}
	if url == "" {
		this.jsonError("认证URL不能为空！")
	}
	if valid.Validate(url, is.URL) != nil {
		this.jsonError("认证URL格式不正确！")
	}

	auth, err := models.AuthModel.GetAuthByAuthId(authId)
	if err != nil {
		this.ErrorLog("修改登录认证 " + authId + " 失败: " + err.Error())
		this.jsonError("修改登录认证失败！")
	}
	if len(auth) == 0 {
		this.jsonError("登录认证不存在！")
	}

	ok, _ := models.AuthModel.HasSameName(authId, name)
	if ok {
		this.jsonError("登录认证名称已经存在！")
	}
	ok, _ = models.AuthModel.HasSameUsernamePrefix(authId, usernamePrefix)
	if ok {
		this.jsonError("用户名前缀已经存在！")
	}

	_, err = models.AuthModel.Update(authId, map[string]interface{}{
		"name":            name,
		"url":             url,
		"username_prefix": usernamePrefix,
		"ext_data":        extData,
	})

	if err != nil {
		this.ErrorLog("修改登录认证 " + authId + " 失败：" + err.Error())
		this.jsonError("修改登录认证失败")
	}
	this.InfoLog("修改登录认证 " + authId + " 成功")
	this.jsonSuccess("修改登录认证成功", nil, "/system/auth/list")
}

func (this *AuthController) Delete() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/auth/list")
	}
	authId := this.GetString("login_auth_id", "")
	if authId == "" {
		this.jsonError("没有选择登录认证！")
	}

	auth, err := models.AuthModel.GetAuthByAuthId(authId)
	if err != nil {
		this.ErrorLog("删除登录认证 " + authId + " 失败: " + err.Error())
		this.jsonError("删除登录认证失败")
	}
	if len(auth) == 0 {
		this.jsonError("登录认证不存在")
	}

	err = models.AuthModel.Delete(authId)
	if err != nil {
		this.ErrorLog("删除登录认证 " + authId + " 失败: " + err.Error())
		this.jsonError("删除登录认证失败")
	}

	this.InfoLog("删除登录认证 " + authId + " 成功")
	this.jsonSuccess("删除登录认证成功", nil, "/system/auth/list")
}

func (this *AuthController) Used() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/auth/list")
	}
	authId := this.GetString("login_auth_id", "")
	if authId == "" {
		this.jsonError("没有选择登录认证！")
	}

	auth, err := models.AuthModel.GetAuthByAuthId(authId)
	if err != nil {
		this.ErrorLog("登录认证 " + authId + " 启用失败: " + err.Error())
		this.jsonError("登录认证启用失败")
	}
	if len(auth) == 0 {
		this.jsonError("登录认证不存在")
	}
	_, err = models.AuthModel.SetAuthUsed(authId)
	if err != nil {
		this.ErrorLog("登录认证 " + authId + " 启用失败: " + err.Error())
		this.jsonError("登录认证启用失败")
	}

	this.InfoLog("启用登录认证 " + authId + " 成功")
	this.jsonSuccess("启用登录认证成功", nil, "/system/auth/list")
}

func (this *AuthController) Doc() {
	this.viewLayout("auth/doc", "auth")
}
