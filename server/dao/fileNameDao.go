package dao

import (
	"2021/magicExcel/server/model"
	"2021/magicExcel/server/store"
	"sync"
)

type FileNameDao interface {
	// 根据文件名查询
	QueryByName(name string) (data *model.FileNameModel, err error)
	// 插入文件名
	InsertFileName(name, option, filePath string) (id int64, err error)
	// 获取所有文件名
	QueryFileNameList() (data []model.FileNameModel, err error)
}

type fileNameDao struct {
	sync.RWMutex
}

func NewFileNameDao() FileNameDao {
	return &fileNameDao{}
}

// 根据文件名查询
func (fd *fileNameDao) QueryByName(name string) (data *model.FileNameModel, err error) {
	filename := new(model.FileNameModel)
	fd.RLock()
	err = store.GormDb.Where("name = ?", name).Limit(1).Find(filename).Error
	fd.RUnlock()
	return filename, err
}

// 插入文件名
func (fd *fileNameDao) InsertFileName(name, option, filePath string) (id int64, err error) {
	filename := &model.FileNameModel{
		Name: name,
		Type: option,
		Path: filePath,
	}
	fd.Lock()
	result := store.GormDb.Create(filename)
	fd.Unlock()
	return filename.ID, result.Error
}

// 获取所有文件名
func (fd *fileNameDao) QueryFileNameList() (data []model.FileNameModel, err error) {
	err = store.GormDb.Find(&data).Error
	return
}
