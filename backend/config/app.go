package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/phachon/mm-wiki/global"
	klog "github.com/phachon/mm-wiki/gopkg/log"
	"github.com/phachon/mm-wiki/gopkg/upload"
	"github.com/phachon/mm-wiki/utils"
	"gopkg.in/yaml.v3"
)

var (
	appConfig *AppConfig
)

// AppConfig 配置数据结构
type AppConfig struct {
	// Global 全局配置
	Global struct {
		Env     string `yaml:"env"`      // 运行环境
		GinMode string `yaml:"gin_mode"` // gin 框架运行环境
	} `yaml:"Global"`
	// Server 服务配置
	Server struct {
		IP           string `yaml:"ip"`            // ip
		Port         int    `yaml:"port"`          // port
		ReadTimeout  int    `yaml:"read_timeout"`  // 读超时时间 ms
		WriteTimeout int    `yaml:"write_timeout"` // 写超时时间 ms
	} `yaml:"Server"`
	// Database 数据库配置
	Database map[string]DatabaseConf `yaml:"Database"`
	// Upload 上传配置
	Upload map[string]UploadConf `yaml:"Upload"`
	// Logger 日志配置
	Logger map[string]klog.Config `yaml:"Logger"`
	// Auth 登录认证配置
	Auth AuthConf `yaml:"Auth"`
	// CORS 跨域配置
	CORS CORSConf `yaml:"CORS"`
}

// DatabaseConf 数据库配置
type DatabaseConf struct {
	Host              string `yaml:"host"`                // host
	Port              int    `yaml:"port"`                // 端口
	Name              string `yaml:"name"`                // 数据库名
	User              string `yaml:"user"`                // 用户名
	Pass              string `yaml:"pass"`                // 密码
	TablePrefix       string `yaml:"table_prefix"`        // 表前缀
	ConnMaxIdle       int    `yaml:"conn_max_idle"`       // 连接最大空闲数
	ConnMaxConnection int    `yaml:"conn_max_connection"` // 最大连接数
	ConnMaxLifeTime   int    `yaml:"conn_max_lifetime"`   // 连接最大活动时间
}

// AuthConf 登录配置
type AuthConf struct {
	JwtSecret   string `yaml:"jwt_secret"`   // jwt 密匙
	ExpireHours int    `yaml:"expire_hours"` // 过期时间小时
}

// CORSConf 跨域配置
type CORSConf struct {
	AllowOrigins []string `yaml:"allow_origins"` // 允许的来源列表，为空则允许所有
	AllowMethods []string `yaml:"allow_methods"` // 允许的HTTP方法
	AllowHeaders []string `yaml:"allow_headers"` // 允许的请求头
}

// UploadConf 上传配置
type UploadConf struct {
	UploadType            upload.UplaoderName `yaml:"type"` // 上传类型
	upload.UploaderConfig `yaml:",inline"`
}

// getAppConfigPath 获取服务启动配置文件路径
//
//	-conf 传入配置文件路径
//	默认路径 ./app.yaml
func getAppConfigPath() (string, error) {
	if global.FlagVar.GetAppConfPath() != global.DefaultAppConfigPath {
		return global.FlagVar.GetAppConfPath(), nil
	}
	path, err := utils.SearchPath(global.DefaultAppConfigName, global.GetRunEnv())
	return path, err
}

// initAppConfig 初始化项目配置
func initAppConfig() {
	// 获取项目配置文件路径
	path, err := getAppConfigPath()
	if err != nil {
		panic("get app config path fail: " + err.Error())
	}
	log.Printf("[Config] get config file: %s \n", path)
	// 解析项目配置
	cfg, err := LoadAppConfig(path)
	if err != nil {
		panic("parse config fail: " + err.Error())
	}
	log.Println(fmt.Sprintf("[Config] config: %+v", cfg))
	SetAppConf(cfg)
}

// CorrectConfig 修正配置
func CorrectConfig(config *AppConfig) error {
	return nil
}

// parseAppConfigYaml 解析配置从 yaml
func parseAppConfigYaml(configPath string) (*AppConfig, error) {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	cfg := defaultAppConfig()
	if err := yaml.Unmarshal(buf, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// 默认的配置
func defaultAppConfig() *AppConfig {
	cfg := &AppConfig{}
	return cfg
}

// LoadAppConfig 从配置文件加载项目配置
func LoadAppConfig(configPath string) (*AppConfig, error) {
	cfg, err := parseAppConfigYaml(configPath)
	if err != nil {
		return nil, err
	}
	if err := CorrectConfig(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// SetGlobalAppConf 设置全局的 app 配置
func SetAppConf(config *AppConfig) {
	appConfig = config
}

// GetAppConf 设置全局的 app 配置
func GetAppConf() *AppConfig {
	return appConfig
}

// GetServerAddr 获取 Server 监听的IP和端口
func GetServerAddr() string {
	return fmt.Sprintf("%s:%d", GetAppConf().Server.IP, GetAppConf().Server.Port)
}

// GetReadTimeout 读超时时间
func GetReadTimeout() time.Duration {
	if GetAppConf().Server.ReadTimeout > 0 {
		return time.Duration(GetAppConf().Server.ReadTimeout) * time.Millisecond
	}
	return 2 * time.Second
}

// GetWriteTimeout 写超时时间
func GetWriteTimeout() time.Duration {
	if GetAppConf().Server.WriteTimeout > 0 {
		return time.Duration(GetAppConf().Server.WriteTimeout) * time.Millisecond
	}
	return 2 * time.Second
}

// GetDatabaseConf 获取数据库配置
func (ac *AppConfig) GetDatabaseConf() map[string]DatabaseConf {
	if len(GetAppConf().Database) == 0 {
		return make(map[string]DatabaseConf)
	}
	return GetAppConf().Database
}

// GetLoggerConf 获取日志配置
func (ac *AppConfig) GetLoggerConf() map[string]klog.Config {
	return ac.Logger
}

// GetAuthConf 获取登录认证配置
func (ac *AppConfig) GetAuthConf() AuthConf {
	if ac.Auth.JwtSecret == "" {
		ac.Auth.JwtSecret = "platy_admin_login_jwt"
	}
	// 默认三个小时过期时间
	if ac.Auth.ExpireHours == 0 {
		ac.Auth.ExpireHours = 3
	}
	return ac.Auth
}

// GetUploadConf 获取上传配置
func (ac *AppConfig) GetUploadConf(sceneName string) UploadConf {
	if len(ac.Upload) == 0 {
		return UploadConf{}
	}
	return ac.Upload[sceneName]
}

// GetCORSConf 获取跨域配置
func (ac *AppConfig) GetCORSConf() CORSConf {
	return ac.CORS
}
