<script setup>
import { ref } from 'vue'
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
    ElMessageBox,
    ElInput,
    ElDatePicker,
    ElRadioGroup,
    ElRadio
} from 'element-plus';

import axios from 'axios'
const dialogVisible = ref(false)
const value1 = ref('')
const radio = ref(0)
const tableData = ref([ 
])
const fileId = ref('')
const fileName = ref('')
const dataNum = ref('')
const history = ref('')
const arr=ref([])
const handleClose = (done) => {
    dialogVisible.value = false
    // ElMessageBox.confirm('Are you sure to close this dialog?')
    //     .then(() => {
    //         done()
    //     })
    //     .catch(() => {
    //         // catch error
    //     })
}
const showBox = (row) => {

    console.log(row)
    // arr.value=row,
    fileId.value=row.id,
    fileName.value=row.name,
    // radio=row.radio,
    // history=row.history,
    dialogVisible.value = true
}


const getBrowseList = () => {
    //后端get数据
    axios.get("browse").then(res => {
        console.log(res);
        tableData.value = res.data
    })
}   

const handleSubmit = () => {
    axios.post("browse/apply", {
       
    "id": fileId.value, 
    "name": fileName.value, 
    "fileOwner": arr.value.fileOwner, 
    "description": arr.value.description,
     "size": arr.value.size, 
     "time": value1.value, 
     "status": radio.value
    }).then(res => {
        console.log(res);
        ElMessage({
    message: '申请成功',
    type: 'success',
  })
    })
    dialogVisible.value = false
}

const handleSubmit2 = (row) => {
    //申请共享
    axios.post("browse/apply", {      
    "id": row.id, 
    "name": row.name, 
    "fileOwner": row.fileOwner, 
    "description": row.description,
     "size": row.size, 
     "time": row.time, 
     "status": row.status
    }).then(res => {
        console.log(res);
        
        getBrowseList()
        
        ElMessage({
    message: '申请成功',
    type: 'success',
  })
    })
    
}

getBrowseList()
</script>




<template>
    <div class="browse">
        <el-card class="box-card">
            <div>
                <el-table :data="tableData" style="width: 100%" border>
                    <el-table-column prop="fileOwner" label="组织名称" />
                    <el-table-column prop="name" label="文件名" />
                    <el-table-column prop="id" label="ID" />
                    <el-table-column prop="size" label="大小" />
                    <el-table-column prop="description" label="描述" width="300" />
                    <el-table-column label="操作">
                        <template v-slot="scope">
                            <el-button style="margin-left: 40px" type="primary" @click="showBox(scope.row)">查看</el-button>
                            <el-button  style="margin-left: 40px" type="primary" @click="handleSubmit2(scope.row)" v-if="scope.row.status!=1">申请共享</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
        </el-card>


        <!--对话框-->
        <el-dialog v-model="dialogVisible" title="查看" width="30%" :before-close="handleClose">
            <div>
                <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">
                    <div style="width: 18%;">文件ID：</div>
                    <div style="width: 82%;">
                        <el-input v-model="fileId"></el-input>
                    </div>
                </div>
                <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">
                    <div style="width: 18%;">文件名：</div>
                    <div style="width: 82%;">
                        <el-input v-model="fileName"></el-input>
                    </div>
                </div>
                <!--                <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">-->
                <!--                    <div style="width: 18%;">数据条目：</div>-->
                <!--                    <div style="width: 82%;">-->
                <!--                        <el-input v-model="dataNum"></el-input>-->
                <!--                    </div>-->
                <!--                </div>-->
                <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">
                    <div style="width: 18%;">上传时间：</div>
                    <div style="width: 82%;">
                        <el-date-picker value-format="YYYY-MM-DD" v-model="value1" type="date" placeholder="Pick a day" :size="size" />
                    </div>
                </div>
                <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">
                    <div style="width: 18%;">可转发：</div>
                    <div style="width: 82%;">
                        <el-radio-group v-model="radio" class="ml-4">
                            <el-radio :label="1">是</el-radio>
                            <el-radio :label="0">否</el-radio>
                        </el-radio-group>
                    </div> 
                </div>
                <!-- <div style="display: flex;width: 90%;align-items: center;margin-bottom: 16px;">
                    <div style="width: 18%;">转发历史：</div>
                    <div style="width: 82%;">
                        <el-input v-model="history"></el-input>
                    </div>
                </div> -->
            </div>
            <template #footer>
                <span class="dialog-footer">
                    <el-button type="primary" @click="handleSubmit()" v-if="radio!=0">
                        申请
                    </el-button>
                    

                </span>
            </template>
        </el-dialog>
    </div>
</template>



<style lang="scss" scoped>
.browse {
    height: 100%;
    width: 100%;
}

.dialog-footer button:first-child {
    margin-right: 10px;
}
</style>
