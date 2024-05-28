<template>
    <el-text size="large">Docker Node Manager</el-text>
    <br>
    <el-button style="margin: 10px" @click="CreateDockerNodeFormVisible = true">Add</el-button>
    <el-table
        :data="dockerNodeList"
        >
        <el-table-column prop="ID" label="ID"/>
        <el-table-column prop="Host" label="Host"/>
        <el-table-column prop="Port" label="Port"/>
    </el-table>
    <el-dialog v-model="CreateDockerNodeFormVisible" title="Create Node" width="500">
        <el-form :model=CreateDockerNode>
            <el-form-item label="Host">
                <el-input v-model="CreateDockerNodeData.Host" />
            </el-form-item>
            <el-form-item label="Port">
                <el-input v-model="CreateDockerNodeData.Port" />
            </el-form-item>
        </el-form>
        <el-button @click=CreateDockerNode>Submit</el-button>
    </el-dialog>
</template>
<script>
import dockerNodeApi from "@/api/dockerNode.js"
import {ref} from "vue";
export default {
    data() {
        return {
            dockerNodeList: [],
            CreateDockerNodeFormVisible: false,
            CreateDockerNodeData: {
                "Host": "",
                "Port": "",
            },
        }
    },
    methods: {
        CreateDockerNode() {
            dockerNodeApi.CreateDockerNode(this.CreateDockerNodeData).then(res => {
                if (res.code === 0){
                    console.log(res.code)
                }
            }).catch(error => {
                console.log(error)
            })
        },
    },
    mounted() {
        dockerNodeApi.GetDockerNodeList().then(res => {
            if (res.code === 0){
                this.dockerNodeList = res.data
            }
        }).catch(error => {
            console.log(error)
        })
    },
}
</script>
