package apiController

import (
	"2021/magicExcel/server/common"
	"2021/magicExcel/server/conf"
	"2021/magicExcel/server/model"
	"2021/magicExcel/server/param"
	"2021/magicExcel/server/service"
	"2021/magicExcel/server/tools"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	fs = service.NewFileName()
	es = service.NewExcelServer()
)

func UploadExcelController(ctx *gin.Context) {
	var (
		msg           string
		fileName      string
		tableName     string
		baseTableName string
		excelParam    param.UploadExcelParam
	)
	// 获取提交数据
	err := ctx.ShouldBind(&excelParam)
	if err != nil {
		zap.L().Error("获取上传参数失败", zap.Error(err))
		common.Failed(ctx, err, common.UploadFailed, "获取上传参数失败")
		return
	}
	excelFileHeader := excelParam.File
	extString := filepath.Ext(excelFileHeader.Filename)
	if excelParam.DataTag == "2" {
		tableName = excelParam.DataOption + "_" + strings.Replace(excelParam.MonthPicker, ",", "_", -1)
		fileName = tableName + extString
	} else {
		tableName = excelParam.DataOption + "_" + excelParam.MonthString
		fileName = tableName + extString
	}
	//fileName := excelFileHeader.Filename
	filePath := "public/upload/excel/" + fileName
	// 文件已存在则删除旧文件
	if tools.IsFIleExist(filePath) {
		removeErr := os.Remove(filePath)
		if removeErr != nil {
			zap.L().Error("删除已存在文件失败,文件名:"+fileName, zap.Error(err))
			msg = "删除已存在文件失败,文件名:" + fileName
		} else {
			msg = "文件:" + fileName + "已存在,已执行替换程序"
		}
	}
	// 文件保存
	saveFileErr := ctx.SaveUploadedFile(excelFileHeader, filePath)
	if saveFileErr != nil {
		msg += "_保存文件失败,文件名:" + filePath
		zap.L().Error("保存文件失败,文件名:"+filePath, zap.Error(err))
		common.Failed(ctx, err, common.SaveFileFailed, msg)
		return
	}
	// 文件名入库
	info, findErr := fs.FindByName(tableName)
	if findErr != nil {
		msg += "_保存文件名到数据库失败,请检查数据库状态"
		zap.L().Error("文件名入库失败:"+tableName, zap.Error(findErr))
		common.Failed(ctx, findErr, common.InsertDBFailed, msg)
		return
	}
	if info.ID > 0 {
		zap.L().Error("文件名已存在取消入库:"+tableName, zap.Error(err))
		err = errors.New("文件名已存在")
		var reSubParam = param.ReTableNameParam{
			TableName:   tableName,
			TableOption: excelParam.DataOption,
			FilePath:    filePath,
		}
		common.Failed(ctx, err, common.HasTableName, reSubParam)
		return
	}
	id, addErr := fs.AddFileName(tableName, excelParam.DataOption, filePath)
	if addErr != nil || id == 0 {
		msg += "_文件名插入数据库失败,请检查数据库状态"
		zap.L().Error("文件名入库失败:"+tableName, zap.Error(addErr))
		common.Failed(ctx, addErr, common.InsertDBFailed, msg)
		return
	}
	// 根据文件名创建数据表
	err = service.CreateTable(excelParam.DataOption)
	if err != nil {
		msg += "_创建数据表失败:"
		msg += tableName
		zap.L().Error(msg, zap.Error(err))
		common.Failed(ctx, err, common.CreateTableFailed, msg)
		return
	}
	// 重命名数据表
	if excelParam.DataOption == "1" {
		baseTableName = "oil_models"
	} else if excelParam.DataOption == "2" {
		baseTableName = "repair_models"
	}
	err = service.RenameTable(baseTableName, tableName)
	if err != nil {
		msg += "_数据表重命名失败:"
		msg += tableName
		zap.L().Error(msg, zap.Error(err))
		common.Failed(ctx, err, common.RenameTableFailed, msg)
		return
	}
	// 数据入库
	f, openFileErr := excelize.OpenFile(filePath)
	if openFileErr != nil {
		msg += "_打开Excel文件失败"
		msg += filePath
		zap.L().Error(msg, zap.Error(err))
		common.Failed(ctx, err, common.OpenExcelFileFailed, msg)
		return
	}
	sheetArray := f.GetSheetMap()
	rows, readErr := f.GetRows(sheetArray[1])
	if readErr != nil {
		msg += "_读取Excel文件数据失败"
		msg += filePath
		zap.L().Error(msg, zap.Error(err))
		common.Failed(ctx, err, common.ReadExcelFileFailed, msg)
		return
	}
	// 获取表头map
	var titleMap = make(map[int]string)
	for key, value := range rows[0] {
		titleMap[key] = value
	}
	// 记录数据类型转换错误
	var parseErr error
	fmt.Printf("titleMap:%+v\n", titleMap)
	if excelParam.DataOption == "1" {
		// 根据表名 批量插入燃油数据
		var data []model.OilModel
		for index, row := range rows {
			if index == 0 {
				continue
			}
			oil := model.OilModel{}
			for key, col := range row {
				switch titleMap[key] {
				case "部门":
					oil.Class = col
				case "车辆":
					oil.CarName = col
				case "日期":
					oil.DateString = col
				case "上次加油里程表数":
					oil.BackupNum = col
				case "当前里程表":
					oil.NowNum = col
				case "油品":
					oil.OilType = col
				case "加油数量":
					floatOilNum, turnErr := strconv.ParseFloat(col, 64)
					if turnErr != nil {
						parseErr = turnErr
						break
					}
					oil.OilNum = floatOilNum
				case "金额":
					floatPay, turnErr := strconv.ParseFloat(col, 64)
					if turnErr != nil {
						parseErr = turnErr
						break
					}
					oil.Pay = floatPay
				case "审批结果":
					oil.Status = col
				default:
					parseErr = errors.New("标题:" + col + "有误")
				}
			}
			if parseErr != nil {
				break
			}
			data = append(data, oil)
		}
		if parseErr != nil {
			zap.L().Error("解析Excel文件数据错误", zap.Error(parseErr))
			common.Failed(ctx, parseErr, common.ParseDataFailed, "解析数据失败")
			return
		}
		count, insertErr := es.AddOilDataByTableName(tableName, data)
		if insertErr != nil {
			zap.L().Error("数据入库失败", zap.Error(insertErr))
			common.Failed(ctx, insertErr, common.InsertDBFailed, "数据入库失败")
			return
		}
		common.Success(ctx, "http://"+conf.AppConf.AppListen+"/upload/excel/"+fileName, "上传数据成功,共处理"+strconv.FormatInt(count, 10))
		return
	} else {
		// 根据表名 批量插入维修数据
		var data []model.RepairModel
		for index, row := range rows {
			if index == 0 {
				continue
			}
			repair := model.RepairModel{}
			for key, col := range row {
				fmt.Printf("row-key:%v,row-col:%v\n", key, col)
				switch titleMap[key] {
				case "部门":
					repair.Class = col
				case "车牌号":
					repair.CarName = col
				case "维修金额":
					payFloat, turnErr := strconv.ParseFloat(col, 64)
					if turnErr != nil {
						parseErr = turnErr
						break
					}
					repair.Pay = payFloat
				case "审批结果":
					repair.Status = col
				case "完成时间":
					repair.DateString = col
				default:
					parseErr = errors.New("标题:" + col + "有误")
				}
			}
			if parseErr != nil {
				break
			}
			data = append(data, repair)
		}
		if parseErr != nil {
			zap.L().Error("解析Excel文件数据错误", zap.Error(parseErr))
			common.Failed(ctx, parseErr, common.ParseDataFailed, "解析数据失败")
			return
		}
		fmt.Printf("data:%v\n", data)
		count, insertErr := es.AddRepairDataByTableName(tableName, data)
		if insertErr != nil {
			zap.L().Error("数据入库失败", zap.Error(insertErr))
			common.Failed(ctx, insertErr, common.InsertDBFailed, "数据入库失败")
			return
		}
		common.Success(ctx, "http://"+conf.AppConf.AppListen+"/upload/excel/"+fileName, "上传数据成功,共处理"+strconv.FormatInt(count, 10))
		return
	}
}

