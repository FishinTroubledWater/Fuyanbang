<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    内容查询-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold;"> 评论查询</div>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入内容" >
            <el-button slot="append" icon="el-icon-search" @click="getUserList"></el-button>
          </el-input>
        </el-col>
        <el-col :span="5">
          <el-button
              v-for="type in types"
              :key="type.name"
              size="small"
          >
            {{ type.name }}
          </el-button>
        </el-col>
        <el-col :span=2>
          <el-button type="primary" round ><i class="el-icon-plus"></i> 添加内容</el-button>
        </el-col>
      </el-row>
    </el-card>
    <!--    动态列表-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold"> 评论列表</div>
      <Table :table-data="commentList" :columns="columns" :show-state="true">
        <template #operation="scope">
          <el-button size="mini" type="success" icon="el-icon-view" round
                     @click="handleEdit(scope.$index, scope.row)">详情
          </el-button>
          <el-button size="mini" type="warning" icon="el-icon-finished" round
                     @click="handleEdit(scope.$index, scope.row)">审核
          </el-button>
          <el-button size="mini" type="danger" icon="el-icon-delete" round
                     @click="handleDelete(scope.$index, scope.row)">删除
          </el-button>
        </template>
      </Table>
    </el-card>
  </div>
</template>

<script>

export default {
  name: "CommentView",
  data(){
    return{
      title: '评论管理',
      commentList:[
        {userID:'wxy',publishTime:'2022',_like:'1500',content:'这是内容aaaaaaaaaaaaaaaa',state:'1' },
        {userID:'wxy',publishTime:'2022',_like:'1500',content:'这是内容aaaa',state:'0' }
      ],
      columns: [
        {prop: 'userID', label: '用户ID', width: '150px'},
        {prop: 'publishTime', label: '发布时间', width: '180px', sortable: true},
        {prop: 'content', label: '评论内容', width: '180px',showOverflowTooltip:true },
        {prop: '_like', label: '点赞数', width: '180px', sortable: true},
      ],
    }
  }
}
</script>

<style scoped>

</style>