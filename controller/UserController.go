package controller

import (
	"github.com/gin-gonic/gin"
	"go-article-master/dao"
	"go-article-master/service"
)

//CreateUser 创建用户信息
func CreateUser(c *gin.Context) {
	var user dao.User
	c.BindJSON(&user)
	err := service.CreateUser(&user)
	if err != nil {
		Error(c, err.Error())
	} else {
		Success(c, "新增成功", user)
	}
}

//GetUserList 获取全部用户信息
func GetUserList(c *gin.Context) {
	todoList, err := service.GetAllUser()
	if err != nil {
		Error(c, err.Error())
	} else {
		Success(c, "请求成功", todoList)
	}
}

//GetUserById 根据id获取用户
func GetUserById(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		todoList, err := service.GetAllUser()
		if err != nil {
			Error(c, err.Error())
		} else {
			Success(c, "请求成功", todoList)
		}
		return
	}
	todoList, err := service.GetUserById(id)
	if err != nil {
		Error(c, err.Error())
	} else {
		Success(c, "请求成功", todoList)
	}
}

//UpdateUser 更新用户
func UpdateUser(c *gin.Context) {
	var user dao.User
	id, ok := c.GetQuery("id")
	if !ok {
		Error(c, "无效的id参数")
		return
	}
	c.BindJSON(&user)
	if err := service.UpdateUser(user, id); err != nil {
		Error(c, err.Error())
	} else {
		Success(c, "更新成功", user)
	}
}

// DeleteUserById 删除用户
func DeleteUserById(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		Error(c, "无效的id参数")
		return
	}
	if err := service.DeleteUserById(id); err != nil {
		Error(c, err.Error())
	} else {
		Success(c, "删除成功", id)
	}
}
