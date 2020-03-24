package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Unknwon/goconfig"
	"gopkg.in/gorp.v1"
)

var cfg *goconfig.ConfigFile

// InitConf -
func InitConf() {
	config, err := goconfig.LoadConfigFile("database.conf") //加载配置文件
	if err != nil {
		fmt.Println("get config file error")
		os.Exit(-1)
	}
	cfg = config
}

// InitDB -
func InitDB() *gorp.DbMap {
	InitConf()

	dbUsername, _ := cfg.GetValue("mysql", "username")
	dbPassword, _ := cfg.GetValue("mysql", "password")
	dbURL, _ := cfg.GetValue("mysql", "url")
	dbName, _ := cfg.GetValue("mysql", "dbname")
	dbEngine, _ := cfg.GetValue("mysql", "engine")
	dbEncoding, _ := cfg.GetValue("mysql", "encoding")

	dbConf := fmt.Sprintf("%v:%v@tcp(%v)/%v", dbUsername, dbPassword, dbURL, dbName)
	db, err := sql.Open("mysql", dbConf)
	if err != nil {
		log.Fatal("connect db fail")
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{dbEngine, dbEncoding}}
	return dbmap
}
