package main

import (
	"github.com/astaxie/beego"
	"github.com/phachon/mm-wiki/app/utils"
	"github.com/phachon/mm-wiki/install/controllers"
	"github.com/phachon/mm-wiki/install/storage"
	"net/http"
	"os"
	"path/filepath"
)

func init() {

	storage.InstallDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	storage.RootDir = filepath.Join(storage.InstallDir, "../")

	beego.AppConfig.Set("sys.name", "mm-wiki-installer")
	beego.BConfig.AppName = beego.AppConfig.String("sys.name")
	beego.BConfig.ServerName = beego.AppConfig.String("sys.name")

	// set static path
	beego.SetStaticPath("/static/", filepath.Join(storage.InstallDir, "../static"))
	// views path
	beego.BConfig.WebConfig.ViewsPath = filepath.Join(storage.InstallDir, "../views/")

	// session
	beego.BConfig.WebConfig.Session.SessionName = "mmwikiinstallssid"
	beego.BConfig.WebConfig.Session.SessionOn = true

	// router
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.RouterCaseSensitive = false

	beego.AutoRouter(&controllers.InstallController{})
	beego.Router("/", &controllers.InstallController{}, "*:Index")
	beego.ErrorHandler("404", http_404)
	beego.ErrorHandler("500", http_500)

	// add template func
	beego.AddFuncMap("dateFormat", utils.NewDate().Format)

}

func http_404(rs http.ResponseWriter, req *http.Request) {
	rs.Write([]byte("404 not found!"))
}

func http_500(rs http.ResponseWriter, req *http.Request) {
	rs.Write([]byte("500 server error!"))
}
