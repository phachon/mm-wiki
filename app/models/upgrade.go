package models

import (
	"mm-wiki/app/utils"
	"github.com/astaxie/beego"
	"errors"
)

type Upgrade struct {

}

type upgradeHandleFunc func() error

type upgradeHandle struct {
	Version string
	Func upgradeHandleFunc
}

var (
	UpgradeModel = Upgrade{}

	upgradeMap = []*upgradeHandle{}
)

// upgrade handle func
func (up *Upgrade) initHandleFunc()  {
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
				beego.Error("upgrade to "+upHandle.Version+" error: "+err.Error())
				return errors.New("upgrade to "+upHandle.Version+" error: "+err.Error())
			}
			beego.Info("upgrade to "+upHandle.Version+" success")
			// update version record
			tmpVersion = upHandle.Version
		}
	}
	return nil
}

// upgrade v0.0.0 ~ v0.1.2
func (up *Upgrade) v0ToV012() error {

	// 1. table mw_privilege add /email/test privilege

	// 2. table mw_email add field 'is_ssl'

	// 3. todo

	return nil
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