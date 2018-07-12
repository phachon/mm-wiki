package main

import (
	"github.com/astaxie/beego"
	"mm-wiki/app/controllers"
	systemControllers "mm-wiki/app/modules/system/controllers"
	"mm-wiki/app/utils"
	"net/http"
	"html/template"
)

func init()  {
	initRouter()
}

func initRouter() {
	// router
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.RouterCaseSensitive = false

	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/author", &controllers.AuthorController{}, "*:Index")
	beego.AutoRouter(&controllers.AuthorController{})
	beego.AutoRouter(&controllers.MainController{})
	beego.AutoRouter(&controllers.SpaceController{})
	beego.AutoRouter(&controllers.CollectionController{})
	beego.AutoRouter(&controllers.FollowController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.DocumentController{})
	beego.AutoRouter(&controllers.PageController{})

	systemNamespace := beego.NewNamespace("/system",
		beego.NSAutoRouter(&systemControllers.MainController{}),
		beego.NSAutoRouter(&systemControllers.ProfileController{}),
		beego.NSAutoRouter(&systemControllers.UserController{}),
		beego.NSAutoRouter(&systemControllers.RoleController{}),
		beego.NSAutoRouter(&systemControllers.PrivilegeController{}),
		beego.NSAutoRouter(&systemControllers.SpaceController{}),
		beego.NSAutoRouter(&systemControllers.Space_UserController{}),
		beego.NSAutoRouter(&systemControllers.LogController{}),
		beego.NSAutoRouter(&systemControllers.EmailController{}),
		beego.NSAutoRouter(&systemControllers.LinkController{}),
		beego.NSAutoRouter(&systemControllers.AuthController{}),
		beego.NSAutoRouter(&systemControllers.ConfigController{}),
		beego.NSAutoRouter(&systemControllers.ContactController{}),
		beego.NSAutoRouter(&systemControllers.StaticController{}),
	)
	beego.AddNamespace(systemNamespace)

	beego.ErrorHandler("404", http_404)
	beego.ErrorHandler("500", http_500)

	// add template func
	beego.AddFuncMap("dateFormat", utils.NewDate().Format)
}

func http_404(rw http.ResponseWriter, req *http.Request) {
	t,_:= template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/error/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

func http_500(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.New("500.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/500.html")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}
