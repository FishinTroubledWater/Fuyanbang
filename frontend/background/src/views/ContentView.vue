<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    用户查询-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold;"> 用户查询</div>
      <el-row :gutter="20">
        <el-col :span="24" style="margin: 5px">用户类别</el-col>
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
          <el-button type="primary" round ><i class="el-icon-plus"></i> 添加用户</el-button>
        </el-col>
      </el-row>
    </el-card>
    <!--    用户列表-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold"> 用户列表</div>
      <el-table :data="tableData" style="width: 100%" border stripe>
        <!--        索引-->
        <el-table-column type="index"></el-table-column>
        <!--        姓名-->
        <el-table-column label="姓名" width="180">
          <template slot-scope="scope">
            <el-popover trigger="hover" placement="top">
              <p>姓名: {{ scope.row.name }}</p>
              <div slot="reference" class="name-wrapper">
                <el-tag size="medium">{{ scope.row.name }}</el-tag>
              </div>
            </el-popover>
          </template>
        </el-table-column>
        <!--        联系方式-->
        <el-table-column label="联系方式" width="180">
          <template slot-scope="scope">
            <i class="el-icon-user"></i>
            <span style="margin-left: 10px">{{ scope.row.phoneNumber }}</span>
          </template>
        </el-table-column>
        <!--        用户院校-->
        <el-table-column label="用户院校" width="180">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.college }}</span>
          </template>
        </el-table-column>
        <!--        用户类别-->
        <el-table-column label="用户类别" width="180">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.role }}</span>
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
  name: "content",
  components: {Pagination,Breadcrumb},
  data(){
    return{
      title: '内容管理'
    }
  }
}
</script>

<style scoped>

</style>