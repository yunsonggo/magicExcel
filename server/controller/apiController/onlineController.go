package apiController

import (
	"2021/magicExcel/server/common"
	"2021/magicExcel/server/conf"
	"2021/magicExcel/server/param"
	"2021/magicExcel/server/service"
	"2021/magicExcel/server/until"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var onLineServer = service.NewOnlineServer()

// 获取数据
func OnlineListController(ctx *gin.Context) {
	var onlineParam param.OnlineListParam
	err := ctx.ShouldBind(&onlineParam)
	if err != nil {
		zap.L().Error("获取在线数据参数失败", zap.Error(err))
		common.Failed(ctx, err, common.ParamFailed, "")
		return
	}
	var result interface{}
	var resultErr error
	switch onlineParam.TableOption {
	case "1":
		result, resultErr = onLineServer.FindOilListByParam(&onlineParam)
	case "2":
		result, resultErr = onLineServer.FindRepairListByParam(&onlineParam)
	case "3":
		result, resultErr = onLineServer.FindMainListByParam(&onlineParam)
	default:
		resultErr = errors.New("无效的数据类型")
		zap.L().Error(resultErr.Error(), zap.Error(resultErr))
		common.Failed(ctx, resultErr, common.ParamFailed, "")
		return
	}
	if resultErr != nil {
		zap.L().Error(resultErr.Error(), zap.Error(resultErr))
		common.Failed(ctx, resultErr, common.FindListFailed, "")
		return
	}
	common.Success(ctx, result, "ok")
	return
}

// 导出数据
func OnlineOutputController(ctx *gin.Context) {
	var outParam param.OutputParam
	err := ctx.ShouldBind(&outParam)
	if err != nil {
		zap.L().Error("获取导出数据参数失败", zap.Error(err))
		common.Failed(ctx, err, common.ParamFailed, "")
		return
	}
	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	headerTitle := []string{"部门", "车辆", "初始里程", "当前里程", "油品", "加油量", "加油金额", "百公里油耗",
		"最后日期", "审批状态", "维修金额", "维修完成日期", "维修审批"}
	err = f.SetSheetRow("Sheet1", "A1", &headerTitle)
	if err != nil {
		zap.L().Error("设置表标题失败", zap.Error(err))
		common.Failed(ctx, err, common.CreateTableFailed, "设置表标题失败")
		return
	}
	// 写入文件
	for key, rowData := range outParam.QueryTableArray {
		var payInt float64
		if len(rowData.Pay) > 0 {
			payInt, _ = strconv.ParseFloat(rowData.Pay, 64)
		}
		var oilPer float64
		if len(rowData.OilPer) > 0 {
			oilPer, _ = strconv.ParseFloat(rowData.OilPer, 64)
		}
		var oilNum float64
		if len(rowData.OilNum) > 0 {
			oilNum, _ = strconv.ParseFloat(rowData.OilNum, 64)
		}
		var repairPay float64
		if len(rowData.RepairPay) > 0 {
			repairPay, _ = strconv.ParseFloat(rowData.RepairPay, 64)
		}
		data := []interface{}{
			rowData.Class,
			rowData.Car,
			rowData.BackupNum,
			rowData.NowNum,
			rowData.OilType,
			oilNum,
			payInt,
			oilPer,
			rowData.DateString,
			rowData.Status,
			repairPay,
			rowData.RepairDateString,
			rowData.RepairStatus,
		}
		rowErr := f.SetSheetRow("Sheet1", "A"+strconv.Itoa(key+2), &data)
		if rowErr != nil {
			err = rowErr
			break
		}
	}
	if err != nil {
		zap.L().Error("写入数据失败", zap.Error(err))
		common.Failed(ctx, err, common.CreateTableFailed, "写入数据失败")
		return
	}
	// TODO 插入图标
	randomCode := until.GetRandomCode()
	saveFilename := "saved/excel/" + strconv.FormatInt(time.Now().Unix(), 10) + randomCode + ".xlsx"
	saveFilePath := "public/" + saveFilename
	f.SetActiveSheet(index)
	err = f.SaveAs(saveFilePath)
	if err != nil {
		zap.L().Error("保存文件失败", zap.Error(err))
		common.Failed(ctx, err, common.SaveFileFailed, "保存文件失败")
		return
	}
	downLoadURL := "http://" + conf.AppConf.AppListen + "/api/static/" + saveFilename
	common.Success(ctx, downLoadURL, "ok")
	return
}
