package service

import (
	"2021/magicExcel/server/store"
	"errors"
)

func CreateTable(dataOption string) (err error) {
	switch dataOption {
	case "1":
		err = store.MigrateOilTable()
	case "2":
		err = store.MigrateRepairTable()
	default:
		err = errors.New("数据类型应该是:'1' 或者 '2'")
	}
	return
}

func DropTable(tableName string) (err error) {
	return store.GormDb.Migrator().DropTable(tableName)
}

func RenameTable(tableName,newName string) (err error) {
	return store.GormDb.Migrator().RenameTable(tableName,newName)
}