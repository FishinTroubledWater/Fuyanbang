<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    反馈列表-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold"> 反馈列表</div>
      <el-table :data="tableData" style="width: 100%" border stripe>
        <!--        索引-->
        <el-table-column type="index"></el-table-column>
        <!--        用户名-->
        <el-table-column label="用户名" width="180">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.userID }}</span>
          </template>
        </el-table-column>
        <!--        时间-->
        <el-table-column label="时间" width="250">
          <template slot-scope="scope">
            <i class="el-icon-date"></i>
            <span style="margin-left: 10px">{{ scope.row.time }}</span>
          </template>
        </el-table-column>
        <!--        反馈内容-->
        <el-table-column label="反馈内容" width="350">
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.content }}</span>
          </template>
        </el-table-column>
        <!--        状态-->
        <el-table-column label="状态" width="150">
          <template slot-scope="scope">
           <el-tag type = "success" v-if="scope.row.state==='true'">已处理</el-tag>
            <el-tag type = "warning" v-else>未处理</el-tag>
          </template>
        </el-table-column>
        <!--        操作-->
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button size="mini" type="primary" icon="el-icon-view" round
                       @click="handleEdit(scope.$index, scope.row)">详情
            </el-button>
            <el-button size="mini" type="success" icon="el-icon-finished" round
                       @click="handleEdit(scope.$index, scope.row)">处理
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
  name: "FeedBackView",
  components: {Pagination, Breadcrumb},
  data() {
    const item = [{
      time: '2016-05-02',
      userID: '王小虎',
      content: '这是反馈内容啊啊啊啊',
      state:'true'
    },
      {
        time: '2016-05-02',
        userID: '王小虎',
        content: '这是反馈内容啊啊啊啊',
        state:'false'
      }];
    return {
      title: '反馈中心',
      tableData: Array(5).concat(item),
    }
  }
}
</script>

<style scoped>

</style>