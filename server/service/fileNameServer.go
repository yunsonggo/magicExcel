package service

import (
	"2021/magicExcel/server/dao"
	"2021/magicExcel/server/model"
)

type FileNameServer interface {
	// 根据文件名查询
	FindByName(name string) (data *model.FileNameModel, err error)
	// 插入文件名
	AddFileName(name, option, filePath string) (id int64, err error)
	// 获取所有文件名
	FindFileNameList() (data []model.FileNameModel, err error)
}

type fileName struct{}

func NewFileName() FileNameServer {
	return &fileName{}
}

var fd = dao.NewFileNameDao()

// 根据文件名查询
func (fn *fileName) FindByName(name string) (data *model.FileNameModel, err error) {
	return fd.QueryByName(name)
}

// 插入文件名
func (fn *fileName) AddFileName(name, option, filePath string) (id int64, err error) {
	return fd.InsertFileName(name, option, filePath)
}

// 获取所有文件名
func (fn *fileName) FindFileNameList() (data []model.FileNameModel, err error) {
	return fd.QueryFileNameList()
}
