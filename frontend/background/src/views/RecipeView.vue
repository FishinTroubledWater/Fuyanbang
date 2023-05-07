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
      <el-table :data="tableData" style="width: 100%" border stripe>
        <!--        索引-->
        <el-table-column type="index"></el-table-column>
        <!--        用户名-->
        <el-table-column label="用户名" width="180">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.author }}</span>
          </template>
        </el-table-column>
        <!--        时间-->
        <el-table-column label="时间" width="250">
          <template slot-scope="scope">
            <i class="el-icon-date"></i>
            <span style="margin-left: 10px">{{ scope.row.publishTime }}</span>
          </template>
        </el-table-column>
        <!--        标题-->
        <el-table-column label="标题" width="250">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.title }}</span>
          </template>
        </el-table-column>

        <!--        操作-->
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button size="mini" type="primary" icon="el-icon-edit" round
                       @click="handleEdit(scope.$index, scope.row)">编辑
            </el-button>
            <el-button size="mini" type="danger" icon="el-icon-delete" round
                       @click="handleDelete(scope.$index, scope.row)">删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <Pagination></Pagination>
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
      tableData:[{
        author:'渚薰',
        publishTime:'2016-05-02',
        title:'死海文书',

      },{
        author:'王小虎',
        publishTime:'2015-05-02',
        title:'死灵之书',

      }],
    }
  }
}
</script>

<style scoped>

</style>