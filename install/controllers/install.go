package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/phachon/mm-wiki/app/utils"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/install/storage"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type InstallController struct {
	BaseController
}

// 安装首页
func (this *InstallController) Index() {
	this.view("install/index")
}

// 许可协议
func (this *InstallController) License() {

	if this.isPost() {
		licenseAgree := this.GetString("license_agree", "")
		if licenseAgree == "" || licenseAgree == "0" {
			this.jsonError("请先同意协议后再继续")
		}
		storage.Data.License = storage.License_Agree
		this.jsonSuccess("", nil, "/install/env")
	} else {
		bytes, _ := ioutil.ReadFile(filepath.Join(storage.RootDir, "./LICENSE"))
		license := string(bytes)
		this.Data["license"] = license
		this.Data["license_agree"] = storage.Data.License

		this.view("install/license")
	}
}

// 环境检测
func (this *InstallController) Env() {

	if this.isPost() {
		if storage.Data.Env == storage.Env_NotAccess {
			this.jsonError("环境检测未通过")
		}
		storage.Data.Env = storage.Env_Access
		this.jsonSuccess("", nil, "/install/config")
	}
	storage.Data.Env = storage.Env_Access
	//获取服务器信息
	host := utils.Misc.GetLocalIp()
	osSys := runtime.GOOS
	server := map[string]string{
		"host":        host,
		"sys":         osSys,
		"install_dir": storage.RootDir,
		"version":     global.SYSTEM_VERSION,
	}

	// 环境检测
	vm, _ := mem.VirtualMemory()
	vmTotal := vm.Total / 1024 / 1024
	cpuCount, _ := cpu.Counts(true)
	memData := map[string]interface{}{
		"name":    "内存",
		"require": "400M",
		"value":   strconv.FormatInt(int64(vmTotal), 10) + "M",
		"result":  "1",
	}
	if int(vmTotal) < 400 {
		storage.Data.Env = storage.Env_NotAccess
		memData["result"] = "0"
	}
	cpuData := map[string]interface{}{
		"name":    "CPU",
		"require": "1核",
		"value":   strconv.Itoa(cpuCount) + "核",
		"result":  "1",
	}
	if cpuCount < 1 {
		storage.Data.Env = storage.Env_NotAccess
		cpuData["result"] = "0"
	}
	envData := []map[string]interface{}{}
	envData = append(envData, memData)
	envData = append(envData, cpuData)

	// 目录文件检测
	fileTool := utils.NewFile()
	templateConfDir := map[string]string{
		"path":    "conf/template.conf",
		"require": "读/写",
		"result":  "1",
	}
	err := fileTool.IsWriterReadable(filepath.Join(storage.RootDir, templateConfDir["path"]))
	if err != nil {
		storage.Data.Env = storage.Env_NotAccess
		templateConfDir["result"] = "0"
	}

	databaseTable := map[string]string{
		"path":    "docs/databases/table.sql",
		"require": "读/写",
		"result":  "1",
	}
	err = fileTool.IsWriterReadable(filepath.Join(storage.RootDir, databaseTable["path"]))
	if err != nil {
		storage.Data.Env = storage.Env_NotAccess
		databaseTable["result"] = "0"
	}

	databaseData := map[string]string{
		"path":    "docs/databases/data.sql",
		"require": "读/写",
		"result":  "1",
	}
	err = fileTool.IsWriterReadable(filepath.Join(storage.RootDir, databaseData["path"]))
	if err != nil {
		storage.Data.Env = storage.Env_NotAccess
		databaseData["result"] = "0"
	}

	viewsDir := map[string]string{
		"path":    "views",
		"require": "存在且不为空",
		"result":  "1",
	}
	isEmpty := utils.File.PathIsEmpty(filepath.Join(storage.RootDir, viewsDir["path"]))
	if isEmpty == true {
		storage.Data.Env = storage.Env_NotAccess
		viewsDir["result"] = "0"
	}

	staticDir := map[string]string{
		"path":    "static",
		"require": "存在且不为空",
		"result":  "1",
	}
	isEmpty = utils.File.PathIsEmpty(filepath.Join(storage.RootDir, staticDir["path"]))
	if isEmpty == true {
		storage.Data.Env = storage.Env_NotAccess
		staticDir["result"] = "0"
	}

	dirData := []map[string]string{}
	dirData = append(dirData, templateConfDir)
	dirData = append(dirData, databaseTable)
	dirData = append(dirData, databaseData)
	dirData = append(dirData, viewsDir)
	dirData = append(dirData, staticDir)

	this.Data["server"] = server
	this.Data["envData"] = envData
	this.Data["dirData"] = dirData
	this.view("install/env")
}

// 系统配置
func (this *InstallController) Config() {

	if this.isPost() {
		addr := strings.TrimSpace(this.GetString("addr", ""))
		documentDir := strings.TrimSpace(this.GetString("document_dir", ""))
		port, _ := this.GetInt32("port", 0)

		if addr == "" {
			this.jsonError("addr 不能为空，默认请填写 0.0.0.0")
		}
		if port == 0 {
			this.jsonError("启动端口不能为空")
		}
		if port > int32(65535) {
			this.jsonError("端口超出范围")
		}
		if documentDir == "" {
			this.jsonError("文档保存目录不能为空")
		}
		if !filepath.IsAbs(documentDir) {
			this.jsonError("文档保存目录不是绝对路径")
		}
		docAbsDir, err := filepath.Abs(documentDir)
		if err != nil {
			this.jsonError("文档保存目录错误!")
		}
		ok, _ := utils.File.PathIsExists(docAbsDir)
		if !ok {
			this.jsonError("文档保存目录不存在!")
		}

		storage.Data.SystemConf = map[string]string{
			"addr":         addr,
			"port":         strconv.FormatInt(int64(port), 10),
			"document_dir": documentDir,
		}
		storage.Data.System = storage.Sys_Access
		this.jsonSuccess("", nil, "/install/database")
	}

	sysConf := storage.Data.SystemConf
	this.Data["sysConf"] = sysConf
	this.view("install/config")
}

