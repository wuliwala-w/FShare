<script setup>
    import {ref,reactive} from 'vue'
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
        ElForm,
        ElFormItem,
        ElInput,
        ElRadioGroup,
        ElRadio,
        ElUpload
    } from 'element-plus';
    import axios from 'axios'
    const form = reactive({
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: '',
        num:'',
        description:''
    })
    const fileList=[]
    const onSubmit = () => {
        console.log('submit!')
    }
    const handleAvatarSuccess=(response, uploadFile)=> {
        console.log(response,uploadFile)
        form.name=uploadFile.name
        form.size=uploadFile.size
        // imageUrl.value = URL.createObjectURL(uploadFile.raw)
    }
    const handleRemove = (file, uploadFiles) => {
        console.log(file, uploadFiles)
        form.name=''
        form.size=''
        form.num=''
        form.resource=''
    }
    const beforeUpload=(file,id)=>{
        console.log(file,id);
        axios.post("upload/confirm",{
            "f1":file,
             "name":form.name ||file.name,
             "description":form.description,
             "size":form.size ||file.size,
             "status":form.resource
        },{
            headers: {'Content-Type': 'multipart/form-data'}
        }).then(res=>{
            ElMessage({
    message: '上传成功',
    type: 'success',
  })
        }).catch(err=>{
            ElMessage({
    message: '上传失败',
    type: 'warning',
  })
        })
        console.log(file,id);
    }
</script>

<template>
    <div class="upload">
        <el-card class="box-card">
            <el-form :model="form" label-width="120px" style="width: 500px">
                <el-form-item label="文件名">
                    <el-input v-model="form.name" />
                </el-form-item>
<!--                <el-form-item label="数据条目">-->
<!--                    <el-input v-model="form.num" />-->
<!--                </el-form-item>-->
                <el-form-item label="数据大小">

                    <el-input v-model="form.size" />
                </el-form-item>
                <el-form-item label="文件描述"  >
                    <div class="input">  
                        <el-input type="textarea" v-model="form.description" />
                    </div>                  
                </el-form-item>


                <el-form-item label="可转发">
                    <el-radio-group v-model="form.resource">
                        <el-radio label="是" :value="1" />
                        <el-radio label="否" :value="0" />
                    </el-radio-group>
                </el-form-item>
                <el-form-item style="display: flex;align-items: center;height: 60px;">
                    <el-upload
                            v-model:file-list="fileList"
                            class="upload-demo"
                            action=""
                            :limit="1"
                            :on-success="handleAvatarSuccess"
                            :on-remove="handleRemove"
                            :before-upload="beforeUpload"
                            style="display: flex;"
                    >
                        <el-button type="primary" >上传文件</el-button>
                    </el-upload>
<!--                    <el-button type="primary" @click="onSubmit">上传文件</el-button>-->
                    <el-button style="margin-left: 20px;">确认上传</el-button>
                </el-form-item>
            </el-form>
        </el-card>

    </div>
</template>



<style lang="scss" scoped>
    .upload{
        height: 100%;
        width: 100%;
    }
    .input{
        height: 500%;
        width: 500%;
    }
</style>
