/**
 * @Author: dexukong
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/08/16 16:45
 */

package dao

import (
	"context"
	"ginmall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	//err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&user).Error
	//if user == nil || err == nil {
	//	return nil, false, err
	//}
	//return user, true, nil
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).
		Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).
		First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

//GetUserById 根据 id 获取用户gorm
func (dao *UserDao) GetUserById(id uint) (user model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

// UpdateUserById 根据 id 更新用户信息
func (dao *UserDao) UpdateUserById(id uint, user model.User) (err error) {
	return dao.DB.Model(&model.User{}).Where("id=?", id).Updates(&user).Error
}
