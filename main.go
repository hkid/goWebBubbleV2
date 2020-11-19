package main

import (
	"awesomeProject2/models"
	"awesomeProject2/dao"
	"awesomeProject2/routers"

)


func main() {
	// 0 创建数据库
	// 1 连接数据库 ->封装为函数
	err:=dao.InitMysql()
	if err!=nil{
		panic(err)
	}
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	r:=routers.SetupRouter()
	r.Run(":80")


}
