package dao

import (
	"2021/magicExcel/server/model"
	"2021/magicExcel/server/store"
	"sync"
)

type OnlineDao interface {
	// 根据条件 查询燃油数据列表
	QueryOilListByParam(tableName string) (data []model.OilModel, err error)
	// 根据条件 查询维修数据列表
	QueryRepairListByParam(tableName string) (data []model.RepairModel, err error)
}

type onlineDao struct {
	sync.RWMutex
}

func NewOnlineDao() OnlineDao {
	return &onlineDao{}
}

// 根据条件 查询燃油数据列表
func (od *onlineDao) QueryOilListByParam(tableName string) (data []model.OilModel, err error) {
	od.RLock()
	err = store.GormDb.Table(tableName).Order("date_string ASC").Find(&data).Error
	od.RUnlock()
	return
}

// 根据条件 查询维修数据列表
func (od *onlineDao) QueryRepairListByParam(tableName string) (data []model.RepairModel, err error) {
	od.RLock()
	err = store.GormDb.Table(tableName).Order("date_string ASC").Find(&data).Error
	od.RUnlock()
	return
}
