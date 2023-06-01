<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    评论查询-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold;"> 评论查询</div>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入帖子id" v-model="queryInfo.query">
            <el-button slot="append" icon="el-icon-search" @click="getCommentList"></el-button>
          </el-input>
        </el-col>
        <el-col :span=2>
          <el-button @click="addDialogVisible = true" type="primary"
                     round icon="el-icon-plus">添加评论
          </el-button>
        </el-col>
      </el-row>
    </el-card>
    <!--    评论列表-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold"> 评论列表</div>
      <Table :table-data="commentList" :columns="columns" :show-state="true">
        <!--        状态区-->
        <template #state="scope" >
          <el-tag v-if="scope.row.State === 0" type="warning">未审核</el-tag>
          <el-tag v-else type="success">已审核</el-tag>

        </template>
        <!--        操作区-->
        <template #operation="scope">
          <el-button size="mini" type="success" icon="el-icon-view" round
                     @click="showDetails(scope.row)">详情
          </el-button>
          <el-button size="mini" type="warning" icon="el-icon-finished" round
                     @click="checkCommentById(scope.row.ID)"
                     :disabled="scope.row.State=== 1">审核
          </el-button>
          <el-button size="mini" type="primary" icon="el-icon-edit" round
                     @click="showEditDialog(scope.row.ID)">编辑
          </el-button>
          <el-button size="mini" type="danger" icon="el-icon-delete" round
                     @click="removeCommentById(scope.row.ID)">删除
          </el-button>
        </template>
      </Table>
    </el-card>
    <!--    添加评论的对话框-->
    <el-dialog title="添加评论" :visible.sync="addDialogVisible" width="50%"
               @close="addDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="addFormRef" :model="addForm" label-width="80px"
               :rules="addFormRules">
        <el-form-item label="用户ID" prop="userID">
          <el-input v-model="addForm.userID"></el-input>
        </el-form-item>
        <el-form-item label="帖子ID" prop="targetPost">
          <el-input v-model="addForm.targetPost"></el-input>
        </el-form-item>
        <el-form-item label="状态" prop="state">
          <el-radio v-model="addForm.state" label='0'>未审核</el-radio>
          <el-radio v-model="addForm.state" label='1'>已审核</el-radio>
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <Editor :value.sync="addForm.content"> </Editor>
<!--          <quill-editor v-model="addForm.content"></quill-editor>-->
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
         <el-button @click="addDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="addPost">确 定</el-button>
      </span>
    </el-dialog>
    <!--    修改评论对话框-->
    <el-dialog title="修改评论" :visible.sync="editDialogVisible" width="50%"
               @close="editDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="editFormRef" :model="editForm" label-width="80px"
               :rules="addFormRules">
        <el-form-item label="用户ID" prop="Author">
          <el-input v-model="editForm.UserID" disabled></el-input>
        </el-form-item>
        <el-form-item label="状态" prop="State">
          <el-radio v-model="editForm.State" :label='0'>未审核</el-radio>
          <el-radio v-model="editForm.State" :label='1'>已审核</el-radio>
        </el-form-item>
        <el-form-item label="内容" prop="Content">
          <Editor :value.sync="editForm.Content"> </Editor>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="editDialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="editPostInfo">确 定</el-button>
  </span>
    </el-dialog>
    <!--    分页器-->
    <Pagination :total="total" :query-info="queryInfo"
                @page-size-change="handlePageSizeChange"
                @page-change="handlePageChange"></Pagination>
    <!--    抽屉-->
    <Drawer :drawer="detailsDrawer" :title="drawerTitle" @closed="drawerClosed">
      <template #details="scope">
        <el-descriptions direction="vertical" :column="2" border class="mg">
          <el-descriptions-item label="用户ID"> {{ editForm.UserID }}</el-descriptions-item>
          <el-descriptions-item label="评论数"> {{ editForm.CommentNum }}</el-descriptions-item>
          <el-descriptions-item label="帖子"> {{ editForm.TargetPost }}</el-descriptions-item>
          <el-descriptions-item label="发布时间"> {{ formattedPublishTime}}</el-descriptions-item>
          <el-descriptions-item label="状态">
            <el-tag v-if="editForm.State === 0" type="warning">未审核</el-tag>
            <el-tag v-else type="success">已审核</el-tag>
          </el-descriptions-item>
        </el-descriptions>
        <div class="mg">内容</div>
        <quill-editor v-model="editForm.Content" @focus="focus($event)" class="mg" ></quill-editor>
      </template>
    </Drawer>

  </div>
</template>

<script>

