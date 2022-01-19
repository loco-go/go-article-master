package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-article-master/controller"
	"go-article-master/dao"
)

func main() {
	//连接数据库
	Db := dao.Mysql
	//程序退出关闭数据库连接
	defer Db.Close()
	//绑定模型
	Db.AutoMigrate(&dao.User{}, &dao.Article{})

	r := gin.Default()
	//注册路由
	Group := r.Group("/")
	{
		//增加
		Group.POST("/users", controller.CreateUser)
		//查看用户
		Group.GET("/users", controller.GetUserList)
		Group.GET("/users/:id", controller.GetUserById)
		//修改用户
		Group.PUT("/users/:id", controller.UpdateUser)
		//删除用户
		Group.DELETE("/users/:id", controller.DeleteUserById)

		//增加帖子
		Group.POST("/article", controller.CreateArticle)
		//查看帖子分页
		Group.GET("/articles", controller.GetArticlePage)
		//查看帖子
		Group.GET("/article", controller.GetArticleById)
		//修改帖子
		Group.PUT("/article", controller.UpdateArticle)
		//删除帖子
		Group.DELETE("/article", controller.DeleteArticleById)
		//增加和减少阅读量
		Group.POST("/addArticle", controller.AddArticleViews)
		Group.POST("/subArticle", controller.SubArticleViews)
		//获取热门帖子
		Group.GET("/articles/hot", controller.GetHotArticle)
	}

	//运行8080端口
	r.Run(":8080")
}
