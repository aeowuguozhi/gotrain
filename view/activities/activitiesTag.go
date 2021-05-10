package activities

import (
	"github.com/kataras/iris"
	"my_demo/models/db"
	paramsUtils "my_demo/utils/params"
	ActivitiesTagException "my_demo/exceptions/activities"
)

//测试成功
//Create --  post          创建标签
func CreateActivitiesTags(ctx iris.Context)  {
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))

	name := params.Str("name","标签名称")

	tags := db.ActivitiesTag{
		Name: name,
	}

	db.Driver.Create(&tags)
	ctx.JSON(iris.Map{
		"id":tags.Id,
	})
}

// Put -- put                   修改
func PutActivitiesTags(ctx iris.Context,cid int){
	//从缓存or数据库拿到的一条活动记录
	var tags db.ActivitiesTag
	if err := db.Driver.GetOne("activities_tag",cid,&tags);err != nil{
		panic(ActivitiesTagException.ActivitiesTagIsNotExsit())
	}
	//前端发来的 请求体
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	//传的修改的参数和之前的不一样才会修改
	params.Diff(&tags)
	//修改对应的数据
	if params.Has("name"){
		tags.Name = params.Str("name","活动标签名称")
	}

	//保存回去db
	db.Driver.Save(&tags)
	//response 告诉前端 ok
	ctx.JSON(iris.Map{
		"id":tags.Id,
	})
}

//Delete     ---  Delete     删除  cid int
func DeleteActivitiesTags(ctx iris.Context,cid int){
	var tags db.ActivitiesTag
	//从db or cache get one data
	if err := db.Driver.GetOne("activities_tag",cid,&tags);err == nil{
		//db delete
		db.Driver.Delete(tags)
	}

	//response
	ctx.JSON(iris.Map{
		"id":cid,
	})
}

//获取很多条记录  List --- get    获取n个活动的  id，名称，等 很少个字段
func ListActivitiesTag(ctx iris.Context){
	//自定义的list，很少个但是很关键的字段
	//(比如用户进来首页面的时候，需要加载很多商品，然后前端就可以用Mget()方法从这个list里面根据Name,id来获取更加详细的商品信息)
	var lists []struct {
		Id         int   `json:"id"`
	}
	//多少条记录
	var count int
	//specify the table on which you would like to run db operations
	table := db.Driver.Table("activities_tag")

	//拿多少条记录
	limit := ctx.URLParamIntDefault("limit", 100)
	//分页
	page := ctx.URLParamIntDefault("page", 1)

	table.Count(&count).Offset((page - 1) * limit).Limit(limit).Select("id").Find(&lists)
	//向前端返回 lists
	ctx.JSON(iris.Map{
		"activitiesTags": lists,
		"total": count,         //总共多少条
		"limit": limit,          //一页多少条
		"page":  page,        //当前页
	})
}

//测试成功
//Mget --- post    根据前端给来的 id 数组，进行获取更详细的goods信息
//db.Driver.getOne("table")    这里的table是数据库表的 name 首字母小写
func MgetActivitiesTag(ctx iris.Context){
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	ids := params.List("ids", "id列表")

	data := make([]interface{}, 0, len(ids))
	activitiesTags := db.Driver.GetMany("activities_tag", ids, db.ActivitiesTag{})
	//for   goods  ->  data
	for _,tag  := range activitiesTags {
		func(data *[]interface{}) {
			//这里对应的是表的字段名
			*data = append(*data, paramsUtils.ModelToDict(tag,[]string{"Id","Name"}))
			defer func() {
				recover()
			}()
		}(&data)
	}
	//返回data
	ctx.JSON(data)
}