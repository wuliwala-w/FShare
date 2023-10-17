# FShare项目
web后端采用go

#### 每个服务器部署之前，在models里更改Node名称 

##### postman调试：
初始化查看我的页面 
GET
127.0.0.1:8080/myfile 

上传文件upload  
POST  
127.0.0.1:8080/upload/confirm   
表单类型第一条是文件，后面是json数据  
文件key:f1，文件实例（图片，文档）  
{"id":"random","name":"file1","fileOwner":"A","description":"this is file1","size":"1kb","time":"random","status":0}

页面三获取所有文件browse  
GET  
127.0.0.1:8080/browse

创建申请createApply  
POST  
127.0.0.1:8080/browse/apply    
{"id":"A16966616127948220009928","name":"file1","fileOwner":"A","description":"this is file1","size":"1kb","time":"B","status":0}

根据id查询单个文件(最后是文件具体id)  
GET
127.0.0.1:8080/myfile/onefile/文件id

更新文件状态  
PUT
127.0.0.1:8080/myfile/applied/文件id  
{"status":3}

更新申请状态  
PUT
127.0.0.1:8080/myfile/update/文件id/申请节点(ABCDE) 
{"status":3}

删除文件  
DELETE
127.0.0.1:8080/myfile/deletefile/文件id

删除申请 
DELETE
127.0.0.1:8080/myfile/deleteapply/文件id/申请节点

上传核验文件 
POST
127.0.0.1:8080/verify/upload2    
文件key:FileVerify，文件实例（文本文件、图片）

提取文件水印  
GET  
127.0.0.1:8080/verify/fingerprint/核验文件的类型

