package service

import (
	"2021/magicExcel/server/dao"
	"2021/magicExcel/server/model"
	"2021/magicExcel/server/param"
	"2021/magicExcel/server/until"
	"errors"
)

type OnlineServer interface {
	// 根据条件 查询燃油数据列表
	FindOilListByParam(onlineParam *param.OnlineListParam) (list map[string]map[string]model.OilDataModel, err error)
	// 根据条件 查询维修数据列表
	FindRepairListByParam(onlineParam *param.OnlineListParam) (list map[string]map[string]model.RepairDataModel, err error)
	// 根据条件 查询综合数据列表
	FindMainListByParam(onlineParam *param.OnlineListParam) (list map[string]map[string]model.MainDataModel, err error)
}

type onlineServer struct{}

func NewOnlineServer() OnlineServer {
	return &onlineServer{}
}

var fs = NewFileName()
var od = dao.NewOnlineDao()

// 根据条件 查询燃油数据列表
func (os *onlineServer) FindOilListByParam(onlineParam *param.OnlineListParam) (list map[string]map[string]model.OilDataModel, err error) {
	var (
		// 保存目标表名
		fileMap = make(map[int]string)
		// map 计数
		i int
		// 条件月份
		paramMinMonth int64
		paramMaxMonth int64
	)
	if onlineParam.TableTag == "1" {
		paramMinMonth, paramMaxMonth, err = until.ParseFileName(onlineParam.MouthString)
		if err != nil {
			return
		}
	} else {
		paramMinMonth, paramMaxMonth, err = until.ParseFileName(onlineParam.MouthPicker)
		if err != nil {
			return
		}
	}
	// 从数据库取出所有表名
	fileNames, namesErr := fs.FindFileNameList()
	if namesErr != nil {
		err = namesErr
		return nil, err
	}
	// 遍历所有表名 寻找符合条件的表名保存到fileMap
	for _, name := range fileNames {
		// 不是燃油数据表 跳过
		if name.Type != "1" {
			continue
		}
		minMonth, maxMonth, parseErr := until.ParseFileName(name.Name)
		if parseErr != nil {
			err = parseErr
			break
		}
		// 对比条件和目标月份
		// 单月数据
		if onlineParam.TableTag == "1" {
			if paramMaxMonth == maxMonth {
				fileMap[i] = name.Name
			}
		} else {
			// 多月数据
			if minMonth <= paramMinMonth && maxMonth <= paramMaxMonth {
				fileMap[i] = name.Name
			}
		}
		i++
	}
	if err != nil {
		return nil, err
	}

	// 根据符合条件的表名 查询数据
	if len(fileMap) == 0 {
		err = errors.New("没有符合条件的数据表")
		return nil, err
	}

	for _, tableName := range fileMap {
		res, resErr := od.QueryOilListByParam(tableName)
		if resErr != nil {
			err = resErr
			break
		}
		list = make(map[string]map[string]model.OilDataModel)
		for _, info := range res {
			if info.Status != "同意" {
				continue
			}
			dataModel := model.OilDataModel{
				DateString: info.DateString,
				BackupNum:  info.BackupNum,
				NowNum:     info.NowNum,
				OilType:    info.OilType,
				Status:     info.Status,
				OilNum:     info.OilNum,
				Pay:        info.Pay,
			}
			if list[info.Class] == nil {
				list[info.Class] = make(map[string]model.OilDataModel)
			}
			if _, ok := list[info.Class][info.CarName]; ok {
				temp := list[info.Class][info.CarName]
				temp.Pay += dataModel.Pay
				temp.OilNum += dataModel.OilNum
				temp.Status = dataModel.Status
				temp.OilType = dataModel.OilType
				temp.NowNum = dataModel.NowNum
				temp.DateString = dataModel.DateString
				list[info.Class][info.CarName] = temp
			} else {
				list[info.Class][info.CarName] = dataModel
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return
}

// 根据条件 查询维修数据列表
func (os *onlineServer) FindRepairListByParam(onlineParam *param.OnlineListParam) (list map[string]map[string]model.RepairDataModel, err error) {
	var (
		// 保存目标表名
		fileMap = make(map[int]string)
		// map 计数
		i int
		// 条件月份
		paramMinMonth int64
		paramMaxMonth int64
	)
	if onlineParam.TableTag == "1" {
		paramMinMonth, paramMaxMonth, err = until.ParseFileName(onlineParam.MouthString)
		if err != nil {
			return
		}
	} else {
		paramMinMonth, paramMaxMonth, err = until.ParseFileName(onlineParam.MouthPicker)
		if err != nil {
			return
		}
	}
	// 从数据库取出所有表名
	fileNames, namesErr := fs.FindFileNameList()
	if namesErr != nil {
		err = namesErr
		return nil, err
	}
	// 遍历所有表名 寻找符合条件的表名保存到fileMap
	for _, name := range fileNames {
		// 不是维修数据表 跳过
		if name.Type != "2" {
			continue
		}
		minMonth, maxMonth, parseErr := until.ParseFileName(name.Name)
		if parseErr != nil {
			err = parseErr
			break
		}
		// 对比条件和目标月份
		// 单月数据
		if onlineParam.TableTag == "1" {
			if paramMaxMonth == maxMonth {
				fileMap[i] = name.Name
			}
		} else {
			// 多月数据
			if minMonth <= paramMinMonth && maxMonth <= paramMaxMonth {
				fileMap[i] = name.Name
			}
		}
		i++
	}
	if err != nil {
		return nil, err
	}

	// 根据符合条件的表名 查询数据
	if len(fileMap) == 0 {
		err = errors.New("没有符合条件的数据表")
		return nil, err
	}

	for _, tableName := range fileMap {
		res, resErr := od.QueryRepairListByParam(tableName)
		if resErr != nil {
			err = resErr
			break
		}
		list = make(map[string]map[string]model.RepairDataModel)
		for _, info := range res {
			if info.Status != "同意" {
				continue
			}
			dataModel := model.RepairDataModel{
				RepairPay:        info.Pay,
				RepairStatus:     info.Status,
				RepairDateString: info.DateString,
			}
			if list[info.Class] == nil {
				list[info.Class] = make(map[string]model.RepairDataModel)
			}
			if _, ok := list[info.Class][info.CarName]; ok {
				temp := list[info.Class][info.CarName]
				temp.RepairPay += dataModel.RepairPay
				temp.RepairStatus = dataModel.RepairStatus
				temp.RepairDateString = dataModel.RepairDateString
				list[info.Class][info.CarName] = temp
			} else {
				list[info.Class][info.CarName] = dataModel
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return
}

// 根据条件 查询综合数据列表
func (os *onlineServer) FindMainListByParam(onlineParam *param.OnlineListParam) (list map[string]map[string]model.MainDataModel, err error) {
	oilListData, oilErr := os.FindOilListByParam(onlineParam)
	if oilErr != nil {
		err = oilErr
		return nil, err
	}
	repairListData, repairErr := os.FindRepairListByParam(onlineParam)
	if repairErr != nil {
		err = repairErr
		return nil, err
	}
	list = make(map[string]map[string]model.MainDataModel)
	for oilClass, oilData := range oilListData {
		for oilCar, oilRow := range oilData {
			if list[oilClass] == nil {
				list[oilClass] = make(map[string]model.MainDataModel)
			}
			dataModel := model.MainDataModel{
				DateString:       oilRow.DateString,
				BackupNum:        oilRow.BackupNum,
				NowNum:           oilRow.NowNum,
				OilType:          oilRow.OilType,
				OilNum:           oilRow.OilNum,
				Pay:              oilRow.Pay,
				Status:           oilRow.Status,
				RepairPay:        0,
				RepairStatus:     "",
				RepairDateString: "",
			}
			list[oilClass][oilCar] = dataModel
		}
	}

	for repairClass, repairData := range repairListData {
		for repairCar, repairRow := range repairData {
			dataModel := list[repairClass][repairCar]
			dataModel.RepairPay = repairRow.RepairPay
			dataModel.RepairStatus = repairRow.RepairStatus
			dataModel.RepairDateString = repairRow.RepairDateString
			list[repairClass][repairCar] = dataModel
		}
	}

	return
}
