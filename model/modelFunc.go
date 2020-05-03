package model

import "fmt"

func GetAreas()([]Area,error) {
	var areas []Area
	if err:=GlobalDB.Find(&areas).Error;err!=nil{
		fmt.Println("获取房源信息错误",err)
		return nil,err
	}
	return areas,nil
}
