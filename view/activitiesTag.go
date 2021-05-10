package view

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	noMatter "my_demo/view/activities"
)

func RegisterActivitiesTagRouters(app *iris.Application){

	//先理解为前缀吧
	activitiesTagRouter := app.Party("/activitiesTags")

	activitiesTagRouter.Post("", hero.Handler(noMatter.CreateActivitiesTags))
	activitiesTagRouter.Put("/{cid:int}", hero.Handler(noMatter.PutActivitiesTags))
	activitiesTagRouter.Delete("/{cid:int}", hero.Handler(noMatter.DeleteActivitiesTags))
	activitiesTagRouter.Get("/list", hero.Handler(noMatter.ListActivitiesTag))
	activitiesTagRouter.Post("/_mget", hero.Handler(noMatter.MgetActivitiesTag))
}
