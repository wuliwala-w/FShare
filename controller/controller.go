package controller

import (
	"FShare/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/*
url     --> controller --> logic   --> model
请求来了 --> 控制器       --> 业务逻辑 --> 模型层的增删改查
*/

func IndexHandler(context *gin.Context) {
	//context.HTML(http.StatusOK, "index.html", nil)
	//这里更改节点名称，服务器分别为ABCDE，用models.Node
	myApply, err := models.GetMyApply() //获取我的申请
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, myApply)
	}
	applyList, err := models.GetApplyList() //获取别人申请我的信息
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, applyList)
	}
	fileList, err := models.GetFileList()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, fileList)
	}
}

func UploadFile(context *gin.Context) {
	var file models.File
	context.BindJSON(&file)
	if err := models.UploadFiles(&file, context); err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, file)
	}
}

func CreateApply(context *gin.Context) {
	//1.从请求中把数据取出来
	//var apply models.Apply
	var file models.File
	context.BindJSON(&file)
	//2.存入数据库
	//3.返回响应
	if err := models.CreateApply(&file); err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, file)
	}
}

func GetFileList(context *gin.Context) {
	fileList, err := models.GetFileList()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, fileList)
	}
}

func GetFileByID(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"error": "id not exist"})
		return
	}
	//查询数据库是否有这个id
	file, err := models.GetFileByID(id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, file)
	}
}

func UpdateTask(context *gin.Context) {
	//拿到请求里的id
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"error": "id not exist"})
		return
	}
	//查询数据库是否有这个id
	todo, err := models.GetFileByID(id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	//放入todo变量
	context.BindJSON(&todo)
	//新信息保存到数据库
	err = models.UpdateATodo(todo)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

func DeleteApply(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"error": "id not exist"})
		return
	}
	//查询数据库是否有这个id并删除
	if err := models.DeleteATodoByID(id); err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{id: "deleted"})
	}

}

// 将上传的文件保存到本地
func UploadFileLocal(context *gin.Context) {
	//1.将上传的文件取出来
	f, err := context.FormFile("FileVerify")
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	//2.将文件保存到本地缓存区
	log.Println(f.Filename)
	dst := fmt.Sprintf("D/Reaserch/System development/FShare/Verify/%s", f.Filename) //设置核验文件保存的本地地址路径
	//将上传的文件保存到指定的本地地址，并返回响应
	if err = context.SaveUploadedFile(f, dst); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status": "Upload success",
		})
	}
}

// 提取文件水印
func GetFingerPrint(context *gin.Context) {

}

// 提取区块链上文件的信息
func TraceBackOnChain(context *gin.Context) {

}

// 展示详细的信息
func DetailInformation(context *gin.Context) {

}
