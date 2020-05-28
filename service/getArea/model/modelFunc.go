package model

import (
	"fmt"
	"ihome555/model"
)

func GetAreaModel() ([]model.Area, error) {
	//使用gorm查询获取地域信息
	var Areas []model.Area
	if err := model.GlobalDB.Find(&Areas).Error; err != nil {
		fmt.Println("获取地域信息错误", err)
		return nil, err
	}
	return Areas, nil
}
