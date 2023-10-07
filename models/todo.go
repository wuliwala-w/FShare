package models

import (
	"FShare/dao"
	"fmt"
	"math/rand"
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
	FileID     string `json:"fileID"`
	Status     int    `json:"status"`
}

var Node string = "A" //节点

/*
	Todo这个model的增删改查放在这里
*/

func UploadFiles(file *File) (err error) {
	t := time.Now()
	file.Time = t.Format("2006-01-02 15:04:05")
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	randnum := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	file.FileID = Node + timestamp + randnum
	if err = dao.DB.Create(&file).Error; err != nil {
		return err
	}
	return
}

func CreateAplly(apply *Apply) (err error) {
	if err = dao.DB.Create(&apply).Error; err != nil {
		return err
	}
	return
}

func GetFileList() (filelist []*File, err error) {
	if err = dao.DB.Find(&filelist).Error; err != nil {
		return nil, err
	}
	return
}

func GetMyApply(node string) (applylist []*Apply, err error) {
	if err = dao.DB.Where("apply_owner <> ?", node).Find(&applylist).Error; err != nil {
		return nil, err
	}
	return
}

// 连接apply数据库表
func GetApplyList(node string) (applylist []*Apply, err error) {
	if err = dao.DB.Where("file_owner <> ?", node).Find(&applylist).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodoByID(id string) (todo *File, err error) {
	todo = new(File)
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *File) (err error) {
	err = dao.DB.Save(todo).Error
	return err

}

func DeleteATodoByID(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&File{}).Error
	return
}
