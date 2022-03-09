package service

import (
	"github.com/jinzhu/gorm"
	"go-article-master/dao"
)

//CreateArticle 创建帖子信息
func CreateArticle(user *dao.Article) (err error) {
	if err = dao.Mysql.Create(user).Error; err != nil {
		return err
	}
	return
}

//GetAllArticle 获取所有帖子
func GetAllArticle() (userList []*dao.Article, err error) {
	if err := dao.Mysql.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

// GetArticleById 根据id查询帖子
func GetArticleById(id string) (user dao.Article, err error) {
	err = dao.Mysql.Find(&user, id).Error
	return
}

//GetArticlePage 获取分页
func GetArticlePage(page int, pageSize int, userId int) (article []dao.Article, err error) {
	if err := dao.Mysql.Limit(pageSize).Offset((page-1)*pageSize).Find(&article, "author=?", userId).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateArticle 更新帖子信息
func UpdateArticle(user dao.Article, id string) (err error) {
	err = dao.Mysql.Model(&user).Where("id=?", id).Update(&user).Error
	return
}

// DeleteArticleById 根据id删除帖子
func DeleteArticleById(id string) (err error) {
	err = dao.Mysql.Where("id=?", id).Delete(&dao.Article{}).Error
	return
}

// AddArticleView 增加帖子阅读量
func AddArticleView(id string) (err error) {
	err = dao.Mysql.Model(&dao.Article{}).Where("id=?", id).Update("views", gorm.Expr("views + ?", 1)).Error
	return
}

//GetHotArticle 获取热门帖子
func GetHotArticle(page int, pageSize int, ArticleType int) (article []dao.Article, err error) {
	if err := dao.Mysql.Order("views DESC").Limit(pageSize).Offset((page-1)*pageSize).Find(&article, "type=?", ArticleType).Error; err != nil {
		return nil, err
	}
	return
}
