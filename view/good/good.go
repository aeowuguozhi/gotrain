package good

import (
	"github.com/kataras/iris"
	"my_demo/models/db"
	paramsUtils "my_demo/utils/params"
)

//Create --  post          创建
func CreateGood(ctx iris.Context){
	//将前端发来的请求体转化成 json格式的字典，便于操作
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	//拿到字典里面对应 key 的 value
	//从字典里面拿出来所有的key value
	name := params.Str("name","名称")
	price := params.Float("price","price")

	good := db.Good{
		Name: name,
		Price: price,
	}

	db.Driver.Create(&good)
	ctx.JSON(iris.Map{
		"id":good.Id,
	})

}

// Put -- put                   修改一条商品记录
func PutGood(ctx iris.Context,cid int){
	//从缓存or数据库拿到的一条商品记录
	var good db.Good
	if err := db.Driver.GetOne("good",cid,&good);err != nil{
		panic("商品不存在")
	}
	//前端发来的 请求体
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	params.Diff(&good)
	//修改对应的数据
	good.Name = params.Str("name","名称")
	good.Price = params.Float("price","price")

	//保存回去db
	db.Driver.Save(&good)
	//response 告诉前端 ok
	ctx.JSON(iris.Map{
		"id":good.Id,
	})
}

//Delete     ---  Delete     删除  cid int
func DeleteGood(ctx iris.Context,cid int){
	var good db.Good
	//从db or cache get one data
	if err := db.Driver.GetOne("good",cid,&good);err == nil{

		//db delete
		db.Driver.Delete(good)

	}

	//response
	ctx.JSON(iris.Map{
		"id":cid,
	})
}

//获取很多条记录  List --- get    获取n个商品的  id，名称，等 很少个字段
func ListGood(ctx iris.Context){
	//自定义的list，很少个但是很关键的字段
	//(比如用户进来首页面的时候，需要加载很多商品，然后前端就可以用Mget()方法从这个list里面根据Name,id来获取更加详细的商品信息)
	var lists []struct {
		Id         int   `json:"id"`
		Name string `json:"name"`
	}
	//多少条记录
	var count int
    //不知道什么意思
	table := db.Driver.Table("good")

	//拿多少条记录
	limit := ctx.URLParamIntDefault("limit", 10)
	//分页
	page := ctx.URLParamIntDefault("page", 1)

	table.Count(&count).Offset((page - 1) * limit).Limit(limit).Select("id, name").Find(&lists)
	//向前端返回 lists
	ctx.JSON(iris.Map{
		"goods": lists,
		"total": count,         //总共多少条
		"limit": limit,          //一页多少条
		"page":  page,        //当前页
	})
}

//Mget --- post    根据前端给来的 id 数组，进行获取更详细的goods信息
func MgetGood(ctx iris.Context){
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	ids := params.List("ids", "id列表")

	data := make([]interface{}, 0, len(ids))
	goods := db.Driver.GetMany("good", ids, db.Good{})
	//for   goods  ->  data
	for _,good  := range goods {
		func(data *[]interface{}) {
			*data = append(*data, paramsUtils.ModelToDict(good,[]string{"Id","Name","Price"}))
			defer func() {
				recover()
			}()
		}(&data)
	}
	//返回data
	ctx.JSON(data)
}

