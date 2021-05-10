package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"my_demo/core/cache"
	"my_demo/middlewares"
	"my_demo/models/db"
	"my_demo/view"
)

func initRouter(app *iris.Application){
	view.RegisterClassifyRouters(app)
	view.RegisterCourseRouters(app)
	//注册我写的good的路由
	view.RegisterShopRouters(app)
	view.RegisterActivitiesRouters(app)
	view.RegisterActivitiesTagRouters(app)
	view.RegisterActivitiesAndTagRouters(app)

}

func main(){
	app := iris.New()

	app.UseGlobal(middlewares.AbnormalHandle,middlewares.RequestLogHandle)
	initRouter(app)

	db.InitDB()
	cache.InitRedisPool()

	app.Run(iris.Addr(":8085"),iris.WithoutServerError(iris.ErrServerClosed))
}

