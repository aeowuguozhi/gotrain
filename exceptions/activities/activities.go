package activities


import "my_demo/models"

func ActivitiesIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "活动不存在",
	}
}

func ActivitiesTagIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "活动标签不存在",
	}
}
