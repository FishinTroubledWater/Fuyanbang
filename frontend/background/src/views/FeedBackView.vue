<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    反馈列表-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold"> 反馈列表</div>
      <Table :table-data="feedbackList" :columns="columns" :show-state="true">
        <!--        状态区-->
        <template #state="scope">
          <el-tag v-if="scope.row.state === '0'" type="warning">未处理</el-tag>
          <el-tag v-else type="success">已处理</el-tag>
        </template>
        <!--        操作区-->
        <template #operation="scope">
          <el-button size="mini" type="success" icon="el-icon-view" round
                     @click="showDetails(scope.row)">详情
          </el-button>
          <el-button size="mini" type="warning" icon="el-icon-finished" round
                     @click="finishFeedBackById(scope.row.ID)"
                     :disabled="scope.row.state=== '1'">处理
          </el-button>
        </template>
      </Table>
    </el-card>
    <Pagination></Pagination>
  </div>
</template>

<script>


export default {
  name: "FeedBackView",
  computed:{
    formattedPublishTime() {
      const Time = this.editForm.Time;
      return this.$moment(Time).format('YYYY-MM-DD HH:mm:ss');
    },
  },
  data() {
    return {
      title: '反馈中心',
      drawerTitle: "反馈详情",
      //控制详情抽屉可见否
      detailsDrawer: false,
      //数据总数
      total: 0,
      //获取列表的参数对象
      queryInfo: {
        query: '',
        pageNum: 1,
        pageSize: 10
      },
      //查询到的信息
      editForm: {

      },
      //反馈数据集合
      feedbackList: [],
      columns: [
        {prop: 'account', label: '用户名', width: '180px'},
        {prop: 'time', label: '时间', width: '200px', sortable: true,
          formatter: (row, column) => {
            const time = row[column.property];
            return this.$moment(time).format('YYYY-MM-DD HH:mm:ss');
          }},
        {prop: 'content', label: '内容', width: '180px', showOverflowTooltip:true},
      ],
    }
  },
  created() {
    this.getFeedBackList()
  },
  methods:{
  // 查询方法
    async getFeedBackList() {
      const {data: res} = await this.axios.post('feedback/list', {
        query: this.queryInfo.query,
        pageSize: this.queryInfo.pageSize,
        pageNum: this.queryInfo.pageNum
      }, {
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      });
      console.log(res)
      if (res.code !== 200) return this.$message.error('获取学长学姐说列表失败！')
      this.feedbackList = res.data.feedbacks
      this.total = res.data.total
      console.log(this.feedbackList);
    },
    //处理反馈
    async finishFeedBackById(id) {
      const result = await this.$confirm('确定要处理该反馈？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'success'
      }).catch(err => err)
      console.log(result)
      if (result !== 'confirm') {
        return this.$message.info('已取消处理')
      }
      console.log(id);
      const {data: res} = await this.axios.patch('feedback/update', {
        'id': this.editForm.ID,
        'state': this.editForm.State
      },{
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) {
        return this.$message.error('处理反馈失败！')
      }
      // const {data: res} = await this.axios.delete('feedback/delete', {
      //   params: {'id': id},
      //   headers: {
      //     'Authorization': window.sessionStorage.getItem("token")
      //   }
      // })
      // if (res.code !== 200) return this.$message.error('删除反馈失败！')
      // this.$message.success('删除反馈成功！')
      // await this.getFeedBackList()
    },
    //处理每页显示数量变化
    handlePageSizeChange(newSize) {
      this.queryInfo.pageSize = newSize;
      this.getFeedBackList()
    },
    //处理页码变化
    handlePageChange(newPage) {
      this.queryInfo.pageNum = newPage;
      this.getFeedBackList()
    },
    // 显示详情
    showDetails(row) {
      console.log(row);
      this.editForm=row;
      this.detailsDrawer = true;
    },
    //关闭详情
    drawerClosed() {
      this.detailsDrawer = false;
    },
    //获取焦点事件
    focus(event){
      event.enable(false);
    }
  }
}
</script>

<style scoped>

</style>