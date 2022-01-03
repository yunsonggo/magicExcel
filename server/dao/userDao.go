package dao

import (
	"2021/magicExcel/server/model"
	"2021/magicExcel/server/store"
	"sync"
)

type UserDao interface {
	// 根据密码查询用户
	QueryUserByPassword(password string) (data *model.UserModel, err error)
	// 保存用户名密码用户
	InsertUserByPassword(user *model.UserModel) (id int64, err error)
	// 根据用户名查询用户
	QueryUserByName(name string) (data *model.UserModel, err error)
	// 修改密码
	EditUserPass(id int64, password string) (err error)
}

type userDao struct {
	sync.RWMutex
}

func NewUserDao() UserDao {
	return &userDao{}
}

// 根据密码查询用户
func (ud *userDao) QueryUserByPassword(password string) (data *model.UserModel, err error) {
	user := new(model.UserModel)
	ud.RLock()
	err = store.GormDb.Where("password = ?", password).Limit(1).Find(user).Error
	ud.RUnlock()
	return user, err
}

// 保存用户名密码用户
func (ud *userDao) InsertUserByPassword(user *model.UserModel) (id int64, err error) {
	ud.RLock()
	result := store.GormDb.Create(user)
	id = user.ID
	err = result.Error
	return
}

// 根据用户名查询用户
func (ud *userDao) QueryUserByName(name string) (data *model.UserModel, err error) {
	user := new(model.UserModel)
	ud.RLock()
	err = store.GormDb.Where("name = ?", name).Limit(1).Find(user).Error
	ud.RUnlock()
	return user, err
}

// 修改密码
func (ud *userDao) EditUserPass(id int64, password string) (err error) {
	user := new(model.UserModel)
	ud.Lock()
	tx := store.GormDb.Begin()
	err = tx.Model(user).Where("id = ?", id).Update("password", password).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//err = store.GormDb.Model(user).Where("id = ?", id).Update("password", password).Error
	ud.Unlock()
	return
}
