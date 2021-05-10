package view

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	noMatter "my_demo/view/activities"
)

func RegisterActivitiesRouters(app *iris.Application){

	//先理解为前缀吧
	activitiesRouter := app.Party("/activities")

	activitiesRouter.Post("", hero.Handler(noMatter.CreateActivities))
	activitiesRouter.Put("/{cid:int}", hero.Handler(noMatter.PutActivities))
	activitiesRouter.Delete("/{cid:int}", hero.Handler(noMatter.DeleteActivities))
	activitiesRouter.Get("/list", hero.Handler(noMatter.ListActivities))
	activitiesRouter.Post("/_mget", hero.Handler(noMatter.MgetActivities))
}