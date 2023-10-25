<script setup>
import { ref, reactive, onMounted } from 'vue'
import axios from 'axios';
import {
    ElAside,
    ElContainer,
    ElHeader,
    ElMain,
    ElMenu,
    ElSubMenu,
    ElMenuItemGroup,
    ElMenuItem,
    ElRow,
    ElCard,
    ElMessage,
    ElButton,
    ElTable,
    ElTableColumn,
    ElDialog,
    ElInput
} from 'element-plus';


const tableData = ref([])
const tableData2 = ref([])
const tableData3 = ref([])
const initValue=ref(0)
const list = []
const dialogVisible1 = ref(false)
const dialogVisible2 = ref(false)
const dialogVisible3 = ref(false)
const fileVisible= ref(false)




const checkFile = reactive({
    description: "",
    fileOwner: "",
    time: "",
    id: "",
    fileName: "",
    hash:"",
    size: "",
    status:"",
    fingerprint: "",
    applyOwner:""
})

const table1 = reactive({
    applyOwner: "",
    fileOwner: "",
    time: "",
    id: "",
    fileName: "",
    hash:"",
})

const table2 = reactive({
    date: "",
    name: "",
})

const table3 = reactive({
    date: "",
    name: "",
    id: "",
    des: "",
    hash: "",
})
const showTable1 = (item) => {

    table1.applyOwner = item.applyOwner
    table1.fileOwner = item.fileOwner
    table1.time = item.time
    table1.id = item.id
    table1.fileName = item.fileName
    dialogVisible1.value = true

}
const showTable2 = (item) => {
    checkFile.fileOwner=item.fileOwner
    checkFile.fileName=item.fileName
    checkFile.applyOwner=item.applyOwner
    checkFile.id=item.id
    checkFile.status=item.status
    checkFile.time=item.time
    checkFile.txHash=item.txHash
    fileVisible.value = true

}
const showTable3 = (item) => {

    table3.time = item.time
    table3.name = item.name
    table3.id = item.id
    table3.description = item.description
    dialogVisible3.value = true

}





//增加了根据id向后端查询文件信息的函数
const getFileById = (row) => {
    console.log(row);
    axios.get(`/myfile/onefile/${row.id}`).then(res => {
        checkFile.fileOwner= res.data.fileOwner
        checkFile.fileName = res.data.name
        checkFile.id = res.data.id
        checkFile.hash=res.data.hash
        checkFile.time=res.data.time
        checkFile.description=res.data.description
        checkFile.size=res.data.size
        checkFile.status=res.data.status
        checkFile.hash=res.data.hash
        fileVisible.value = true
        if(res.data.status==1){
            checkFile.status="是"
        }else{
            checkFile.status="否"
        }
        
        console.log(res);
    })
}

const getData = () => {
    axios.get("/myfile").then(res => {
        tableData.value = res.data.myapply
        tableData2.value = res.data.applylist
        tableData3.value = res.data.filelist
            tableData2.value=tableData2.value.map(item=>{
            item.isOper=true
            return item
        console.log(tableData2.value);
        initValue.value=1
        })
    })
    
}

const handleDelete = (row) => {
    //删除
    //console.log("   kklll")
    axios.delete(`/myfile/deleteapply/${row.id}/${row.applyOwner}`).then(res => {
        console.log(res);
        getData()
    })
}

const updateapply = (row) => {
    //更新申请状态--拒绝
    console.log(row);
    axios.put(`/myfile/update/${row.id}/${row.applyOwner}`, { status: 3 }).then(res => {
        console.log(res);
        getData()
    })
    tableData2.value=tableData2.value.map((item)=>{
        if(row.id===item.id){
            item.isOper=false
        }
        return item
    })
}

const updateapply2 = (row) => {
    //更新申请状态--可用不可转发
    console.log(row);
    axios.put(`/myfile/update/${row.id}/${row.applyOwner}`, { status:4}).then(res => {
        console.log(res);
        getData()
    })
    tableData2.value=tableData2.value.map((item)=>{
        if(row.id===item.id){
            item.isOper=false
        }
        return item
    })
}

const updateapply3 = (row) => {
    //更新申请状态--可用可转发
    console.log(row);
    axios.put(`/myfile/update/${row.id}/${row.applyOwner}`, { status: 5}).then(res => {
        console.log(res);
        getData()
    })
    tableData2.value=tableData2.value.map((item)=>{
        if(row.id===item.id){
            item.isOper=false
        }
        return item
    })
}
const download=(row)=>{
    console.log(row);

    const filename = row.fileName
    const fileowner = row.fileOwner
    //下载
    // axios.get(`/myfile/download/${row.fileName}/${row.fileOwner}`).then(res=>{

    //     console.log(res);

    // })
    axios({
        url: `myfile/download/${filename}/${fileowner}`, 
        method: 'GET',
        responseType: 'blob', // 设置响应类型为二进制数据流
      })
        .then((response) => {
          const blob = new Blob([response.data], {
            type: response.headers['content-type'], // 根据响应头获取文件类型
          });
          const url = window.URL.createObjectURL(blob);
          const a = document.createElement('a');
          a.href = url;
          a.download = filename; // 替换成你想要的文件名和扩展名
          a.style.display = 'none';
          document.body.appendChild(a);
          a.click();
          window.URL.revokeObjectURL(url);
        })
        .catch((error) => {
          console.error('下载文件失败:', error);
        });

}
onMounted(async ()=>{
    await getData()
   
})
</script>

