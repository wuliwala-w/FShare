package models

import (
	"FShare/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"time"
)

type File struct {
	FileID      string `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	FileOwner   string `json:"fileOwner"`
	Description string `json:"description"`
	Size        string `json:"size"`
	Time        string `json:"time"`
	Status      int    `json:"status"` //1:没被申请；2：正在被申请中；3：申请被拒绝；4：可用不可转发；5：可用可转发
}

type Apply struct {
	ApplyOwner string `json:"applyOwner"`
	FileOwner  string `json:"fileOwner"`
	Time       string `json:"time"`
	FileID     string `json:"id"`
	Status     int    `json:"status"`
}

var Node string = "A" //节点

/*
	Todo这个model的增删改查放在这里
*/

func UploadFiles(context *gin.Context) (err error) {
	var file File

	file.FileOwner = context.PostForm("fileOwner")
	file.Name = context.PostForm("name")
	file.Description = context.PostForm("description")
	file.Size = context.PostForm("size")
	file.Status, _ = strconv.Atoi(context.PostForm("status"))

	f, err := context.FormFile("f1")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		//保存读取的文件到本地服务器
		dst := path.Join("./", f.Filename)
		_ = context.SaveUploadedFile(f, dst)
		context.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
		//生成文件ID
		t := time.Now()
		file.Time = t.Format("2006-01-02 15:04:05")
		timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
		randnum := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
		file.FileID = Node + timestamp + randnum
		file.Status = 1
		if err = dao.DB.Create(&file).Error; err != nil {
			return err
		}
	}
	return
}

func CreateApply(file *File) (err error) {
	var apply Apply
	apply.ApplyOwner = Node
	apply.FileOwner = file.FileOwner
	apply.FileID = file.FileID
	apply.Status = 2
	t := time.Now()
	apply.Time = t.Format("2006-01-02 15:04:05")
	if err = dao.DB.Create(&apply).Error; err != nil {
		return err
	}
	return
}

func GetFileList() (fileList []*File, err error) {
	if err = dao.DB.Where("file_owner = ?", Node).Find(&fileList).Error; err != nil {
		return nil, err
	}
	return
}

func GetMyApply() (applyList []*Apply, err error) {
	if err = dao.DB.Where("apply_owner = ?", Node).Find(&applyList).Error; err != nil {
		return nil, err
	}
	return
}

// 连接apply数据库表
func GetApplyList() (applyList []*Apply, err error) {
	if err = dao.DB.Where("file_owner = ?", Node).Find(&applyList).Error; err != nil {
		return nil, err
	}
	return
}

func GetFileByID(id string) (file *File, err error) {
	file = new(File)
	if err = dao.DB.Where("file_id = ?", id).First(&file).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateFile(file *File) (err error) {
	err = dao.DB.Save(file).Error
	return err

}

func DeleteATodoByID(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&File{}).Error
	return
}

func SaveFilelocal(context *gin.Context) (err error) {
	//将上传的文件取出来
	f, err := context.FormFile("FileVerify")
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": err.Error()})
	} else { //将上传的文件保存到指定的本地地址，并返回响应
		log.Println(f.Filename)
		dst := fmt.Sprintf("./verifyfile/%s", f.Filename) //设置核验文件保存的本地地址路径z
		if err = context.SaveUploadedFile(f, dst); err != nil {
			return err
		}
	}
	return
}
