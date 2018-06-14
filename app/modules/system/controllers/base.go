package controllers

import (
	"encoding/json"
	"strings"
	"github.com/astaxie/beego"
	"mm-wiki/app/utils"
	"fmt"
	"mm-wiki/app/models"
)

type BaseController struct {
	beego.Controller
	UserID string
	User   map[string]string
	controllerName string
}

type JsonResponse struct {
	Code     int                    `json:"code"`
	Message  interface{}            `json:"message"`
	Data     interface{}            `json:"data"`
	Redirect map[string]interface{} `json:"redirect"`
}

// prepare
func (this *BaseController) Prepare() {
	controllerName, _ := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])

	if this.controllerName == "author" {
		return
	}

	if !this.isLogin() {
		if this.IsAjax() {
			this.jsonError("no login", nil, "/author/index")
		}else {
			this.Redirect("/author/index", 302)
		}
		this.StopRun()
	}

	this.User = this.GetSession("author").(map[string]string)
	this.UserID = this.User["user_id"]
	this.Data["user"] = this.User

	if !this.checkUserAccess() {
		if this.IsAjax() {
			this.jsonError("没有操作权限", nil, "/")
		}else {
			//this.viewError("没有操作权限", "template", "/")
		}
		return
	}

	this.Data["navName"] = this.controllerName
	this.Layout = "layouts/default.html"
}

// check is login
func (this *BaseController) isLogin() bool {
	passport := beego.AppConfig.String("author.passport")
	cookie := this.Ctx.GetCookie(passport)
	// cookie is empty
	if cookie == "" {
		return false
	}
	user := this.GetSession("author")
	// session is empty
	if user == nil {
		return false
	}
	cookieValue, _ := utils.Encrypt.Base64Decode(cookie)
	identifyList := strings.Split(cookieValue, "@")
	if cookieValue == "" || len(identifyList) != 2 {
		fmt.Println(identifyList)
		return false
	}
	username := identifyList[0]
	identify := identifyList[1]
	userValue := user.(map[string]string)

	// cookie  session name
	if username != userValue["username"] {
		return false
	}
	// UAG and IP
	if identify != utils.Encrypt.Md5Encode(this.Ctx.Request.UserAgent()+this.getClientIp()+userValue["password"]) {
		return false
	}
	// success
	return true
}


// check user access
func (this *BaseController) checkUserAccess() bool {
	if this.roleIsAdmin() || this.roleIsRoot() {
		return true
	}
	if this.controllerName == "user" {
		return false
	}
	return true
}

// view layout
func (this *BaseController) viewLayout(viewName, layout string) {
	this.Layout = "system/layouts/" + layout + ".html"
	this.TplName = "system/"+viewName + ".html"
	this.Data["title"] = ""
	this.Render()
}

// return json success
func (this *BaseController) jsonSuccess(message interface{}, data ...interface{}) {
	url := ""
	sleep := 2000
	var _data interface{}
	if len(data) > 0 {
		_data = data[0]
	}
	if len(data) > 1 {
		url = data[1].(string)
	}
	if len(data) > 2 {
		sleep = data[2].(int)
	}
	this.Data["json"] = JsonResponse{
		Code:    1,
		Message: message,
		Data:    _data,
		Redirect: map[string]interface{}{
			"url":   url,
			"sleep": sleep,
		},
	}

	j, err := json.MarshalIndent(this.Data["json"], "", "\t")
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Abort(string(j))
	}
}

// return json error
func (this *BaseController) jsonError(message interface{}, data ...interface{}) {
	url := ""
	sleep := 2000
	var _data interface{}
	if len(data) > 0 {
		_data = data[0]
	}
	if len(data) > 1 {
		url = data[1].(string)
	}
	if len(data) > 2 {
		sleep = data[2].(int)
	}
	this.Data["json"] = JsonResponse{
		Code:    0,
		Message: message,
		Data:    _data,
		Redirect: map[string]interface{}{
			"url":   url,
			"sleep": sleep,
		},
	}
	j, err := json.MarshalIndent(this.Data["json"], "", " \t")
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Abort(string(j))
	}
}

// get client ip
func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// paginator
func (this *BaseController) SetPaginator(per int, nums int64) *utils.Paginator {
	p := utils.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}

func (this *BaseController) roleIsRoot() bool {
	return this.User["role"] == fmt.Sprintf("%d", models.USER_ROLE_ROOT)
}

func (this *BaseController) roleIsAdmin() bool {
	return this.User["role"] == fmt.Sprintf("%d", models.USER_ROLE_ADMIN)
}

func (this *BaseController) roleIsUser() bool {
	return this.User["role"] == fmt.Sprintf("%d", models.USER_ROLE_USER)
}

// insert action log
func (this *BaseController) RecordLog(message string, level int) {
	//controllerName, actionName := this.GetControllerAndAction()
	//controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	//methodName := strings.ToLower(actionName)
	//userAgent := this.Ctx.Request.UserAgent()
	//referer := this.Ctx.Request.Referer()
	//getParams := this.Ctx.Request.URL.String()
	//this.Ctx.Request.ParseForm()
	//postParamsMap := map[string][]string(this.Ctx.Request.PostForm)
	//postParams, _ := json.Marshal(postParamsMap)
	//user := this.GetSession("author").(map[string]string)
	//
	//logValue := map[string]interface{}{
	//	"level": level,
	//	"controller": controllerName,
	//	"action": methodName,
	//	"get": getParams,
	//	"post": string(postParams),
	//	"message": message,
	//	"ip": this.getClientIp(),
	//	"user_agent": userAgent,
	//	"referer": referer,
	//	"user_id": user["user_id"],
	//	"username": user["username"],
	//	"create_time": time.Now().Unix(),
	//}
	//
	//models.LogModel.Insert(logValue)
}

func (this *BaseController) ErrorLog(message string)  {
	//this.RecordLog(message, models.Log_Level_Error)
}

func (this *BaseController) WarningLog(message string)  {
	//this.RecordLog(message, models.Log_Level_Warning)
}

func (this *BaseController) InfoLog(message string)  {
	//this.RecordLog(message, models.Log_Level_Info)
}

func (this *BaseController) DebugLog(message string)  {
	//this.RecordLog(message, models.Log_Level_Debug)
}