export default {
  name: "CommentView",
  computed:{
    formattedPublishTime() {
      const publishTime = this.editForm.PublishTime;
      return this.$moment(publishTime).format('YYYY-MM-DD HH:mm:ss');
    }
  },
  data(){
    return{
      title: '评论管理',
      drawerTitle: "评论详情",
      //添加对话框
      addDialogVisible: false,
      //控制详情抽屉可见否
      detailsDrawer: false,
      //控制修改对话框的显示与隐藏
      editDialogVisible: false,
      //数据总数
      total: 0,
      //获取列表的参数对象
      queryInfo: {
        query: "",
        pageNum: 1,
        pageSize: 10
      },
      //查询到的信息
      editForm: {
        Content:"",
      },
      //添加的表单数据
      addForm: {
        targetPost:'',
        userID: '',
        state: '0',
        content:''
      },
      // 添加规则
      addFormRules: {
        userID: [
          {required: true, message: '请输入用户id', trigger: 'blur'},
        ],
        targetPost: [
          {required: true, message: '请输入帖子id', trigger: 'blur'},
        ],
      },
      commentList:[],
      columns: [
        {prop: 'UserID', label: '用户ID', width: '150px'},
        {prop: 'PublishTime', label: '发布时间', width: '200px', sortable: true,
          formatter:(row, column)=>{
            const publishTime = row[column.property];
            return this.$moment(publishTime).format('YYYY-MM-DD HH:mm:ss');
          }},
        {prop: 'Content', label: '评论', width: '180px',showOverflowTooltip:true },
        {prop: 'CommentNum', label: '评论数', width: '180px', sortable: true},
      ],
    }
  },
  created() {
    this.getCommentList()
  },
  methods:{
    // 查询方法
    async getCommentList() {
      const {data: res} = await this.axios.post('comment/list', {
        query: this.queryInfo.query,
        pageSize: this.queryInfo.pageSize,
        pageNum: this.queryInfo.pageNum
      }, {
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      });
      console.log(res)
      if (res.code !== 200) return this.$message.error('获取评论列表失败！')
      this.commentList = res.data.comments
      this.total = res.data.total
      console.log(this.commentList);
    },
    //处理每页显示数量变化
    handlePageSizeChange(newSize) {
      this.queryInfo.pageSize = newSize;
      this.getCommentList()
    },
    //处理页码变化
    handlePageChange(newPage) {
      this.queryInfo.pageNum = newPage;
      this.getCommentList()
    },

    //删除
    async removeCommentById(id) {
      const result = await this.$confirm('确定要删除该评论？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).catch(err => err)
      console.log(result)
      if (result !== 'confirm') {
        return this.$message.info('已取消删除')
      }
      console.log(id);
      const {data: res} = await this.axios.delete('comment/delete', {
        params: {'id': id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) return this.$message.error('删除评论失败！')
      this.$message.success('删除评论成功！')
      await this.getCommentList()
    },
    //审核
    async checkCommentById(id) {
      const result = await this.$confirm('确定要通过该评论？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'success'
      }).catch(err => err)
      console.log(result)
      if (result !== 'confirm') {
        return this.$message.info('已取消通过')
      }
      console.log(id);
      // 发起审核评论信息的数据请求
      this.editForm.State = '1'
      console.log(id);
      const {data: res} = await this.axios.patch('comment/update', {
        'id': id,
        'state': this.editForm.State
      },{
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) {
        return this.$message.error('审核评论失败！')
      }
      this.$message.success('审核成功')
      await this.getCommentList()
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
    //添加对话框关闭
    addDialogClosed() {
      this.$refs.addFormRef.resetFields()
    },
    //点击按钮添加
    addPost() {
      this.$refs.addFormRef.validate(async valid => {
        if (!valid) return
        console.log(this.addForm)
        // 发起添加评论网络请求
        const {data: res} = await this.axios.post('comment/add', this.addForm, {
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          this.$message.error('添加评论失败！')
        } else this.$message.success('添加评论成功！')
        //隐藏对话框
        this.addDialogVisible = false
        //刷新用户列表
        await this.getCommentList()
      })
    },
    //展示编辑的对话框
    async showEditDialog(id) {
      console.log(id);
      const {data: res} = await this.axios.get('comment/searchById', {
        params: {'id':id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) {
        return this.$message.error('查询评论信息失败！')
      }
      this.editForm = res.data
      this.editDialogVisible = true
    },
    // 修改对话框关闭
    editDialogClosed() {
      this.$refs.editFormRef.resetFields()
    },
    // 修改信息并提交
    editPostInfo() {
      this.$refs.editFormRef.validate(async valid => {
        console.log(valid)
        if (!valid) return
        console.log(this.editForm)
        // 发起修改评论信息的数据请求
        const {data: res} = await this.axios.patch('comment/update', {
          'id': this.editForm.ID,
          'content': this.editForm.Content,
          'state': this.editForm.State
        },{
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          return this.$message.error('更新评论失败！')
        }
        // 关闭对话框
        this.editDialogVisible = false
        // 刷新数据列表
        await this.getCommentList()
        this.$message.success('更新评论成功！')
      })
    },
    //获取焦点事件
    focus(event){
      event.enable(false);
    },
  },

}
</script>

<style scoped>

</style>