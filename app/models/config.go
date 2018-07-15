package models

import (
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const Table_Config_Name = "config"

type Config struct {
	
}

var ConfigModel = Config{}

// get config by config_id
func (u *Config) GetConfigByConfigId(configId string) (config map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Config_Name).Where(map[string]interface{}{
		"config_id":   configId,
	}))
	if err != nil {
		return
	}
	config = rs.Row()
	return
}

// update config by config_id
func (u *Config) Update(configId string, configValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	configValue["update_time"] =  time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Config_Name, configValue, map[string]interface{}{
		"config_id":   configId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update config by key
func (u *Config) UpdateByKey(key string, value string) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	configValue := map[string]interface{}{}
	configValue["value"] = value
	configValue["update_time"] =  time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Config_Name, configValue, map[string]interface{}{
		"key": key,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get all configs
func (u *Config) GetConfigs() (configs []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Config_Name))
	if err != nil {
		return
	}
	configs = rs.Rows()
	return
}

// get config by many config_id
func (u *Config) GetConfigByConfigIds(configIds []string) (configs []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Config_Name).Where(map[string]interface{}{
		"config_id":   configIds,
	}))
	if err != nil {
		return
	}
	configs = rs.Rows()
	return
}

// insert batch configs
func (u *Config) InsertBatch(insertValues []map[string]interface{}) (id int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().InsertBatch(Table_Config_Name, insertValues))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}