func ResubmitTableNameController(ctx *gin.Context) {
	var (
		nameParam     param.ReTableNameParam
		msg           string
		tableName     string
		baseTableName string
		filePath      string
	)
	err := ctx.ShouldBind(&nameParam)
	if err != nil {
		zap.L().Error("获取重新入库表名参数失败", zap.Error(err))
		common.Failed(ctx, err, common.ParamFailed, "")
		return
	}
	if nameParam.TableOption == "1" {
		baseTableName = "oil_models"
	} else if nameParam.TableOption == "2" {
		baseTableName = "repair_models"
	}
	tableName = nameParam.TableName
	filePath = nameParam.FilePath
	err = service.DropTable(tableName)
	if err != nil {
		msg += "_清除数据表失败:"
		msg += tableName
		zap.L().Error(msg, zap.Error(err))
		common.Failed(ctx, err, common.CreateTableFailed, msg)
		return
	}
	// 根据文件名创建数据表
	err = service.CreateTable(nameParam.TableOption)
	if err != nil {
		msg += "_创建数据表失败:"
		msg += tableName
		zap.L().Error(msg, zap.Error(err))
		common.Failed(ctx, err, common.CreateTableFailed, msg)
		return
	}
	// 重命名数据表
	err = service.RenameTable(baseTableName, tableName)
	if err != nil {
		msg += "_数据表重命名失败:"
		msg += tableName
		zap.L().Error(msg, zap.Error(err))
		common.Failed(ctx, err, common.RenameTableFailed, msg)
		return
	}
	// 数据入库
	f, openFileErr := excelize.OpenFile(filePath)
	if openFileErr != nil {
		msg += "_打开Excel文件失败"
		msg += filePath
		zap.L().Error(msg, zap.Error(err))
		common.Failed(ctx, err, common.OpenExcelFileFailed, msg)
		return
	}
	sheetArray := f.GetSheetMap()
	rows, readErr := f.GetRows(sheetArray[1])
	if readErr != nil {
		msg += "_读取Excel文件数据失败"
		msg += filePath
		zap.L().Error(msg, zap.Error(err))
		common.Failed(ctx, err, common.ReadExcelFileFailed, msg)
		return
	}
	// 获取表头map
	var titleMap = make(map[int]string)
	for key, value := range rows[0] {
		titleMap[key] = value
	}
	// 记录数据类型转换错误
	var parseErr error
	fmt.Printf("titleMap:%+v\n", titleMap)
	if nameParam.TableOption == "1" {
		// 根据表名 批量插入燃油数据
		var data []model.OilModel
		for index, row := range rows {
			if index == 0 {
				continue
			}
			oil := model.OilModel{}
			for key, col := range row {
				switch titleMap[key] {
				case "部门":
					oil.Class = col
				case "车辆":
					oil.CarName = col
				case "日期":
					oil.DateString = col
				case "上次加油里程表数":
					oil.BackupNum = col
				case "当前里程表":
					oil.NowNum = col
				case "油品":
					oil.OilType = col
				case "加油数量":
					floatOilNum, turnErr := strconv.ParseFloat(col, 64)
					if turnErr != nil {
						parseErr = turnErr
						break
					}
					oil.OilNum = floatOilNum
				case "金额":
					floatPay, turnErr := strconv.ParseFloat(col, 64)
					if turnErr != nil {
						parseErr = turnErr
						break
					}
					oil.Pay = floatPay
				case "审批结果":
					oil.Status = col
				default:
					parseErr = errors.New("标题:" + col + "有误")
				}
			}
			if parseErr != nil {
				break
			}
			data = append(data, oil)
		}
		if parseErr != nil {
			zap.L().Error("解析Excel文件数据错误", zap.Error(parseErr))
			common.Failed(ctx, parseErr, common.ParseDataFailed, "解析数据失败")
			return
		}
		count, insertErr := es.AddOilDataByTableName(tableName, data)
		if insertErr != nil {
			zap.L().Error("数据入库失败", zap.Error(insertErr))
			common.Failed(ctx, insertErr, common.InsertDBFailed, "数据入库失败")
			return
		}
		common.Success(ctx, "http://"+conf.AppConf.AppListen+filePath, "上传数据成功,共处理"+strconv.FormatInt(count, 10))
		return
	} else {
		// 根据表名 批量插入维修数据
		var data []model.RepairModel
		for index, row := range rows {
			if index == 0 {
				continue
			}
			repair := model.RepairModel{}
			for key, col := range row {
				fmt.Printf("row-key:%v,row-col:%v\n", key, col)
				switch titleMap[key] {
				case "部门":
					repair.Class = col
				case "车牌号":
					repair.CarName = col
				case "维修金额":
					payFloat, turnErr := strconv.ParseFloat(col, 64)
					if turnErr != nil {
						parseErr = turnErr
						break
					}
					repair.Pay = payFloat
				case "审批结果":
					repair.Status = col
				case "完成时间":
					repair.DateString = col
				default:
					parseErr = errors.New("标题:" + col + "有误")
				}
			}
			if parseErr != nil {
				break
			}
			data = append(data, repair)
		}
		if parseErr != nil {
			zap.L().Error("解析Excel文件数据错误", zap.Error(parseErr))
			common.Failed(ctx, parseErr, common.ParseDataFailed, "解析数据失败")
			return
		}
		fmt.Printf("data:%v\n", data)
		count, insertErr := es.AddRepairDataByTableName(tableName, data)
		if insertErr != nil {
			zap.L().Error("数据入库失败", zap.Error(insertErr))
			common.Failed(ctx, insertErr, common.InsertDBFailed, "数据入库失败")
			return
		}
		common.Success(ctx, "http://"+conf.AppConf.AppListen+filePath, "上传数据成功,共处理"+strconv.FormatInt(count, 10))
		return
	}
}
