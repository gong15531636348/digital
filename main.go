package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ihome555/controller"
	"ihome555/model"
)



func main()  {
	router:=gin.Default()
	err:=model.InitDb()
	if err != nil {
		fmt.Println(err)
		return
	}


	router.Static("/home","view")
	r1:=router.Group("/api/v1.0")
	{
		r1.GET("/areas", controller.GetArea)

	}
	router.Run(":8066")
}