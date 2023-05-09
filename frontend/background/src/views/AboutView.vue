<template>
  <div class="about">
    <Table :table-data="userList" :columns="columns">
      <template #default="scope">
        <el-row :gutter="20">
          <el-col :span="10"><el-button type="primary" @click="handle(scope.row)">编辑</el-button></el-col>
          <el-col :span="10"><el-button type="success">ni</el-button></el-col>
        </el-row>
      </template>
    </Table>
  </div>
</template>
<script>


export default {
  data() {
    return {
      userList: [],
      queryInfo: {
        query: '',
        pageNum: 1,
        pageSize: 2
      },
      students: [
        {name: '张三', age: 18, score: 90},
        {name: '李四', age: 19, score: 85},
        {name: '王五', age: 20, score: 95}
      ],
      columns: [
        {prop: 'Account', label: '姓名', width: '120px'},
        {prop: 'Age', label: '年龄', width: '100px'},
        {prop: 'College', label: '成绩', width: '100px'}],
    }
  },
  created() {
    this.getUserList()
  },
  methods:{
    async getUserList() {
      const {data: res} = await this.axios.post('user/list', {
        pageSize: this.queryInfo.pageSize,
        pageNum: this.queryInfo.pageNum
      }, {
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      });
      console.log(res)
      if (res.code !== 200) return this.$message.error('获取用户列表失败！')
      this.userList = res.data.users
      this.total = res.data.total
      console.log(this.userList);
    },
    handle(row) {
      console.log(row.Account);
    },

  }
}
</script>