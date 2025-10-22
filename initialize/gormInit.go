package initialize

import (
	"ZoranYuan_blog/global"
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	mysqlConfig := global.Config.Mysql

	db, err := gorm.Open(mysql.Open(mysqlConfig.Dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(mysqlConfig.LogLevel()),
	})

	if err != nil {
		return nil, errors.New("failed to create gorm client")
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	sqlDb.SetMaxOpenConns(mysqlConfig.MaxOpenConns)

	return db, nil
}
