package db

//设计参考csdn：https://blog.csn.net/thc1987/article/details/80426063?utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7EBlogCommendFromMachineLearnPai2%7Edefault-1.control&dist_request_id=&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7EBlogCommendFromMachineLearnPai2%7Edefault-1.control
/*
SPU
SPU(Standard Product Unit)：标准化产品单元。是商品信息聚合的最小单位，是一组可复用、易检索的标准化信息的集合，该集合描述了一个产品的特性。通俗点讲，属性值、特性相同的商品就可以称为一个SPU。

SKU
SKU=Stock Keeping Unit（库存量单位）。即库存进出计量的基本单元，可以是以件，盒，托盘等为单位。

举个例子：iPhone6是一个SPU，iPhone6 32G 白色是一个SKU，iPhone6 128G 白色是另一个SKU。
*/


//SPU  Standard Product Unit
type Good struct{
	ID int `gorm:"primary_key" json:"id"`                        //1         2

	//商品名称
	Name string `json:"name" gorm:"not null"`                    //苹果       衬衫

	//品牌
	BrandID int `json:"brand_id"`                 //FK

	//分类 属于什么类别
	ClassifyID int `json:"classify_id"`            //FK
}

//规格
type Standard struct{
	ID int `gorm:"primary_key" json:"id"`                      //1        2

	//规格名
	Name string `json:"name"`                                  //直径      领口
}

//规格值
type StandardValue struct {
	ID int `gorm:"primary_key" json:"id"`                    // 1          2

	//规格ID
	StandardID int `json:"standard_id"`          //FK           1          2

	//规格值
	Value string `json:"value"`                     //          5cm        圆领
}

//SPU-规格    DB+字段名
type GoodStandard struct{
	ID int `gorm:"primary_key" json:"id"`        //             1               2

	//商品ID
	GoodID int `gorm:"primary_key" json:"good_id"`       //     1苹果            2衬衫     3手机

	//规格ID
	StandardID int `gorm:"primary_key" json:"standard_id"`  //  1直径            2圆领     4内存
}

//品牌
type Brand struct {
	ID int `gorm:"primary_key" json:"id"`           //      1            2

	//品牌名称
	Name string `json:"name"`                       //      富士山        南极人
}

//分类
type Classify struct {
	ID int `gorm:"primary_key" json:"id"`      //          1         2            3

	//类别名称
	Name string `json:"name"`                  //          水果       上衣          衣服
} // 1 电脑 衣服 生活用品 水果  海鲜


//以上是SPU相关信息的表的设计；
//以下是SKU相关信息的表的设计

//SKU
type SKU struct {
	ID int `gorm:"primary_key" json:"id"`                             //        1           2

	//商品ID
	GoodID int `json:"good_id"`                  //FK  spu_id                   1苹果        2衬衫

    //店铺ID
 	StoreID int `json:"store_id"`               //FK                            1富士山旗舰店  2南极人旗舰店
}

//SKU-规格值表ID
type SKUStandardvalue struct {
	ID int `gorm:"primary_key" json:"id"`              //                        1             2
	SKUID int `json:"skuid"`                          //FK                       1             2
	StandardValueID int `json:"standard_value_id"`    //FK    规格值表ID           1 5cm         2 圆领
	Price float64 `json:"price"`                                       //        6.5元/斤     54元/件 (计量单位？如何设计)
}

//店铺
type Store struct {
	ID int `gorm:"primary_key" json:"id"`             // 1             2
	Name string `json:"name"`                        //  富士山旗舰店     南极人旗舰店
}

//增值保障  不了解
type VAG struct {
	ID int `gorm:"primary_key" json:"id"`           //1                2
	Name string `json:"name"`                       //
	Content string `json:"content"`                 //
}

//增值保障-SKU
type SKU_vAG struct {
	ID int `gorm:"primary_key" json:"id"`
	VAGID int `json:"vagid"`
	SKUID int `json:"skuid"`
}