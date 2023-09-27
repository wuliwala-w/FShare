package routers

import (
	"FShare/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/myfile", controller.IndexHandler)

	v1Group := r.Group("myfile")
	{
		//添加
		v1Group.POST("/upload", controller.UploadFile)
		//查看所有待办
		v1Group.GET("/todo", controller.GetFileList)
		////查看一个待办
		//v1Group.GET("/todo/:id", func(context *gin.Context) {
		//
		//})
		//修改待办
		v1Group.PUT("/applied/:id", controller.UpdateTask)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteApply)
	}
	v2Group := r.Group("/upload")
	{
		v2Group.POST("/confirm", controller.UploadFile)
	}
	v3Group := r.Group("/browse")
	{
		v3Group.POST("/apply", controller.CreateApply)
	}

	return r
}
