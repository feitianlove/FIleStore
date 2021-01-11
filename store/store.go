package store

import (
	"fmt"
	"github.com/feitianlove/FIleStore/config"
	"github.com/feitianlove/FIleStore/models"
	"github.com/feitianlove/golib/common/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"time"
)

//mysql
type Store struct {
	db *gorm.DB
}

//初始化
func NewStore(conf *config.Config) (*Store, error) {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
			conf.MysqlConf.User,
			conf.MysqlConf.Pass,
			conf.MysqlConf.Host,
			conf.MysqlConf.Port,
			conf.MysqlConf.Database))
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxOpenConns(10)
	db.LogMode(true)
	db.SetLogger(logger.Mysql)
	db.DB().SetConnMaxLifetime(60 * time.Second)
	db.AutoMigrate(&models.TblFile{})
	return &Store{db}, nil
}
