package dao

import (
	"formllm/model"
	"github.com/jinzhu/gorm"
)

type UserDao interface {
	QueryUserByPrimaryKey(tx *gorm.DB, id int64) (user *model.User, err error)
	QueryUsersByCondition(tx *gorm.DB, condition *model.UserCondition) (users []*model.User, err error)
	QueryUsersCountByCondition(tx *gorm.DB, condition *model.UserCondition) (count int64, err error)
	InsertUser(tx *gorm.DB, user *model.User) (ok bool, err error)
	InsertUsers(tx *gorm.DB, users []*model.User) (ok bool, err error)
	UpdateUsersByCondition(tx *gorm.DB, condition *model.UserCondition, selective map[string]interface{}) (ok bool, err error)
}

func (d *Dao) QueryUserByPrimaryKey(tx *gorm.DB, id int64) (user *model.User, err error) {
	user = new(model.User)
	if tx == nil {
		tx = d.DB
	}
	err = tx.Where("id = ?", id).First(user).Error
	return
}
func (d *Dao) QueryUsersByCondition(tx *gorm.DB, condition *model.UserCondition) (users []*model.User, err error) {
	users = make([]*model.User, 0)
	if tx == nil {
		tx = d.DB
	}
	tx.Where(condition.Where, condition.Params...)
	if condition.Offset != 0 {
		tx.Offset(condition.Offset)
	}
	if condition.Limit != 0 {
		tx.Limit(condition.Limit)
	}
	if condition.Order != "" {
		tx.Limit(condition.Order)
	}
	err = tx.Find(&users).Error
	return
}
func (d *Dao) QueryUsersCountByCondition(tx *gorm.DB, condition *model.UserCondition) (count int64, err error) {
	if tx == nil {
		tx = d.DB
	}
	err = tx.Where(condition.Where, condition.Params...).Count(&count).Error
	return
}
func (d *Dao) InsertUser(tx *gorm.DB, user *model.User) (ok bool, err error) {
	if tx == nil {
		tx = d.DB
	}
	err = tx.Omit("ctime", "mtime").Create(user).Error
	return
}
func (d *Dao) InsertUsers(tx *gorm.DB, users []*model.User) (ok bool, err error) {
	if tx == nil {
		tx = d.DB
	}
	err = tx.Omit("ctime", "mtime").Create(users).Error
	return
}
func (d *Dao) UpdateUsersByCondition(tx *gorm.DB, condition *model.UserCondition, selective map[string]interface{}) (ok bool, err error) {
	if tx == nil {
		tx = d.DB
	}
	err = tx.Where(condition.Where, condition.Params...).Updates(selective).Error
	return
}
