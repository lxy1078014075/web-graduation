package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

// Init 初始化mysql数据库
func Init() (err error) {
	dsn:=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
		)
	db,err=sqlx.Connect("mysql",dsn)
	if err!=nil {
		zap.L().Error("connect mysql failed",zap.Error(err))
		return err
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conn"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conn"))
	return
}

// Close 关闭mysql
func Close() {
	_ = db.Close()
}