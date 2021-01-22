package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/xormplus/xorm"
	"go.uber.org/zap"
	"web-graduation/models/sql"
)

var db *xorm.Engine

// Init 初始化mysql数据库
func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed", zap.Error(err))
		return err
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conn"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conn"))
	db.Charset("utf8mb4")
	db.StoreEngine("innodb")
	err = db.Sync2(sql.ModelList...)
	if err != nil {
		zap.L().Error("Sync mysql failed", zap.Error(err))
		return err
	}
	// 数据导出
	//db.DumpAllToFile("./models/mysql.sql")
	return
}

// Close 关闭mysql
func Close() {
	_ = db.Close()
}
