### FShare项目
web后端采用go

#### 每个服务器部署之前，在models里更改Node名称

##### postman调试：
初始化查看我的页面
127.0.0.1:8080/myfile

上传文件upload
127.0.0.1:8080/upload/confirm
{"id":"","name":"file1","fileOwner":"A","description":"this is file1","size":"1kb","time":"B","status":0}

创建申请createApply
127.0.0.1:8080/browse/apply
{"id":"","name":"file1","fileOwner":"A","description":"this is file1","size":"1kb","time":"B","status":0}

根据id查询单个文件
127.0.0.1:8080/myfile/onefile/A16966616127948220009928