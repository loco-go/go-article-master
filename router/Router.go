package router

import (
	"github.com/gin-gonic/gin"
	"go-article-master/controller"
)

func Router() *gin.Engine {

	r := gin.Default()
	//注册路由
	Group := r.Group("/")
	{
		//增加用户
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
		//获取缓存热门帖子
		Group.GET("/articles/hotCache", controller.GetArticleCache)

		Group.GET("/articles/hotRedis", controller.GetArticleRedis)
	}
	return r
}
