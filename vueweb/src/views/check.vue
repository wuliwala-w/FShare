<script setup>
    import {ref} from 'vue'
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
        ElUpload,
        ElMessageBox,
        ElInput,
        ElRadioGroup,
        ElRadio,
        ElDescriptions,
        ElDescriptionsItem
    } from 'element-plus';
    const fileList=[]
    const radio=ref('')
    const tableData = [
        {
            date: '2016-05-03',
            name: 'Tom',
            address: 'No. 189, Grove St, Los Angeles',
        },
        {
            date: '2016-05-02',
            name: 'Tom',
            address: 'No. 189, Grove St, Los Angeles',
        },
        {
            date: '2016-05-04',
            name: 'Tom',
            address: 'No. 189, Grove St, Los Angeles',
        },
        {
            date: '2016-05-01',
            name: 'Tom',
            address: 'No. 189, Grove St, Los Angeles',
        },
    ]
    const handleRemove = (file, uploadFiles) => {
        console.log(file, uploadFiles)
    }
    const handlePreview = (uploadFile) => {
        console.log(uploadFile)
    }


    const beforeRemove= (uploadFile, uploadFiles) => {
        return ElMessageBox.confirm(
            `Cancel the transfer of ${uploadFile.name} ?`
        ).then(
            () => true,
            () => false
        )
    }
    const handleAvatarSuccess=(response, uploadFile)=> {
        console.log(response,uploadFile)
        // imageUrl.value = URL.createObjectURL(uploadFile.raw)
    }
    const handleCheck=()=>{
        ElMessage({
            message: '校验成功',
            type: 'success',
        })
    }

</script>

<template>
    <div class="check">
        <el-card>
            <div class="header">
                <div style="display: flex;justify-content: center;">
                    <el-upload
                            v-model:file-list="fileList"
                            class="upload-demo"
                            action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
                            :on-preview="handlePreview"
                            :on-remove="handleRemove"
                            :before-remove="beforeRemove"
                            :limit="1"
                            :on-success="handleAvatarSuccess"
                            style="display: flex;width: 500px;"
                    >
                        <el-button type="primary">上传文件</el-button>
                    </el-upload>
                </div>
                <div style="display: flex;justify-content: center;margin: 16px 0;">
                    <div style="width: 500px;display: flex;">
                        <el-button type="success">提取指纹</el-button>
                        <el-input style="width: 300px;margin-left: 20px"></el-input>
                    </div>
                </div>
                <div style="display: flex;justify-content: center;margin: 16px 0;">
                    <div style="width: 500px;display: flex;">
                        <el-button type="warning">链上追溯</el-button>
                        <el-radio-group v-model="radio" class="ml-4" style="margin-left: 20px;">
                            <el-radio label="A">A</el-radio>
                            <el-radio label="B">B</el-radio>
                            <el-radio label="C">C</el-radio>
                            <el-radio label="D">D</el-radio>
                            <el-radio label="E">E</el-radio>
                            <el-radio label="F">公开</el-radio>
                        </el-radio-group>
                    </div>
                </div>
<!--                <el-button type="success" @click="handleCheck">核验</el-button>-->
            </div>
        </el-card>
        <div class="con">
<!--            <el-card style="width: 49%;min-height: 600px;">-->
<!--                <template #header>-->
<!--                    <div class="card-header">-->
<!--                        <h3 style="text-align: center;">文件信息</h3>-->
<!--                        &lt;!&ndash;                    <el-button class="button" text>Operation button</el-button>&ndash;&gt;-->
<!--                    </div>-->
<!--                </template>-->
<!--                <el-table :data="tableData" style="width: 100%">-->
<!--                    <el-table-column prop="date" label="Date" width="180" />-->
<!--                    <el-table-column prop="name" label="Name" width="180" />-->
<!--                    <el-table-column prop="address" label="Address" />-->
<!--                </el-table>-->
<!--            </el-card>-->
<!--            <el-card style="width: 49%;min-height: 600px;">-->
<!--                <template #header>-->
<!--                    <div class="card-header">-->
<!--                        <h3 style="text-align: center;">区块链信息</h3>-->

<!--                        &lt;!&ndash;                    <el-button class="button" text>Operation button</el-button>&ndash;&gt;-->
<!--                    </div>-->
<!--                </template>-->
<!--                <el-table :data="tableData" style="width: 100%">-->
<!--                    <el-table-column prop="date" label="Date" width="180" />-->
<!--                    <el-table-column prop="name" label="Name" width="180" />-->
<!--                    <el-table-column prop="address" label="Address" />-->
<!--                </el-table>-->
<!--            </el-card>-->
            <el-card style="width: 100%;padding: 0 25%;">
                <el-descriptions :column="2" size="large" class="mt-4">
                    <el-descriptions-item label="发送方"
                    >kooriookami</el-descriptions-item
                    >
                    <el-descriptions-item label="时间"
                    >2022-02-21</el-descriptions-item
                    >
                    <el-descriptions-item label="接收方">kooriookami</el-descriptions-item>
                    <el-descriptions-item label="共享类型">
                        测试
                    </el-descriptions-item>
                </el-descriptions>
            </el-card>
        </div>
    </div>
</template>



<style lang="scss">
    .check{
        height: 100%;
        width: 100%;
        .header{
            /*display: flex;*/
            /*justify-content: center;*/
        }
        .el-upload-list{
            width: 300px !important;
        }
        .con{
            width: 100%;
            margin-top: 20px;
            display: flex;
            justify-content: space-between;
        }
        .el-descriptions__label{
            font-size: 20px !important;
            font-weight: bold;
        }
    }
</style>
