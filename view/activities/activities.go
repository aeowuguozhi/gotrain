package activities

import (
	"github.com/Masterminds/squirrel"
	"github.com/kataras/iris"
	ActivitiesException "my_demo/exceptions/activities"
	"my_demo/models/db"
	logUtils "my_demo/utils/log"
	paramsUtils "my_demo/utils/params"
)

/*
db.Driver.getOne("table")     数据库表名首字母小写   驼峰命名法会变成小写_小写_ eg:tableName->table_name
db.Driver.getMany("table")     数据库表名首字母小写
                                                          数据库字段名
*data = append(*data, paramsUtils.ModelToDict(activity,[]string{"Id","Title","Content","StartTime","EndTime","IsPublic","SuitPeople","Cost"}))


table.Count(&count).Offset((page - 1) * limit).Limit(limit).Select("id, title").Find(&lists)
*/


//Create --  post          创建
func CreateActivities(ctx iris.Context)  {
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))

	title := params.Str("title","活动标题")
	content := params.Str("content","活动内容")
	startTime := params.Time("startTime","活动开始时间")
	endTime := params.Time("endTime","结束时间")
	isPublc := params.Bool("isPublic","是否公布")
	suitPeople := params.Str("suitPeople","适合人群")
	cost := params.Float("cost","费用")
	tags := params.List("tag","标签")

	activies := db.Activities{
		Title: title,
		Content: content,
        StartTime: startTime,
        EndTime: endTime,
		IsPublic: isPublc,
		SuitPeople: suitPeople,
		Cost: cost,
	}

	db.Driver.Create(&activies)

	//标签挂载      活动对象， tags是标签列表 TagMnt需要activities.id
	TagMnt(activies,tags)

	ctx.JSON(iris.Map{
		"id":activies.Id,
	})
}

// Put -- put                   修改一条活动记录
func PutActivities(ctx iris.Context,cid int){
	//从缓存or数据库拿到的一条活动记录
	var activities db.Activities
	if err := db.Driver.GetOne("activities",cid,&activities);err != nil{
		//这里的报错信息使用方法是:包名.类名
		panic(ActivitiesException.ActivitiesIsNotExsit())
	}
	//前端发来的 请求体
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	params.Diff(&activities)
	//修改对应的数据
	if params.Has("title"){
		activities.Title = params.Str("title","名称")
	}
	if params.Has("content") {
		activities.Content = params.Str("content","内容")
	}
	if params.Has("startTime") {
		activities.StartTime = params.Time("startTime","开始时间")
	}
	if params.Has("endTime") {
		activities.EndTime = params.Time("endTime","结束时间")
	}
	if params.Has("isPublic") {
		activities.IsPublic = params.Bool("isPublic","是否发布")
	}
	if params.Has("suitPeople") {
		activities.SuitPeople = params.Str("suitPeople","适合人群")
	}
	if params.Has("cost") {
		activities.Cost = params.Float("cost","花费")
	}

	//保存回去db
	db.Driver.Save(&activities)
	//response 告诉前端 ok
	ctx.JSON(iris.Map{
		"id":activities.Id,
	})
}

//Delete     ---  Delete     删除  cid int
func DeleteActivities(ctx iris.Context,cid int){
	var activities db.Activities
	//从db or cache get one data
	if err := db.Driver.GetOne("activities",cid,&activities);err == nil{
		//db delete
		db.Driver.Delete(activities)
	}

	//response
	ctx.JSON(iris.Map{
		"id":cid,
	})
}

//获取很多条记录  List --- get    获取n个活动的  id，名称，等 很少个字段
func ListActivities(ctx iris.Context){
	//自定义的list，很少个但是很关键的字段
	//(比如用户进来首页面的时候，需要加载很多商品，然后前端就可以用Mget()方法从这个list里面根据Name,id来获取更加详细的商品信息)
	//这里定义的struct是table的字段的集合的子集
	var lists []struct {
		Id         int   `json:"id"`
		Title string `json:"title"`
	}
	//多少条记录
	var count int
	//specify the table on which you would like to run db operations
	table := db.Driver.Table("activities")

	//拿多少条记录
	limit := ctx.URLParamIntDefault("limit", 10)
	//分页
	page := ctx.URLParamIntDefault("page", 1)

	table.Count(&count).Offset((page - 1) * limit).Limit(limit).Select("id, title").Find(&lists)
	//向前端返回 lists
	ctx.JSON(iris.Map{
		"activities": lists,
		"total": count,         //总共多少条
		"limit": limit,          //一页多少条
		"page":  page,        //当前页
	})
}


//Mget --- post    根据前端给来的 id 数组，进行获取更详细的goods信息
func MgetActivities(ctx iris.Context){
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	ids := params.List("ids", "id列表")

	data := make([]interface{}, 0, len(ids))
	activities := db.Driver.GetMany("activities", ids, db.Activities{})
	//for   goods  ->  data
	for _,activity  := range activities {
		func(data *[]interface{}) {
			//这里对应的是数据库表的字段名 Select:选择“指定查询时要从数据库检索的字段”，默认情况下，将选择所有字段；创建/更新时，指定要保存到数据库的字段
			*data = append(*data, paramsUtils.ModelToDict(activity,[]string{"Id","Title","Content","StartTime","EndTime","IsPublic","SuitPeople","Cost"}))
			defer func() {
				recover()
			}()
		}(&data)
	}
	//返回data
	ctx.JSON(data)
}

//标签挂载
func TagMnt(activity db.Activities,t []interface{})  {
	var tags []db.ActivitiesTag
	if err:= db.Driver.Where("name in (?)",t).Find(&tags).Error;err != nil || len(tags) == 0 {
		logUtils.Println(err)
		return
	}
	//sql批处理                    table 首字母小写，驼峰会改为_
	sql := squirrel.Insert("activities_and_tag").Columns(
		"activities_id", "activities_tag_id",
	)
	for _,tag := range tags{
		sql = sql.Values(
			activity.Id,
			tag.Id,
		)
	}
	if s,args,err := sql.ToSql();err != nil {
		logUtils.Println(err)
	}else{
		if err := db.Driver.Exec(s,args...).Error;err != nil {
			logUtils.Println(err)
			return
		}
	}
}

































