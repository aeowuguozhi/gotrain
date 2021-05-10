package view

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	noMatter "my_demo/view/good"
	//上面这个应该是我写的对应CRUD方法的那个package
)

//设置路由,写完路由后，就要在main.go注册这个路由
func RegisterShopRouters(app *iris.Application){

	goodRouter := app.Party("shopping/good")

	goodRouter.Post("", hero.Handler(noMatter.CreateGood))
	goodRouter.Put("/{cid:int}", hero.Handler(noMatter.PutGood))
	goodRouter.Delete("/{cid:int}", hero.Handler(noMatter.DeleteGood))
	goodRouter.Get("/list", hero.Handler(noMatter.ListGood))
	goodRouter.Post("/_mget", hero.Handler(noMatter.MgetGood))
}
