package main

import (
	"fmt"
	"google.golang.org/grpc"
	"ihome555/model"
	"ihome555/service/getArea/controller"
	"ihome555/service/getArea/message"
	"net"
)

func main() {
	//注册数据库 略
	err:=model.InitDb()
	if err != nil {
		fmt.Println("数据库初始化失败")
		return
	}
	server:=grpc.NewServer()
	message.RegisterGetAreaSerServer(server,new(controller.GetAreaSer))
	conn,err:=net.Listen("tcp",":8012")
	if err != nil {
		fmt.Println("net.Listen err:",err)
		return
	}
	err=server.Serve(conn)
	if err != nil {
		fmt.Println("server.Serve err",err)
	}
}
