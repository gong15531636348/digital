package controller

import (
	"fmt"
	"golang.org/x/net/context"
	"ihome555/service/getArea/message"
	"ihome555/service/getArea/model"
	"ihome555/utils"
)

type GetAreaSer struct {
}

func (g *GetAreaSer) GetArea(ctx context.Context, req *message.Requset) (*message.Response, error) {
	//查询数据库,把对应的数据赋值给Response
	fmt.Println("--------------------")
	var resp *message.Response = new(message.Response)
	areas, err := model.GetAreaModel()
	if err != nil {
		fmt.Println("model.GetAreaModel err:", err)
		resp.Errno = utils.RECODE_DBERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return resp, err
	}
	fmt.Println("--------------------")
	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	for _, v := range areas {
		var areaInfo message.Area
		areaInfo.Aid = int32(v.Id)
		areaInfo.Aname = v.Name
		resp.Data = append(resp.Data, &areaInfo)
	}
	fmt.Println("--------------------")
	fmt.Println(resp)
	return resp, nil

}
