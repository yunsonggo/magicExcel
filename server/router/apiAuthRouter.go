package router

import (
	"2021/magicExcel/server/controller/apiController"
	"2021/magicExcel/server/middleware"
	"github.com/gin-gonic/gin"
)

func ApiAuthRouter(ag *gin.RouterGroup) {
	aag := ag.Group("/auth")
	aag.Use(middleware.ApiAuth())
	{
		// 上传数据
		aag.POST("/excel/upload", apiController.UploadExcelController)
		// 覆盖数据
		aag.POST("/excel/resubmit", apiController.ResubmitTableNameController)
		// 获取数据
		aag.POST("/online/list", apiController.OnlineListController)
		// 导出数据
		aag.POST("/online/output/list", apiController.OnlineOutputController)
		// 修改密码
		aag.POST("/user/reset", apiController.ResetUserPass)
	}
}