<template>
    <div class="mine">
        <el-dialog v-model="fileVisible" width="30%">
            <!--所有的查看详细信息都用这一个页面-->
            <div>
                <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">
                    <div style="width: 18%;">文件拥有方：</div>
                    <div style="width: 82%;">
                        <el-input v-model="checkFile.fileOwner" readonly></el-input>
                    </div>
                </div>
                <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">
                    <div style="width: 18%;">文件名称：</div>
                    <div style="width: 82%;">
                        <el-input v-model="checkFile.fileName" readonly></el-input>
                    </div>
                </div>
                <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">
                    <div style="width: 18%;">上传时间：</div>
                    <div style="width: 82%;">
                        <el-input v-model="checkFile.time" readonly></el-input>
                    </div>
                </div>
                <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">
                    <div style="width: 18%;">文件ID:</div>
                    <div style="width: 82%;">
                        <el-input v-model="checkFile.id" readonly></el-input>
                    </div>
                </div>
                <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">
                    <div style="width: 18%;">文件大小:</div>
                    <div style="width: 82%;">
                        <el-input v-model="checkFile.size" readonly></el-input>
                    </div>
                </div>
                <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">
                    <div style="width: 18%;">是否可转发:</div>
                    <div style="width: 82%;">
                        <el-input v-model="checkFile.status" readonly></el-input>
                    </div>
                </div>
                <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">
                    <div style="width: 18%;">区块链哈希:</div>
                    <div style="width: 82%;">
                        <el-input v-model="checkFile.hash" type="textarea" resize="none" :autosize="{ minRows: 2, maxRows: 2 }" readonly></el-input>
                    </div>
                </div>
                <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">
                    <div style="width: 18%;">文件描述:</div>
                    <div style="width: 82%;">
                        <el-input v-model="checkFile.description" readonly></el-input>
                    </div>
                </div>
            </div>
        </el-dialog>

      
        <el-card class="box-card" style="margin-top: 20px">
            <template #header>
                <div class="card-header">
                    <h3 style="text-align: center;">我的文件</h3>
                    <!--                    <el-button class="button" text>Operation button</el-button>-->
                </div>
            </template>
            <div>
                <el-table :data="tableData3" style="width: 100%" border>
                    <el-table-column prop="id" label="ID" />
                    <el-table-column prop="name" label="文件名称" />
                    <el-table-column prop="description" label="文件描述" />
                    <el-table-column prop="time" label="时间" />
                    <el-table-column label="操作">
                        <template v-slot="scope">
                            <div style="display: flex;">
                                <el-button style="margin-left: 40px" type="primary"
                                    @click="getFileById(scope.row)">查看</el-button>
                                <el-button style="margin-left: 40px" type="danger">删除</el-button>
                            </div>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
        </el-card>
        <el-card class="box-card" style="margin-top: 20px">
            <template #header>
                <div class="card-header">
                    <h3 style="text-align: center;">我的申请</h3>
                    <!--                    <el-button class="button" text>Operation button</el-button>-->
                </div>
            </template>
            <div>
                <el-table :data="tableData" style="width: 100%" border>
                    <el-table-column prop="fileOwner" label="组织" width="180" />
                    <el-table-column prop="fileName" label="文件名称" width="180" />
                    <el-table-column prop="time" label="申请时间" width="180" />
                    <el-table-column prop="id" label="文件ID" width="180" />
                    <el-table-column label="状态" width="180">
                        <template #default="scope">
                            {{
                                scope.row.status == 1 ? '没被申请' : scope.row.status == 2 ? '正在被申请中' : scope.row.status == 3 ? '申请被拒绝' : scope.row.status == 4 ? '可用不可转发' :
                                    scope.row.status == 5 ? '可用可转发':""}}
                        </template>
                    </el-table-column>

                    <el-table-column label="">
                        <template v-slot="scope">
                            <div style="display: flex;">
                                <div v-html="scope.row.address"></div>
                                <el-button style="margin-left: 40px" v-if="scope.row.status == 4"
                                    type="success" @click="download(scope.row)">下载</el-button>
                               
                                <el-button style="margin-left: 40px" type="primary"
                                    @click="getFileById(scope.row)">查看</el-button>
                            </div>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
        </el-card>
        <el-card class="box-card" style="margin-top: 20px">
            <template #header>
                <div class="card-header">
                    <h3 style="text-align: center;">申请信息</h3>
                    <!--                    <el-button class="button" text>Operation button</el-button>-->
                </div>
            </template>
            <div>
                <el-table :data="tableData2" style="width: 100%" border>
                    <el-table-column prop="applyOwner" label="申请节点" width="180" />
                    <el-table-column prop="fileName" label="文件名称" width="180" />

                    <el-table-column label="">
                        <template v-slot="scope">
                            <div style="display: flex;">
                                <el-button style="margin-left: 40px" type="primary"
                                    @click="showTable2(scope.row)">查看</el-button>
                                <el-button style="margin-left: 40px" type="danger"
                                    @click="updateapply(scope.row)" v-if="scope.row.isOper==true">拒绝</el-button>
                                <el-button style="margin-left: 40px" type="warning"
                                    @click="updateapply2(scope.row)" v-if="scope.row.isOper==true">授权亦可用</el-button>
                                <el-button style="margin-left: 40px" type="primary"
                                    @click="updateapply3(scope.row)" v-if="scope.row.isOper==true">可用可转发</el-button>
                                <el-button style="margin-left: 40px" type="danger"
                                    @click="handleDelete(scope.row)">删除</el-button>
                                <el-text type="success" disabled="true" v-if="!scope.row.isOper">已操作</el-text>
                               
                            </div>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
        </el-card>
       
    </div>
</template>



<style lang="scss" scoped>
.mine {
    height: 100%;
    width: 100%;
}
</style>
