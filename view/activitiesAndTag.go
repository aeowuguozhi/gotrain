package view

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	noMatter "my_demo/view/activities"
)

func RegisterActivitiesAndTagRouters(app *iris.Application){

	//先理解为前缀吧
	activitiesAndTagRouter := app.Party("/activitiesAndTags")

	activitiesAndTagRouter.Post("", hero.Handler(noMatter.CreateActivitiesAndTags))
	activitiesAndTagRouter.Put("/{cid:int}", hero.Handler(noMatter.PutActivitiesAndTags))
	activitiesAndTagRouter.Delete("/{cid:int}", hero.Handler(noMatter.DeleteActivitiesAndTags))
	activitiesAndTagRouter.Get("/list", hero.Handler(noMatter.ListActivitiesAndTag))
	activitiesAndTagRouter.Post("/_mget", hero.Handler(noMatter.MgetActivitiesAndTag))
}