package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-article-master/dao"
	"go-article-master/router"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var g errgroup.Group

func main() {
	//连接数据库
	Db := dao.Mysql
	//程序退出关闭数据库连接
	defer Db.Close()
	//绑定模型
	Db.AutoMigrate(&dao.User{}, &dao.Article{})

	//连接redis
	if err := dao.InitClient(); err != nil {
		return
	}

	//运行多个端口进程
	s1 := &http.Server{
		Addr:         ":8080",
		Handler:      router.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s2 := &http.Server{
		Addr:         ":8081",
		Handler:      router.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s3 := &http.Server{
		Addr:         ":8082",
		Handler:      router.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return s1.ListenAndServe()
	})

	g.Go(func() error {
		return s2.ListenAndServe()
	})

	g.Go(func() error {
		return s3.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
