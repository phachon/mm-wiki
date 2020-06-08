package app

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/fatih/color"
	"github.com/go-ego/riot/types"
	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/utils"
	"github.com/phachon/mm-wiki/app/work"
	"github.com/phachon/mm-wiki/global"
	"github.com/snail007/go-activerecord/mysql"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

var (
	defaultConf = "conf/mm-wiki.conf"

	confPath = flag.String("conf", "", "please set mm-wiki conf path")

	version = flag.Bool("version", false, "mm-wiki version")

	upgrade = flag.Bool("upgrade", false, "mm-wiki upgrade")

	Version = global.SYSTEM_VERSION

	CopyRight = beego.Str2html(global.SYSTEM_COPYRIGHT)

	StartTime = int64(0)

	RootDir = ""

	DocumentAbsDir = ""

	MarkdownAbsDir = ""

	ImageAbsDir = ""

	AttachmentAbsDir = ""

	SearchIndexAbsDir = ""
)

func init() {
	initFlag()
	poster()
	initConfig()
	initDB()
	checkUpgrade()
	initDocumentDir()
	initSearch()
	initWork()
	StartTime = time.Now().Unix()
}

// init flag
func initFlag() {
	flag.Parse()
	// --version
	if *version == true {
		fmt.Printf(Version)
		os.Exit(0)
	}
}

// poster logo
func poster() {
	fg := color.New(color.FgBlue)
	logo := `
                                            _   _      _ 
 _ __ ___    _ __ ___           __      __ (_) | | __ (_)
| '_ ' _ \  | '_ ' _ \   _____  \ \ /\ / / | | | |/ / | |
| | | | | | | | | | | | |_____|  \ V  V /  | | |   <  | |
|_| |_| |_| |_| |_| |_|           \_/\_/   |_| |_|\_\ |_|
` +
		"Author: phachon\r\n" +
		"Version: " + Version + "\r\n" +
		"Link: https://github.com/phachon/mm-wiki"
	fg.Println(logo)
}

// init beego config
func initConfig() {

	RootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println("init config error: " + err.Error())
		os.Exit(1)
	}
	confFile := *confPath
	if *confPath == "" {
		confFile = filepath.Join(RootDir, defaultConf)
	}
	ok, _ := utils.NewFile().PathIsExists(confFile)
	if ok == false {
		log.Println("conf file " + confFile + " not exists!")
		os.Exit(1)
	}
	// init config file
	beego.LoadAppConfig("ini", confFile)

	// init name
	beego.AppConfig.Set("sys.name", "mm-wiki")
	beego.BConfig.AppName = beego.AppConfig.String("sys.name")
	beego.BConfig.ServerName = beego.AppConfig.String("sys.name")

	// set static path
	beego.SetStaticPath("/static/", filepath.Join(RootDir, "./static"))
	// views path
	beego.BConfig.WebConfig.ViewsPath = filepath.Join(RootDir, "./views/")

	// session
	//beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = ".session"
	//beego.BConfig.WebConfig.Session.SessionName = "mmwikissid"
	//beego.BConfig.WebConfig.Session.SessionOn = true

	// log
	logConfigs, err := beego.AppConfig.GetSection("log")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	for adapter, config := range logConfigs {
		logs.SetLogger(adapter, config)
	}
	logs.SetLogFuncCall(true)
}

//init db
func initDB() {
	host := beego.AppConfig.String("db::host")
	port, _ := beego.AppConfig.Int("db::port")
	user := beego.AppConfig.String("db::user")
	pass := beego.AppConfig.String("db::pass")
	dbname := beego.AppConfig.String("db::name")
	dbTablePrefix := beego.AppConfig.String("db::table_prefix")
	maxIdle, _ := beego.AppConfig.Int("db::conn_max_idle")
	maxConn, _ := beego.AppConfig.Int("db::conn_max_connection")
	models.G = mysql.NewDBGroup("default")
	cfg := mysql.NewDBConfigWith(host, port, dbname, user, pass)
	cfg.SetMaxIdleConns = maxIdle
	cfg.SetMaxOpenConns = maxConn
	cfg.TablePrefix = dbTablePrefix
	cfg.TablePrefixSqlIdentifier = "__PREFIX__"
	err := models.G.Regist("default", cfg)
	if err != nil {
		logs.Error(fmt.Errorf("database error:%s,with config : %v", err, cfg))
		os.Exit(1)
	}
}

