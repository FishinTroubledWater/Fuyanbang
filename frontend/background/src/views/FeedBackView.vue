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
          <el-tag v-if="scope.row.State === 0" type="warning">未处理</el-tag>
          <el-tag v-else type="success">已处理</el-tag>
        </template>
        <!--        操作区-->
        <template #operation="scope">
          <el-button size="mini" type="success" icon="el-icon-view" round
                     @click="showDetails(scope.row)">详情
          </el-button>
          <el-button size="mini" type="warning" icon="el-icon-finished" round
                     @click="finishFeedBackById(scope.row.ID)"
                     :disabled="scope.row.State=== 1">处理
          </el-button>
        </template>
      </Table>
    </el-card>
    <!--    分页器-->
    <Pagination :total="total" :query-info="queryInfo"
                @page-size-change="handlePageSizeChange"
                @page-change="handlePageChange"></Pagination>
    <!--    抽屉-->
    <Drawer :drawer="detailsDrawer" :title="drawerTitle" @closed="drawerClosed">
      <template #details="scope">
        <el-descriptions direction="vertical" :column="2" border class="mg">
          <el-descriptions-item label="作者"> {{ editForm.Author.Account }}</el-descriptions-item>
          <el-descriptions-item label="时间"> {{ formattedPublishTime }}</el-descriptions-item>
          <el-descriptions-item label="状态">
            <el-tag v-if="editForm.State === 0" type="warning">未审核</el-tag>
            <el-tag v-else type="success">已审核</el-tag>
          </el-descriptions-item>
        </el-descriptions>
        <div class="mg">内容</div>
        <quill-editor v-model="editForm.Content" @focus="focus($event)" class="mg"></quill-editor>
      </template>
    </Drawer>
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
        query:'',
        pageNum: 1,
        pageSize: 10
      },
      //查询到的信息
      editForm: {

      },
      //反馈数据集合
      feedbackList: [],
      columns: [
        {prop: 'Author.Account', label: '用户名', width: '180px'},
        {prop: 'Time', label: '时间', width: '200px', sortable: true,
          formatter: (row, column) => {
            const time = row[column.property];
            return this.$moment(time).format('YYYY-MM-DD HH:mm:ss');
          }},
        {prop: 'Content', label: '内容', width: '180px', showOverflowTooltip:true},
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
      console.log(this.queryInfo);
      console.log(res)
      if (res.code !== 200) return this.$message.error('获取反馈列表失败！')
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
      this.editForm.State = 1
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
      this.$message.success('处理成功')
      await this.getFeedBackList()
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