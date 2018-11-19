package model

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type DataBae struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *DataBae

func (db *DataBae) Init() {
	DB = &DataBae{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}

func (db *DataBae) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parstTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local",
	)

	db, err := gorm.Open("mysql", config)

	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	setupDB(db)
	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(0)
}

func InitSelfDB() *gorm.DB {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

func InitDockerDB() *gorm.DB {
	// 当前没有配置docker
	return InitSelfDB()
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}
