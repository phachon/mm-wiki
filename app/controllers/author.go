package controllers

import (
	"strings"
	"mm-wiki/app/models"
	"github.com/astaxie/beego"
	"mm-wiki/app/utils"
	"fmt"
)

type AuthorController struct {
	BaseController
}

// login index
func (this *AuthorController) Index() {
	this.viewLayout("author/login", "author")
}

// login
func (this *AuthorController) Login()  {

	username := strings.TrimSpace(this.GetString("username"))
	password := strings.TrimSpace(this.GetString("password"))

	user, err := models.UserModel.GetUserByUsername(username)
	if err != nil {
		this.ErrorLog("查找用户失败："+err.Error())
		this.jsonError("登录出错")
	}
	if len(user) == 0 {
		this.jsonError("用户名或密码错误!")
	}
	if user["is_forbidden"] == fmt.Sprintf("%d", models.User_Forbidden_True) {
		this.jsonError("用户已被禁用!")
	}

	password = utils.Encrypt.Md5Encode(password)

	if user["password"] != password {
		this.jsonError("账号或密码错误!")
	}

	// save session
	this.SetSession("author", user)
	// save cookie
	identify := utils.Encrypt.Md5Encode(this.Ctx.Request.UserAgent() + this.GetClientIp() + password)
	passportValue := utils.Encrypt.Base64Encode(username + "@" + identify)
	passport := beego.AppConfig.String("author::passport")
	cookieExpired, _ := beego.AppConfig.Int64("author::cookie_expired")
	this.Ctx.SetCookie(passport, passportValue, cookieExpired)

	this.Ctx.Request.PostForm.Del("password")

	this.InfoLog("登录成功")
	this.jsonSuccess("登录成功", nil, "/main/index")
}

//logout
func (this *AuthorController) Logout(){
	this.InfoLog("退出成功")
	passport := beego.AppConfig.String("author::passport")
	this.Ctx.SetCookie(passport, "")
	this.SetSession("author", nil)
	this.DelSession("author")

	this.Redirect("/", 302)
}