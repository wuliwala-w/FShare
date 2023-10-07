package routers

import (
	"FShare/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	//r.Static("/static", "static")
	//r.LoadHTMLGlob("templates/*")
	r.GET("/myfile", controller.IndexHandler)

	//页面一路由
	v1Group := r.Group("myfile")
	{
		//添加

		//查看所有信息
		v1Group.GET("/todo", controller.GetFileList)
		//查看一个文件信息
		v1Group.GET("/onefile/:id", controller.GetFileByID)
		//按钮触发更改状态
		v1Group.PUT("/applied/:id", controller.UpdateStatus)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteApply)
	}

	//页面二路由
	v2Group := r.Group("/upload")
	{
		v2Group.POST("/confirm", controller.UploadFile)
	}

	//页面三路由
	v3Group := r.Group("/browse")
	{
		v3Group.POST("/apply", controller.CreateApply)
	}

	//页面四路由
	v4Group := r.Group("/verify")
	{
		v4Group.GET("/index", controller.IndexHandler)
		v4Group.POST("/upload", controller.UploadFileLocal)
		v4Group.GET("/fingerprint", controller.GetFingerPrint)
		v4Group.GET("/traceback", controller.TraceBackOnChain)
		v4Group.GET("/detail", controller.DetailInformation)
	}

	return r
}
