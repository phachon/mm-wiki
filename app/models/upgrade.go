package models

import (
	"errors"
	"github.com/astaxie/beego"
	"mm-wiki/app/utils"
)

type Upgrade struct {
}

type upgradeHandleFunc func() error

type upgradeHandle struct {
	Version string
	Func    upgradeHandleFunc
}

var (
	UpgradeModel = Upgrade{}

	upgradeMap = []*upgradeHandle{}
)

// upgrade handle func
func (up *Upgrade) initHandleFunc() {
	// v0 ~ v0.1.2
	upgradeMap = append(upgradeMap, &upgradeHandle{Version: "v0.1.2", Func: up.v0ToV012})
	// v0.1.2 ~ v0.1.8
	//upgradeMap = append(upgradeMap, &upgradeHandle{Version: "v0.1.8", Func: up.v012ToV018})
	// v0.1.8 ~ v0.2.1
	//upgradeMap = append(upgradeMap, &upgradeHandle{Version: "v0.2.1", Func: up.v018ToV021})
	// v0.2.1 ~ v0.2.7
	//upgradeMap = append(upgradeMap, &upgradeHandle{Version: "v0.2.7", Func: up.v021ToV027})
	// v0.2.7 ~ v0.3.3
	//upgradeMap = append(upgradeMap, &upgradeHandle{Version: "v0.3.3", Func: up.v027ToV033})
}

// upgrade start
func (up *Upgrade) Start(dbVersion string) (err error) {
	up.initHandleFunc()

	var tmpVersion = dbVersion
	for _, upHandle := range upgradeMap {
		// upgrade now version, exit
		if tmpVersion == Version {
			break
		}
		// tmpVersion < upHandle.version
		if utils.VersionCompare.Lt(tmpVersion, upHandle.Version) {
			// upgrade handle
			err = upHandle.Func()
			if err != nil {
				beego.Error("upgrade to " + upHandle.Version + " error: " + err.Error())
				return errors.New("upgrade to " + upHandle.Version + " error: " + err.Error())
			}
			// update system database version
			err = up.upgradeAfter(upHandle.Version)
			if err != nil {
				beego.Error("upgrade to database " + upHandle.Version + " error: " + err.Error())
				return errors.New("upgrade to database " + upHandle.Version + " error: " + err.Error())
			}
			beego.Info("upgrade to " + upHandle.Version + " success")
			// update version record
			tmpVersion = upHandle.Version
		}
	}
	return nil
}

// upgrade v0.0.0 ~ v0.1.2
func (up *Upgrade) v0ToV012() (err error) {

	// 1. add privilege '/email/test'
	// INSERT INTO mw_privilege (name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES ('测试邮件服务器', 53, 'controller', 'email', 'test', 'glyphicon-list', 0, 80, unix_timestamp(now()), unix_timestamp(now()));
	privilege := map[string]interface{}{
		"name":       "测试邮件服务器",
		"type":       "controller",
		"parent_id":  53,
		"controller": "email",
		"action":     "test",
		"target":     "",
		"icon":       "glyphicon-list",
		"is_display": 0,
		"sequence":   80,
	}
	_, err = PrivilegeModel.InsertNotExists(privilege)
	if err != nil {
		return
	}

	// 2. table mw_email add field 'is_ssl'
	// alter table mw_email add `is_ssl` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否使用ssl， 0 默认不使用 1 使用' after `password`
	db := G.DB()
	db.Exec(db.AR().Raw("alter table mw_email DROP COLUMN `is_ssl`"))
	_, err = db.Exec(db.AR().Raw("alter table mw_email add `is_ssl` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否使用ssl， 0 默认不使用 1 使用' after `password`"))

	return
}

// upgrade v0.1.2 ~ v0.1.8
func (up *Upgrade) v012ToV018() error {
	return nil
}

// upgrade v0.1.8 ~ v0.2.1
func (up *Upgrade) v018ToV021() error {
	return nil
}

// upgrade v0.2.1 ~ v0.2.7
func (up *Upgrade) v021ToV027() error {
	return nil
}

// upgrade v0.2.7 ~ v0.3.3
func (up *Upgrade) v027ToV033() error {
	return nil
}

func (up *Upgrade) upgradeAfter(version string) (err error) {
	// update system version
	config, err := ConfigModel.GetConfigByKey(Config_Key_SystemVersion)
	if err != nil {
		return
	}
	if len(config) == 0 {
		configValue := map[string]interface{}{
			"name":  "系统版本号",
			"key":   "system_version",
			"value": version,
		}
		_, err = ConfigModel.Insert(configValue)
	} else {
		_, err = ConfigModel.UpdateByKey(Config_Key_SystemVersion, version)
	}

	return err
}
