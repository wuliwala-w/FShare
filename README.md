# FShare项目
web后端采用go

#### 每个服务器部署之前，在models里更改Node名称

##### postman调试：
初始化查看我的页面
127.0.0.1:8080/myfile

上传文件upload
127.0.0.1:8080/upload/confirm
文件key:f1
{"id":"","name":"file1","fileOwner":"A","description":"this is file1","size":"1kb","time":"B","status":0}

创建申请createApply
127.0.0.1:8080/browse/apply
{"id":"A16966616127948220009928","name":"file1","fileOwner":"A","description":"this is file1","size":"1kb","time":"B","status":0}

根据id查询单个文件(最后是文件id)
127.0.0.1:8080/myfile/onefile/文件id

更新文件状态
127.0.0.1:8080/myfile/applied/文件id
{"status":3}

更新申请状态
127.0.0.1:8080/myfile/update/文件id/申请节点
{"status":3}

删除文件
127.0.0.1:8080/myfile/deletefile/文件id

删除申请
127.0.0.1:8080/myfile/deleteapply/文件id/申请节点