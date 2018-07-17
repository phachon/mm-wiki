package models

import (
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const Table_Config_Name = "config"

const (
	Config_Key_MainTitle = "main_title"
	Config_Key_MainDescription = "main_description"
	Config_Key_AutoFollowDoc = "auto_follow_doc_open"
	Config_Key_SendEmail = "send_email_open"
	Config_Key_AuthLogin = "sso_open"
)

type Config struct {
	
}

var ConfigModel = Config{}

// get config by config_id
func (c *Config) GetConfigByConfigId(configId string) (config map[string]string, err error) {
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
func (c *Config) Update(configId string, configValue map[string]interface{}) (id int64, err error) {
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
func (c *Config) UpdateByKey(key string, value string) (id int64, err error) {
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
func (c *Config) GetConfigs() (configs []map[string]string, err error) {

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
func (c *Config) GetConfigByConfigIds(configIds []string) (configs []map[string]string, err error) {
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
func (c *Config) InsertBatch(insertValues []map[string]interface{}) (id int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().InsertBatch(Table_Config_Name, insertValues))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get config by config key
func (c *Config) GetConfigByKey(key string) (config map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Config_Name).Where(map[string]interface{}{
		"key": key,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	config = rs.Row()
	return
}