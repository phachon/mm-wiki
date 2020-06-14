package models

import (
	"github.com/phachon/mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const TablePluginName = "plugin"

type Plugin struct {
}

var PluginModel = Plugin{}

// get plugin by plugin_id
func (p *Plugin) GetPluginByPluginId(pluginId string) (plugin map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(TablePluginName).Where(map[string]interface{}{
		"plugin_id": pluginId,
	}))
	if err != nil {
		return
	}
	plugin = rs.Row()
	return
}

// update plugin by plugin_id
func (p *Plugin) Update(pluginId string, pluginValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	pluginValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(TablePluginName, pluginValue, map[string]interface{}{
		"plugin_id": pluginId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get limit plugins by search keyword
func (p *Plugin) GetPluginsByKeywordAndLimit(keyword string, limit int, number int) (plugins []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(TablePluginName).Where(map[string]interface{}{
		"title LIKE": "%" + keyword + "%",
	}).Limit(limit, number).OrderBy("plugin_id", "DESC"))
	if err != nil {
		return
	}
	plugins = rs.Rows()

	return
}

// get limit plugins
func (p *Plugin) GetPluginsByLimit(limit int, number int) (plugins []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(TablePluginName).
			Limit(limit, number).
			OrderBy("plugin_id", "DESC"))
	if err != nil {
		return
	}
	plugins = rs.Rows()

	return
}

// update plugin conf value by pluginKey
func (p *Plugin) UpdateConfValueByKey(pluginKey string, confValue string) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	pluginValue := map[string]interface{}{}
	pluginValue["conf_value"] = confValue
	pluginValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(TablePluginName, pluginValue, map[string]interface{}{
		"plugin_key": pluginKey,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update plugin conf value by pluginId
func (p *Plugin) UpdateConfValueByPluginId(pluginId string, confValue string) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	pluginValue := map[string]interface{}{}
	pluginValue["conf_value"] = confValue
	pluginValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(TablePluginName, pluginValue, map[string]interface{}{
		"plugin_id": pluginId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get all plugins
func (p *Plugin) GetPlugins() (plugins []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(TablePluginName))
	if err != nil {
		return
	}
	plugins = rs.Rows()
	return
}

// get all plugins key map
func (p *Plugin) GetPluginsKeyMap() (pluginMaps map[string]map[string]string, err error) {
	pluginMaps = make(map[string]map[string]string)
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(TablePluginName))
	if err != nil {
		return
	}
	plugins := rs.Rows()
	for _, plugin := range plugins {
		key, ok := plugin["key"]
		if ok {
			pluginMaps[key] = plugin
		}
	}
	return
}

// get plugin by many plugin_id
func (p *Plugin) GetPluginByPluginIds(pluginIds []string) (plugins []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(TablePluginName).Where(map[string]interface{}{
		"plugin_id": pluginIds,
	}))
	if err != nil {
		return
	}
	plugins = rs.Rows()
	return
}

// insert plugin
func (p *Plugin) Insert(insertValue map[string]interface{}) (id int64, err error) {

	insertValue["create_time"] = time.Now().Unix()
	insertValue["update_time"] = time.Now().Unix()
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(TablePluginName, insertValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get plugin by pluginKey
func (p *Plugin) GetPluginByKey(pluginKey string) (plugin map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(TablePluginName).Where(map[string]interface{}{
		"plugin_key": pluginKey,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	plugin = rs.Row()
	return
}

// get plugin value by plugin key
func (p *Plugin) GetConfValueByKey(pluginKey string, defaultValue string) (value string) {

	pluginData, err := p.GetPluginByKey(pluginKey)
	if err != nil {
		return defaultValue
	}
	if len(pluginData) == 0 {
		return defaultValue
	}
	if value, ok := pluginData["conf_value"]; ok {
		return value
	}
	return defaultValue
}

// get plugin count
func (p *Plugin) CountPlugins() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(TablePluginName))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get role count by keyword
func (p *Plugin) CountPluginsByKeyword(keyword string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(TablePluginName).
		Where(map[string]interface{}{
			"title LIKE": "%" + keyword + "%",
		}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}