// init document dir
func initDocumentDir() {
	docRootDir := beego.AppConfig.String("document::root_dir")
	if docRootDir == "" {
		logs.Error("document root dir " + docRootDir + " is not empty!")
		os.Exit(1)
	}
	ok, _ := utils.File.PathIsExists(docRootDir)
	if !ok {
		logs.Error("document root dir " + docRootDir + " is not exists!")
		os.Exit(1)
	}

	documentAbsDir, err := filepath.Abs(docRootDir)
	if err != nil {
		logs.Error("document root dir " + docRootDir + " is error!")
		os.Exit(1)
	}

	DocumentAbsDir = documentAbsDir

	// markdown save dir
	markDownAbsDir := path.Join(documentAbsDir, "markdowns")
	// image save dir
	imagesAbsDir := path.Join(documentAbsDir, "images")
	// attachment save dir
	attachmentAbsDir := path.Join(documentAbsDir, "attachment")
	// search index dir
	searchIndexAbsDir := path.Join(documentAbsDir, "search-index")

	MarkdownAbsDir = markDownAbsDir
	ImageAbsDir = imagesAbsDir
	AttachmentAbsDir = attachmentAbsDir
	SearchIndexAbsDir = searchIndexAbsDir

	dirList := []string{MarkdownAbsDir, ImageAbsDir, AttachmentAbsDir, SearchIndexAbsDir}
	// create dir
	for _, dir := range dirList {
		ok, _ = utils.File.PathIsExists(dir)
		if !ok {
			err := os.Mkdir(dir, 0777)
			if err != nil {
				logs.Error("create document dir "+dir+" error=%s", err.Error())
				os.Exit(1)
			}
		}
	}

	// utils document
	utils.Document.MarkdownAbsDir = markDownAbsDir
	utils.Document.DocumentAbsDir = documentAbsDir

	beego.SetStaticPath("/images/", ImageAbsDir)
	// todo
	beego.SetStaticPath("/images/:space_id/:document_id/", ImageAbsDir)
}

// check upgrade
func checkUpgrade() {
	if *upgrade == true {
		logs.Info("Start checking whether MM-Wiki needs upgrading.")
		var versionDb = "v0.0.0"
		versionConf := models.ConfigModel.GetConfigValueByKey(models.ConfigKeySystemVersion, "v0.0.0")
		if versionConf != "" {
			versionDb = versionConf
		}
		logs.Info("MM-Wiki Database version：" + versionDb)
		logs.Info("MM-Wiki Now version: " + Version)

		if versionDb == Version {
			logs.Info("MM-Wiki does not need updating.")
		} else {
			logs.Info("MM-Wiki start upgrading.")
			err := models.UpgradeModel.Start(versionDb)
			if err != nil {
				logs.Error("MM-Wiki upgrade failed.")
				os.Exit(1)
			}
			logs.Info("MM-Wiki upgrade finish.")
		}
		os.Exit(0)
	}
}

func initSearch() {

	gseFile := filepath.Join(RootDir, "docs/search_dict/dictionary.txt")
	stopFile := filepath.Join(RootDir, "docs/search_dict/stop_tokens.txt")
	ok, _ := utils.File.PathIsExists(gseFile)
	if !ok {
		logs.Error("search dict file " + gseFile + " is not exists!")
		os.Exit(1)
	}
	ok, _ = utils.File.PathIsExists(stopFile)
	if !ok {
		logs.Error("search stop dict file " + stopFile + " is not exists!")
		os.Exit(1)
	}
	global.DocSearcher.Init(types.EngineOpts{
		UseStore:    true,
		StoreFolder: SearchIndexAbsDir,
		Using:       3,
		//GseDict:       "zh",
		GseDict:       gseFile,
		StopTokenFile: stopFile,
		IndexerOpts: &types.IndexerOpts{
			IndexType: types.LocsIndex,
		},
	})
}

func initWork() {
	work.DocSearchWorker.Start()
}
