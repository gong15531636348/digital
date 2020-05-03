package controller

import (

	"github.com/gin-gonic/gin"
	"ihome555/model"
	"ihome555/utils"
	"net/http"
)

func GetArea(ctx *gin.Context)  {
	 resp:=make(map[string]interface{})
     areas,err:=model.GetAreas()
	if err != nil {
		resp["errno"]=utils.RECODE_DBERR
		resp["errmsg"]=utils.RecodeText(utils.RECODE_DBERR)
		ctx.JSON(http.StatusOK,resp)
	}
	resp["errno"]=utils.RECODE_OK
	resp["errmsg"]=utils.RecodeText(utils.RECODE_OK)
	resp["data"]=areas
	ctx.JSON(http.StatusOK,resp)
}