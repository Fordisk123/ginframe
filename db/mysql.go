package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	flog "github.com/Fordisk123/ginframe/log"
	"github.com/ory/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var db *gorm.DB

func GetDb(ctx context.Context) *gorm.DB {
	scope := os.Getenv("scope")
	if scope == "true" {
		return db.WithContext(ctx)
	} else {
		return db.WithContext(ctx).Unscoped()
	}
}

func InitDb() {
	db = openMysqlDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.address"),
		viper.GetString("db.name"))
}

func openMysqlDB(username, password, addr, name string) *gorm.DB {
	var logLevel gormLogger.LogLevel
	logLevelConf := strings.TrimSpace(strings.ToLower(viper.GetString("db.log_level")))
	switch logLevelConf {
	case "info":
		{
			logLevel = gormLogger.Info
		}
	case "silent":
		{
			logLevel = gormLogger.Silent
		}
	case "warn":
		{
			logLevel = gormLogger.Warn
		}
	case "error":
		{
			logLevel = gormLogger.Error
		}
	default:
		logLevel = gormLogger.Info
	}

	newLogger := gormLogger.New(
		log.New(flog.DefaultLogger.GetWriter(), "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold:             3 * time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,        // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,            // Disable color
		},
	)

	df := viper.GetBool("db.disable_foreignKey")

	addrs := strings.Split(addr, ":")
	if len(addrs) != 2 {
		panic("db-addr format error! eg : 127.0.0.1:3306")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, addrs[0], addrs[1], name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: df,
	})

	if err != nil {
		flog.DefaultLogger.Sugar().Errorf(err.Error(), "Database connection failed. Database name: %s", name)
	} else {
		sqlDb, err := db.DB()
		if err != nil {
			panic("初始化数据库失败：" + err.Error())
		}
		sqlDb.SetMaxOpenConns(10)
		sqlDb.SetMaxIdleConns(5)
		sqlDb.SetConnMaxLifetime(30 * time.Second)
	}

	return db
}