// 数据库配置
func (this *InstallController) Database() {

	if !this.isPost() {
		this.Data["databaseConf"] = storage.Data.DatabaseConf
		this.viewLayoutTitle("mm-wiki-安装-数据库配置", "install/database", "install")
		return
	}

	host := strings.TrimSpace(this.GetString("host", ""))
	port := strings.TrimSpace(this.GetString("port", ""))
	name := strings.TrimSpace(this.GetString("name", ""))
	user := strings.TrimSpace(this.GetString("user", ""))
	pass := strings.TrimSpace(this.GetString("pass", ""))
	connMaxIdle := strings.TrimSpace(this.GetString("conn_max_idle", "0"))
	connMaxConn := strings.TrimSpace(this.GetString("conn_max_connection", "0"))
	adminName := strings.TrimSpace(this.GetString("admin_name", ""))
	adminPass := strings.TrimSpace(this.GetString("admin_pass", ""))

	if host == "" {
		this.jsonError("数据库 host 不能为空！")
	}
	if port == "" {
		this.jsonError("数据库端口不能为空！")
	}
	if name == "" {
		this.jsonError("数据库名不能为空！")
	}
	if user == "" {
		this.jsonError("数据库用户名不能为空！")
	}
	if pass == "" {
		this.jsonError("数据库密码不能为空！")
	}
	if connMaxIdle == "0" {
		this.jsonError("数据库连接数不能为0！")
	}
	if connMaxConn == "0" {
		this.jsonError("最大连接数不能为0！")
	}
	if adminName == "" {
		this.jsonError("超级管理员用户名不能为空！")
	} else {
		v := validation.Validation{}
		if !v.AlphaNumeric(adminName, "admin_name").Ok {
			this.jsonError("用户名格式不正确！")
		}
	}

	if adminPass == "" {
		this.jsonError("超级管理员密码不能为空！")
	}

	storage.Data.DatabaseConf = map[string]string{
		"host":                host,
		"port":                port,
		"name":                name,
		"user":                user,
		"pass":                pass,
		"conn_max_idle":       connMaxIdle,
		"conn_max_connection": connMaxConn,
		"admin_name":          adminName,
		"admin_pass":          adminPass,
	}
	storage.Data.Database = storage.Database_Access
	this.jsonSuccess("", nil, "/install/ready")
}

// 安装准备
func (this *InstallController) Ready() {

	if this.isPost() {
		if (storage.Data.License != storage.License_Agree) ||
			(storage.Data.Env != storage.Env_Access) ||
			(storage.Data.System != storage.Sys_Access) ||
			(storage.Data.Database != storage.Database_Access) {
			this.jsonError("请先完成安装准备")
		}
		storage.StartInstall()
		this.jsonSuccess("", nil, "/install/end")
	}

	// 协议
	licenseConf := map[string]interface{}{
		"name":   "许可协议",
		"value":  "同意",
		"result": "1",
		"url":    "/install/license",
	}
	if storage.Data.License != storage.License_Agree {
		licenseConf["value"] = "未同意"
		licenseConf["result"] = "0"
	}
	//环境检测
	envConf := map[string]interface{}{
		"name":   "环境检测",
		"value":  "通过",
		"result": "1",
		"url":    "/install/env",
	}
	if storage.Data.Env != storage.Env_Access {
		envConf["value"] = "未通过"
		envConf["result"] = "0"
	}
	//系统配置
	sysConf := map[string]interface{}{
		"name":   "系统配置",
		"value":  "完成",
		"result": "1",
		"url":    "/install/config",
	}
	if storage.Data.System != storage.Sys_Access {
		sysConf["value"] = "未完成"
		sysConf["result"] = "0"
	}
	//数据库配置
	databaseConf := map[string]interface{}{
		"name":   "数据库配置",
		"value":  "完成",
		"result": "1",
		"url":    "/install/database",
	}
	if storage.Data.Database != storage.Database_Access {
		databaseConf["value"] = "未完成"
		databaseConf["result"] = "0"
	}

	readyConf := []map[string]interface{}{}
	readyConf = append(readyConf, licenseConf)
	readyConf = append(readyConf, envConf)
	readyConf = append(readyConf, sysConf)
	readyConf = append(readyConf, databaseConf)

	this.Data["readyConf"] = readyConf
	this.view("install/ready")
}

// 安装完成
func (this *InstallController) End() {

	if storage.Data.Status == storage.Install_Ready {
		this.Redirect("/install/ready", 302)
	}

	this.view("install/end")
}

// 获取状态
func (this *InstallController) Status() {

	data := map[string]interface{}{
		"status":     storage.Data.Status,
		"is_success": storage.Data.IsSuccess,
		"result":     storage.Data.Result,
	}

	this.jsonSuccess("ok", data)
}
