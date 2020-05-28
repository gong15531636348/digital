package controller

import (
	"fmt"
	"ihome555/model"
	"ihome555/service/getArea/message"
	"ihome555/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func GetArea(ctx *gin.Context) {
	//resp:=make(map[string]interface{})
	//areas,err:=model.GetAreas()
	//if err != nil {
	//	resp["errno"]=utils.RECODE_DBERR
	//	resp["errmsg"]=utils.RecodeText(utils.RECODE_DBERR)
	//	ctx.JSON(http.StatusOK,resp)
	//}
	//resp["errno"]=utils.RECODE_OK
	//resp["errmsg"]=utils.RecodeText(utils.RECODE_OK)
	//resp["data"]=areas
	//ctx.JSON(http.StatusOK,resp)
	//使用grpc完成业务
	fmt.Println("--------------------")
	conn, err := grpc.Dial("localhost:8012", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	fmt.Println("--------------------")

	SerClient := message.NewGetAreaSerClient(conn)
	orderRequset := message.Requset{}
	orderInfo, _ := SerClient.GetArea(context.Background(), &orderRequset)
	fmt.Println("-------------------", orderInfo)
	ctx.JSON(http.StatusOK, orderInfo)
}

func GetSession(ctx *gin.Context) {
	resp := make(map[string]interface{})
	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	ctx.JSON(http.StatusOK, resp)
}

//{
//"app_id":"100"
//"tt_enable":"0"
//"email":"1"
//"sms_enable":"1"
//"email_enable":"0"
//"call_enable":"0"
//"psa_role": [
//{
//"role_id": "1"
//},
//{
//"role_id":"2"
//}
//]
//"receive_users":[
//{
//"user_id":"1"
//},
//{
//"user_id":"2"
//}
//]
//"type": "0"
//}
//type AlertConfig struct {
//	APP_id        string               `json:"app_id"`
//	TT_enable     bool                 `json:"tt_enable"`
//	Email_enable  bool                 `json:"email_enable"`
//	Sms_enable    bool                 `json:"sms_enable"`
//	Call_enable   bool                 `json:"call_enable"`
//	Psa_role      []*model.PsaRole     `json:"psa_role"`
//	Receive_users []*model.ReceiveUser `json:"receive_users"`
//	Type          int                  `json:"type"`
//}

func CreatAlertConfig(ctx *gin.Context) {
	resp := make(map[string]interface{})
	alertconfig := model.AlertConfig{}
	err := ctx.ShouldBindJSON(&alertconfig)
	if err != nil {
		fmt.Println("获取数据失败")
		resp["errmsg"] = "获取数据失败"
		ctx.JSON(http.StatusOK, resp)
		return
	}
	//把数据存入数据库
}
