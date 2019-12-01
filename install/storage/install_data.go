package storage

import "sync"

var InstallData = NewInstallData()

func NewInstallData() *installData {
	return &installData{
		lock: sync.Mutex{},
		SystemConf: &SystemConf{},
		DatabaseConf: &DatabaseConf{
			Type: Database_Type_Mysql,
		},
		AdminConf: &AdminConf{},
		CheckStatus: 0x00,
	}
}

const (
	License_Status   = 1 << 0x00 // 协议
	Env_Status       = 1 << 0x01 // 环境
	System_Status    = 1 << 0x02 // 系统
	Database_Status  = 1 << 0x03 // 数据库
	AdminUser_Status = 1 << 0x04 // 系统管理员

	Install_Check_Pass = License_Status | Env_Status | System_Status | Database_Status | AdminUser_Status

	Install_Status_Ready = 0 // 安装准备
	Install_Status_Start = 1 // 安装开始
	Install_Status_End   = 2 // 安装完成

	Install_Result_Default = 0 // 默认
	Install_Result_Failed  = 1 // 安装失败
	Install_Result_Success = 2 // 安装成功
)

const (
	Database_Type_Mysql = 0
	Database_Type_Sqlite = 1
)

// install data
type installData struct {
	lock sync.Mutex
	// system
	SystemConf *SystemConf
	// database
	DatabaseConf *DatabaseConf
	// admin user
	AdminConf *AdminConf
	// checkStatus
	CheckStatus int
}

// system info
type SystemConf struct {
	Status      int
	Addr        string
	Port        string
	DocumentDir string
}

// database info
type DatabaseConf struct {
	Type  int
	MysqlConf *MysqlConf
	SqlLiteConf *SqlLiteConf
}

// mysql conf
type MysqlConf struct {
	Host        string
	Port        string
	Name        string
	User        string
	Pass        string
	ConnMaxIdle int64
	ConnMax     int64
}

// sqlite conf
type SqlLiteConf struct {
	Path string
}

// admin user
type AdminConf struct {
	Username string
	Password string
}

func (d *installData) LicenseIsPass() bool {
	return (d.CheckStatus & License_Status) == License_Status
}

func (d *installData) EnvIsPass() bool {
	return (d.CheckStatus & Env_Status) == Env_Status
}

func (d *installData) SystemIsPass() bool {

	return (d.CheckStatus & System_Status) == System_Status
}

func (d *installData) DatabaseIsPass() bool {
	return (d.CheckStatus & Database_Status) == Database_Status
}

func (d *installData) AdminUserIsPass() bool {
	return (d.CheckStatus & AdminUser_Status) == AdminUser_Status
}

func (d *installData) CheckIsPass() bool {
	return d.CheckStatus == Install_Check_Pass
}