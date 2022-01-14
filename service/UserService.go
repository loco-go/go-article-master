package service

import (
	"go-article-master/dao"
)

//CreateUser 创建用户信息
func CreateUser(user *dao.User) (err error) {
	if err = dao.Mysql.Create(user).Error; err != nil {
		return err
	}
	return
}

//GetAllUser 获取所有用户
func GetAllUser() (userList []*dao.User, err error) {
	if err := dao.Mysql.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

// GetUserById 根据id查询用户User
func GetUserById(id string) (user dao.User, err error) {
	err = dao.Mysql.Find(&user, id).Error
	return
}

// UpdateUser 更新用户信息
func UpdateUser(user dao.User, id string) (err error) {
	err = dao.Mysql.Model(&user).Where("id=?", id).Update(&user).Error
	return
}

// DeleteUserById 根据id删除user
func DeleteUserById(id string) (err error) {
	err = dao.Mysql.Where("id=?", id).Delete(&dao.User{}).Error
	return
}
