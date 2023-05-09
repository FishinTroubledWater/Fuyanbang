<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    内容查询-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold;"> 学长学姐说查询</div>
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
      <div style="font-size: 20px;font-weight: bold"> 学长学姐说列表</div>
      <Table :table-data="recipeList" :columns="columns" :show-state="true">
        <!--        状态区-->
        <template #state="scope">
          <el-tag v-if="scope.row.state === '0'" type="warning">未审核</el-tag>
          <el-tag v-else type="success">已审核</el-tag>
        </template>
        <!--        操作区-->
        <template #operation="scope">
          <el-button size="mini" type="success" icon="el-icon-view" round
                     @click="handleEdit(scope.$index, scope.row)">详情
          </el-button>
          <el-button size="mini" type="warning" icon="el-icon-finished" round
                     @click="handleEdit(scope.$index, scope.row)">审核
          </el-button>
          <el-button size="mini" type="primary" icon="el-icon-edit" round
                     @click="handleEdit(scope.$index, scope.row)">编辑
          </el-button>
          <el-button size="mini" type="danger" icon="el-icon-delete" round
                     @click="handleDelete(scope.$index, scope.row)">删除
          </el-button>
        </template>
      </Table>
    </el-card>
    <Pagination :total="total" :query-info="queryInfo"
                @page-size-change="handlePageSizeChange"
                @page-change="handlePageChange"></Pagination>
  </div>
</template>

<script>
import Pagination from "@/components/Pagination";
import Breadcrumb from "@/components/Breadcrumb";

export default {
  name: "RecipeView",
  components: {Pagination,Breadcrumb},
  data(){
    return{
      title: '学长学姐说管理',
      recipeList:[
        {author:'wxy',publishTime:'2022',title:'测试'}
      ],
      columns: [
        {prop: 'author', label: '用户名', width: '150px'},
        {prop: 'publishTime', label: '发布时间', width: '180px', sortable: true},
        {prop: 'title', label: '标题', width: '180px'},
      ],
    }
  }
}
</script>

<style scoped>

</style>