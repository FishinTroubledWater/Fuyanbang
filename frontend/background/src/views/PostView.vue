<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    帖子查询-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold;"> 帖子查询</div>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入帖子" v-model="queryInfo.query">
            <el-button slot="append" icon="el-icon-search" @click="getPostList"></el-button>
          </el-input>
        </el-col>

        <el-col :span=2>
          <el-button type="primary" round><i class="el-icon-plus"></i> 添加帖子</el-button>
        </el-col>
      </el-row>
    </el-card>
    <!--    帖子列表-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold"> 帖子列表</div>
      <Table :table-data="postList" :columns="columns" :show-state="true">
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
                     @click="handleEdit(scope.$index, scope.row)"
                     :disabled="scope.row.state=== '1'">审核
          </el-button>
          <el-button size="mini" type="primary" icon="el-icon-edit" round
                     @click="handleEdit(scope.$index, scope.row)">编辑
          </el-button>
          <el-button size="mini" type="danger" icon="el-icon-delete" round
                     @click="removePostById(scope.row.ID)">删除
          </el-button>
        </template>
      </Table>
    </el-card>
    <!--    分页器-->
    <Pagination :total="total" :query-info="queryInfo"
                @page-size-change="handlePageSizeChange"
                @page-change="handlePageChange"></Pagination>
  </div>
</template>

<script>


export default {
  name: "post",
  data() {
    return {
      title: '帖子管理',
      //数据总数
      total: 0,
      //获取帖子列表的参数对象
      queryInfo: {
        query: '',
        pageNum: 1,
        pageSize: 10
      },
      //帖子数据集合
      postList: [],
      columns: [
        {prop: 'AuthorID', label: '用户名', width: '150px'},
        {prop: 'PublishTime', label: '发布时间', width: '180px', sortable: true},
        {prop: 'PartID', label: '板块', width: '180px', sortable: true},
        {prop: 'Summary', label: '内容', width: '180px' , showOverflowTooltip:true},
      ],
    }
  },
  created() {
    this.getPostList()
  },
  methods: {
    // 查询方法
    async getPostList() {
      const {data: res} = await this.axios.post('post/list', {
        pageSize: this.queryInfo.pageSize,
        pageNum: this.queryInfo.pageNum
      }, {
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      });
      console.log(res)
      if (res.code !== 200) return this.$message.error('获取帖子列表失败！')
      this.postList = res.data.posts
      this.total = res.data.total
      console.log(this.postList);
    },
    //处理每页显示数量变化
    handlePageSizeChange(newSize) {
      this.queryInfo.pageSize = newSize;
      this.getPostList()
    },
    //处理页码变化
    handlePageChange(newPage) {
      this.queryInfo.pageNum = newPage;
      this.getPostList()
    },

    //删除帖子
    async removePostById(id) {
      const result = await this.$confirm('确定要删除该帖子？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).catch(err => err)
      console.log(result)
      if (result !== 'confirm') {
        return this.$message.info('已取消删除')
      }
      console.log(id);
      const {data: res} = await this.axios.delete('post/delete', {
        params:{'id': id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) return this.$message.error('删除帖子失败！')
      this.$message.success('删除帖子成功！')
      await  this.getPostList()
    }
  }
}
</script>

<style scoped>

</style>