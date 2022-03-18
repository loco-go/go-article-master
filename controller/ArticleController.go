package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/patrickmn/go-cache"
	"go-article-master/dao"
	"go-article-master/service"
	"strconv"
	"time"
)

//CreateArticle 创建帖子
func CreateArticle(c *gin.Context) {
	var article dao.Article
	c.BindJSON(&article)
	err := service.CreateArticle(&article)
	if err != nil {
		Error(c, err.Error())
	} else {
		Success(c, "新增成功", article)
	}
}

//GetArticlePage 获取帖子分页
func GetArticlePage(c *gin.Context) {
	//获取url参数
	page, _ := strconv.Atoi(c.Query("pg_id"))
	pageSize, _ := strconv.Atoi(c.Query("pg_sz"))
	userId, _ := strconv.Atoi(c.Query("user_id"))

	//进行参数校验
	validate := dao.ArticleValidator{UserId: userId, Page: page, PageSize: pageSize}
	err := validator.New().Struct(validate)
	if err != nil {
		Error(c, err.Error())
		return
	}

	//进行分页请求
	res, err := service.GetArticlePage(validate.Page, validate.PageSize, validate.UserId)
	if err != nil {
		Error(c, err.Error())
	} else {
		Success(c, "请求成功", res)
	}
	return
}

//GetArticleById 根据id获取帖子
func GetArticleById(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		res, err := service.GetAllArticle()
		if err != nil {
			Error(c, err.Error())
		} else {
			Success(c, "请求成功", res)
		}
		return
	}
	res, err := service.GetArticleById(id)
	if err != nil {
		Error(c, err.Error())
	} else {
		Success(c, "请求成功", res)
	}
}

//UpdateArticle 更新帖子
func UpdateArticle(c *gin.Context) {
	var article dao.Article
	id, ok := c.GetQuery("id")
	if !ok {
		Error(c, "无效的id参数")
		return
	}
	c.BindJSON(&article)
	if err := service.UpdateArticle(article, id); err != nil {
		Error(c, err.Error())
	} else {
		Success(c, "更新成功", article)
	}
}

// DeleteArticleById 删除帖子
func DeleteArticleById(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		Error(c, "无效的id参数")
		return
	}
	if err := service.DeleteArticleById(id); err != nil {
		Error(c, err.Error())
	} else {
		Success(c, "删除成功", id)
	}
}

// AddArticleViews 增加阅读量
func AddArticleViews(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		Error(c, "无效的id参数")
		return
	}
	err := service.AddArticleView(id)
	if err != nil {
		Error(c, err.Error())
	} else {
		Success(c, "增加阅读量成功", err)
	}
}

//GetHotArticle 获取热门帖子
func GetHotArticle(c *gin.Context) {
	//获取url参数
	//page, _ := strconv.Atoi(c.Query("pg_id"))
	//pageSize, _ := strconv.Atoi(c.Query("pg_sz"))
	//articleType, _ := strconv.Atoi(c.Query("type"))

	//进行参数校验
	validate := dao.HotValidator{Type: 1, Page: 1, PageSize: 10}
	err := validator.New().Struct(validate)
	if err != nil {
		Error(c, err.Error())
		return
	}

	//进行分页请求
	res, err := service.GetHotArticle(validate.Page, validate.PageSize, validate.Type)
	if err != nil {
		Error(c, err.Error())
	} else {
		Success(c, "请求成功", res)
	}
	return
}

// 设置超时时间和清理时间
var cacheData = cache.New(60*time.Minute, 1*time.Minute)

//GetArticleCache 进行缓存处理
func GetArticleCache(c *gin.Context) {

	//获取url参数
	//page, _ := strconv.Atoi(c.Query("pg_id"))
	//pageSize, _ := strconv.Atoi(c.Query("pg_sz"))
	//articleType, _ := strconv.Atoi(c.Query("type"))

	//进行参数校验
	validate := dao.HotValidator{Type: 1, Page: 1, PageSize: 10}
	err := validator.New().Struct(validate)
	if err != nil {
		Error(c, err.Error())
		return
	}

	//进行缓存
	var hotArticle []dao.Article
	x, found := cacheData.Get("hotArticle")
	if !found {
		//进行分页请求
		res, err := service.GetHotArticle(validate.Page, validate.PageSize, validate.Type)
		if err != nil {
			Error(c, err.Error())
		} else {
			Success(c, "请求成功", res)
		}
		cacheData.Set("hotArticle", res, cache.DefaultExpiration)
		fmt.Println("缓存更新")
		return
	} else {
		hotArticle = x.([]dao.Article)
		Success(c, "请求成功", hotArticle)
	}
}

func GetArticleRedis(c *gin.Context) {

	//获取url参数
	//page, _ := strconv.Atoi(c.Query("pg_id"))
	//pageSize, _ := strconv.Atoi(c.Query("pg_sz"))
	//articleType, _ := strconv.Atoi(c.Query("type"))

	//进行参数校验
	validate := dao.HotValidator{Type: 1, Page: 1, PageSize: 10}
	err := validator.New().Struct(validate)
	if err != nil {
		Error(c, err.Error())
		return
	}

	//进行缓存
	var hotArticle []dao.Article
	ctx := context.Background()

	val, err := dao.Rdb.Get(ctx, "hotArticle").Result()

	if err != nil {
		//进行分页请求
		res, err := service.GetHotArticle(validate.Page, validate.PageSize, validate.Type)
		if err != nil {
			Error(c, err.Error())
		} else {
			Success(c, "请求成功", res)
		}
		resJson, _ := json.Marshal(res)
		err = dao.Rdb.Set(ctx, "hotArticle", resJson, 60*time.Minute).Err()
		if err != nil {
			panic(err)
		}
		fmt.Println("redis缓存更新")
		return
	} else {
		//反序列化
		json.Unmarshal([]byte(val), &hotArticle)
		Success(c, "请求成功", hotArticle)
	}
}
