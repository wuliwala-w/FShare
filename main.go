package main

import (
	"FShare/dao"
	"FShare/models"
	"FShare/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.File{}, &models.Apply{})
	r := routers.SetupRouter()
	r.Run()
}
