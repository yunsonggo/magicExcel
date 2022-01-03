package dao

import (
	"2021/magicExcel/server/model"
	"2021/magicExcel/server/store"
	"sync"
)

type ExcelDao interface {
	// 指定燃油数据表名插入Excel文件数据
	InsertOilDataByTableName(tableName string,data []model.OilModel) (count int64,err error)
	// 指定维修数据表名插入Excel文件数据
	InsertRepairDataByTableName(tableName string,data []model.RepairModel) (count int64,err error)
}

type excelDao struct {
	sync.RWMutex
}

func NewExcelDao() ExcelDao {
	return &excelDao{}
}

// 指定燃油数据表名插入Excel文件数据
func (ed *excelDao) InsertOilDataByTableName(tableName string, data []model.OilModel) (count int64,err error) {
	ed.Lock()
	result := store.GormDb.Table(tableName).CreateInBatches(data,100)
	ed.Unlock()
	return result.RowsAffected,result.Error
}

// 指定维修数据表名插入Excel文件数据
func (ed *excelDao) InsertRepairDataByTableName(tableName string,data []model.RepairModel) (count int64,err error) {
	ed.Lock()
	result := store.GormDb.Table(tableName).CreateInBatches(data,100)
	ed.Unlock()
	return result.RowsAffected,result.Error
}