package service

import (
	"2021/magicExcel/server/dao"
	"2021/magicExcel/server/model"
)

type ExcelServer interface {
	// 指定燃油数据表名插入Excel文件数据
	AddOilDataByTableName(tableName string,data []model.OilModel) (count int64,err error)
	// 指定维修数据表名插入Excel文件数据
	AddRepairDataByTableName(tableName string,data []model.RepairModel) (count int64,err error)
}

type excelServer struct {}

var ed = dao.NewExcelDao()

func NewExcelServer() ExcelServer {
	return &excelServer{}
}

// 指定燃油数据表名插入Excel文件数据
func (es *excelServer) AddOilDataByTableName(tableName string,data []model.OilModel) (count int64,err error) {
	return ed.InsertOilDataByTableName(tableName,data)
}

// 指定维修数据表名插入Excel文件数据
func (es *excelServer) AddRepairDataByTableName(tableName string,data []model.RepairModel) (count int64,err error) {
	return ed.InsertRepairDataByTableName(tableName,data)
}