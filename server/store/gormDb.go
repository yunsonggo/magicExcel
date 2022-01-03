package store

import (
	"2021/magicExcel/server/conf"
	"2021/magicExcel/server/model"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() (err error) {
	user := conf.AppConf.MysqlUser
	pass := conf.AppConf.MysqlPass
	addr := conf.AppConf.MysqlAddr
	dbname := conf.AppConf.MysqlDbName
	dsn := user + ":" + pass + "@tcp(" + addr + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	//mysql.Open(dsn)
	GormDb, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DisableDatetimePrecision: true,
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return
	}
	return migrateTable()
}

func migrateTable() (err error) {
	return GormDb.AutoMigrate(&model.FileNameModel{}, &model.UserModel{})
}

func MigrateOilTable() (err error) {
	return GormDb.AutoMigrate(&model.OilModel{})
}

func MigrateRepairTable() (err error) {
	return GormDb.AutoMigrate(&model.RepairModel{})
}

func RenameTable(tableName, newName string) (err error) {
	return GormDb.Migrator().RenameTable(tableName, newName)
}
