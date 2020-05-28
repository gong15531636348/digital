package main

import (
	"fmt"
	"ihome555/controller"
	"ihome555/model"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	err := model.InitDb()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Static("/home", "view")
	r1 := router.Group("/api/v1.0")
	{
		r1.GET("/areas", controller.GetArea)
		r1.GET("/session", controller.GetSession)
		r1.POST("/alertconfig", controller.CreatAlertConfig)
		//r1.GET("/alertconfig",controller.GetAlertConfig)

	}
	router.Run(":8071")
}
