<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    反馈列表-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold"> 反馈列表</div>
      <Table :table-data="feedbackList" :columns="columns">
        <template #default="scope">
          <el-button size="mini" type="success" icon="el-icon-view" round
                     @click="handleEdit(scope.$index, scope.row)">详情
          </el-button>
          <el-button size="mini" type="warning" icon="el-icon-finished" round
                     @click="handleEdit(scope.$index, scope.row)">处理
          </el-button>
        </template>
      </Table>
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

    return {
      title: '反馈中心',
      feedbackList: [
        {userID: 'wxy', time: '2022', email: '测试', content: '这是反馈内容啊啊啊啊啊'}
      ],
      columns: [
        {prop: 'userID', label: '用户ID', width: '150px'},
        {prop: 'time', label: '时间', width: '180px', sortable: true},
        {prop: 'email', label: '邮箱', width: '180px', sortable: true},
        {prop: 'content', label: '内容', width: '300px', showOverflowTooltip:true},
      ],
    }
  }
}
</script>

<style scoped>

</style>