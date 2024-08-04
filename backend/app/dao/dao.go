// Package dal 数据层
package dao

import (
	"fmt"
	"time"

	"github.com/phachon/mm-wiki/config"
	"github.com/phachon/mm-wiki/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	dbNameMK = "mm_wiki" // mm-wiki 数据库连接
)

var dbs = make(map[string]*gorm.DB)

// register 注册一个 db
func register(dbName string, db *gorm.DB) {
	dbs[dbName] = db
}

// GetDB 获取一个 db
func GetDB(dbName string) *gorm.DB {
	return dbs[dbName]
}

// InitDB 初始化数据库实例
func InitDB() {
	dbConfig := config.GetAppConf().GetDatabaseConf()
	if len(dbConfig) == 0 {
		logger.Fatalf("[InitDB] db config is empty")
	}
	for dbName, dbConf := range dbConfig {
		gormDB, err := getDB(&dbConf)
		if err != nil {
			continue
		}
		if gormDB == nil {
			continue
		}
		register(dbName, gormDB)
	}
}

func getDB(dbConf *config.DatabaseConf) (gormDB *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConf.User,
		dbConf.Pass,
		dbConf.Host,
		dbConf.Port,
		dbConf.Name,
	)
	logger.Infof("[InitDB] db dsn=%s", dsn)
	gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbConf.TablePrefix,
			SingularTable: false,
		}},
	)
	if err != nil {
		return gormDB, err
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		return gormDB, err
	}
	if dbConf.ConnMaxIdle > 0 {
		sqlDB.SetMaxIdleConns(dbConf.ConnMaxIdle)
	}
	if dbConf.ConnMaxConnection > 0 {
		sqlDB.SetMaxOpenConns(dbConf.ConnMaxConnection)
	}
	if dbConf.ConnMaxLifeTime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(dbConf.ConnMaxLifeTime) * time.Second)
	}
	return gormDB, nil
}

// CloseDB 关闭单个 db
func CloseDB(dbName string) {
	gormDB := GetDB(dbName)
	if gormDB == nil {
		return
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		logger.Errorf("[CloseDB] dbName=%s err: %v", dbName, err)
		return
	}
	sqlDB.Close()
}

// CloseDBs 关闭所有的 db
func CloseDBs() {
	for dbName := range dbs {
		CloseDB(dbName)
	}
}
