package models

import (
	"FShare/dao"
)

type File struct {
	FileOwner string `json:"fileOwner"`
	FileID    int    `json:"id"`
	Name      string `json:"name"`
	Status    int    `json:"status"` //1:没被申请；2：正在被申请中；3：申请被拒绝；4：可用不可转发；5：可用可转发
}

type Apply struct {
	ApplyOwner string `json:"applyOwner"`
	FileOwner  string `json:"fileOwner"`
	Timestamp  string `json:"timestamp"`
	FileID     int    `json:"fileID"`
	Status     int    `json:"status"`
}

/*
	Todo这个model的增删改查放在这里
*/

func UploadFiles(file *File) (err error) {
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
