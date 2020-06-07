package models

import (
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const Table_Config_Name = "config"

const (
	ConfigKeyMainTitle       = "main_title"
	ConfigKeyMainDescription = "main_description"
	ConfigKeyAutoFollowdoc   = "auto_follow_doc_open"
	ConfigKeySendEmail       = "send_email_open"
	ConfigKeyAuthLogin       = "sso_open"
	ConfigKeySystemVersion   = "system_version"
	ConfigKeyFulltextSearch  = "fulltext_search_open"
	ConfigKeyDocSearchTimer  = "doc_search_timer"
	ConfigKeySystemName      = "system_name"
)

type Config struct {
}

var ConfigModel = Config{}

// get config by config_id
func (c *Config) GetConfigByConfigId(configId string) (config map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Config_Name).Where(map[string]interface{}{
		"config_id": configId,
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
	configValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Config_Name, configValue, map[string]interface{}{
		"config_id": configId,
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
	configValue["update_time"] = time.Now().Unix()
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

// get all configs key map
func (c *Config) GetConfigsKeyMap() (configMaps map[string]map[string]string, err error) {
	configMaps = make(map[string]map[string]string)
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Config_Name))
	if err != nil {
		return
	}
	configs := rs.Rows()
	for _, config := range configs {
		key, ok := config["key"]
		if ok {
			configMaps[key] = config
		}
	}
	return
}

// get config by many config_id
func (c *Config) GetConfigByConfigIds(configIds []string) (configs []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Config_Name).Where(map[string]interface{}{
		"config_id": configIds,
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

// insert config
func (c *Config) Insert(insertValue map[string]interface{}) (id int64, err error) {

	insertValue["create_time"] = time.Now().Unix()
	insertValue["update_time"] = time.Now().Unix()
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Config_Name, insertValue))
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

// get config value by config key
func (c *Config) GetConfigValueByKey(key string, defaultValue string) (value string) {

	configData, err := c.GetConfigByKey(key)
	if err != nil {
		return defaultValue
	}
	if len(configData) == 0 {
		return defaultValue
	}
	if value, ok := configData["value"]; ok {
		return value
	}
	return defaultValue
}
