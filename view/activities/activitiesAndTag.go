package activities

import (
	"fmt"
	"github.com/kataras/iris"
	"my_demo/models/db"
	paramsUtils "my_demo/utils/params"
)

//Create --  post          创建标签
func CreateActivitiesAndTags(ctx iris.Context)  {
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))

	activitiesId := params.Int("activities_id","活动ID")

	activitiesTagId := params.Int("activities_tag_id","活动标签ID")

	activitiesAndTags := db.ActivitiesAndTag{
		ActivitiesId: activitiesId,
		ActivitiesTagId: activitiesTagId,
	}

	db.Driver.Create(&activitiesAndTags)
	ctx.JSON(iris.Map{
		"id":activitiesAndTags.Id,
	})
}

// Put -- put                   修改
func PutActivitiesAndTags(ctx iris.Context,cid int){
	//从缓存or数据库拿到的一条活动记录
	var activitiesAndTag db.ActivitiesAndTag
	if err := db.Driver.GetOne("activities_and_tag",cid,&activitiesAndTag);err != nil{
		panic("活动-活动标签不存在")
	}
	//前端发来的 请求体
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	//传的修改的参数和之前的不一样才会修改
	params.Diff(&activitiesAndTag)
	//修改对应的数据
	if params.Has("activities_id"){
		activitiesAndTag.ActivitiesTagId = params.Int("activities_id","活动Id")
		activitiesAndTag.ActivitiesTagId = params.Int("activities_tag_id","活动标签Id")
	}

	//保存回去db
	db.Driver.Save(&activitiesAndTag)
	//response 告诉前端 ok
	ctx.JSON(iris.Map{
		"id": activitiesAndTag.Id,
	})
}

//Delete     ---  Delete     删除  cid int
func DeleteActivitiesAndTags(ctx iris.Context,cid int){
	var activitiesAndTag db.ActivitiesAndTag
	//从db or cache get one data
	if err := db.Driver.GetOne("activities_and_tag",cid,&activitiesAndTag);err == nil{
		//db delete
		db.Driver.Delete(activitiesAndTag)
	}

	//response
	ctx.JSON(iris.Map{
		"id":cid,
	})
}

//获取很多条记录  List --- get    获取n个活动的  id，名称，等 很少个字段
func ListActivitiesAndTag(ctx iris.Context){
	//自定义的list，很少个但是很关键的字段
	//(比如用户进来首页面的时候，需要加载很多商品，然后前端就可以用Mget()方法从这个list里面根据Name,id来获取更加详细的商品信息)
	//这里定义的struct是table的字段的集合的子集
	var lists []struct {
		Id   int   `json:"id"`
	}
	//多少条记录
	var count int
	//specify the table on which you would like to run db operations
	table := db.Driver.Table("activities_and_tag")

	//拿多少条记录
	limit := ctx.URLParamIntDefault("limit", 10)
	//分页
	page := ctx.URLParamIntDefault("page", 1)

	table.Count(&count).Offset((page - 1) * limit).Limit(limit).Select("id").Find(&lists)
	for _,id := range lists {
		fmt.Println(id)
	}

	//向前端返回 lists
	ctx.JSON(iris.Map{
		"act_tag_ids": lists,
		"total": count,         //总共多少条
		"limit": limit,          //一页多少条
		"page":  page,        //当前页
	})
}


//Mget --- post    根据前端给来的 id 数组，进行获取更详细的goods信息
func MgetActivitiesAndTag(ctx iris.Context){
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	ids := params.List("ids", "id列表")

	data := make([]interface{}, 0, len(ids))
	activitiesAndTags := db.Driver.GetMany("activities_and_tag", ids, db.ActivitiesAndTag{})
	//for   goods  ->  data
	for _,actTag  := range activitiesAndTags {
		func(data *[]interface{}) {
			//这里对应的是表的字段名
			*data = append(*data, paramsUtils.ModelToDict(actTag,[]string{"Id","ActivitiesId","ActivitiesTagId"}))
			defer func() {
				recover()
			}()
		}(&data)
	}
	//返回data
	ctx.JSON(data)
}
