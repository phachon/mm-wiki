package controllers

import (
	"strings"
	"mm-wiki/app/models"
	"github.com/astaxie/beego"
	"mm-wiki/app/utils"
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

	user, err := models.UserModel.GetUserByName(username)
	if err != nil {
		this.ErrorLog("获取用户失败："+err.Error())
		this.jsonError("登录出错")
	}
	if len(user) == 0 {
		this.jsonError("用户名或密码错误!")
	}

	password = utils.Encrypt.Md5Encode(password)

	if user["password"] != password {
		this.jsonError("账号或密码错误!")
	}

	// save session
	this.SetSession("author", user)
	// save cookie
	identify := utils.Encrypt.Md5Encode(this.Ctx.Request.UserAgent() + this.getClientIp() + password)
	passportValue := utils.Encrypt.Base64Encode(username + "@" + identify)
	passport := beego.AppConfig.String("author.passport")
	this.Ctx.SetCookie(passport, passportValue, 3600)

	this.Ctx.Request.PostForm.Del("password")

	this.InfoLog("登录成功")
	this.jsonSuccess("登录成功", nil, "/main/index")
}

//logout
func (this *AuthorController) Logout(){
	this.InfoLog("退出成功")
	passport := beego.AppConfig.String("author.passport")
	this.Ctx.SetCookie(passport, "")
	this.SetSession("author", "")

	this.Redirect("/", 302)
}