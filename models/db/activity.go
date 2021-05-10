package db

//活动表
//Int64 就是时间戳
type Activities struct {
	Id int `gorm:"primary_key" json:"id"`

	Title string `json:"title" gorm:"not null"`

	Content string `json:"content" gorm:"not null"`

	StartTime int64 `json:"start_time" gorm:"not null"`

	EndTime int64 `json:"end_time"`

	IsPublic bool `json:"is_public" gorm:"not null"`

	SuitPeople string  `json:"suit_people"`

	Cost float64 `json:"cost" gorm:"not null"`

}

//类别表
type ActivitiesTag struct {
	Id int ` json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

//活动表和标签表的映射关系
type ActivitiesAndTag struct {
	Id int `gorm:"primary_key" json:"id"`

	ActivitiesId int `json:"activities_id" gorm:"not null"`

	ActivitiesTagId int `json:"activities_tag_id" gorm:"not null"`
}
