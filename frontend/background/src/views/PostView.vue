<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    帖子查询-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold;"> 帖子查询</div>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入用户名" v-model="queryInfo.query">
            <el-button slot="append" icon="el-icon-search" @click="getPostList"></el-button>
          </el-input>
        </el-col>
        <el-col :span=2>
          <el-button @click="addDialogVisible = true" type="primary"
                     round icon="el-icon-plus">添加帖子
          </el-button>
        </el-col>
      </el-row>
    </el-card>
    <!--    帖子列表-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold"> 帖子列表</div>
      <Table :table-data="postList" :columns="columns" :show-state="true">
        <!--        状态区-->
        <template #state="scope" >
          <el-tag v-if="scope.row.State === '0'" type="warning">未审核</el-tag>
          <el-tag v-else type="success">已审核</el-tag>

        </template>
        <!--        操作区-->
        <template #operation="scope">
          <el-button size="mini" type="success" icon="el-icon-view" round
                     @click="showDetails(scope.row)">详情
          </el-button>
          <el-button size="mini" type="warning" icon="el-icon-finished" round
                     @click="checkPostById(scope.row.ID)"
                     :disabled="scope.row.State=== '1'">审核
          </el-button>
          <el-button size="mini" type="primary" icon="el-icon-edit" round
                     @click="showEditDialog(scope.row)">编辑
          </el-button>
          <el-button size="mini" type="danger" icon="el-icon-delete" round
                     @click="removePostById(scope.row.ID)">删除
          </el-button>
        </template>
      </Table>
    </el-card>
    <!--    添加帖子的对话框-->
    <el-dialog title="添加帖子" :visible.sync="addDialogVisible" width="50%"
               @close="addDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="addFormRef" :model="addForm" label-width="80px"
               :rules="addFormRules">
        <el-form-item label="用户名" prop="account">
          <el-input v-model="addForm.account"></el-input>
        </el-form-item>
        <el-form-item label="板块" prop="part">
          <el-select v-model="addForm.part" placeholder="请选择">
            <el-option
                v-for="item in parts"
                :key="item.value"
                :label="item.label"
                :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="state">
          <el-radio v-model="addForm.state" label="0">未审核</el-radio>
          <el-radio v-model="addForm.state" label="1">已审核</el-radio>
        </el-form-item>
        <el-form-item label="概要" prop="summary">
          <el-input type="textarea" v-model="addForm.summary" :rows="5" readonly resize="none"></el-input>
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <quill-editor v-model="addForm.content"></quill-editor>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
         <el-button @click="addDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="addPost">确 定</el-button>
      </span>
    </el-dialog>
    <!--    修改帖子对话框-->
    <el-dialog title="修改帖子" :visible.sync="editDialogVisible" width="50%"
               @close="editDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="editFormRef" :model="editForm" label-width="80px"
               :rules="addFormRules">
        <el-form-item label="用户名" prop="Author">
          <el-input v-model="editForm.Author.Account" disabled></el-input>
        </el-form-item>
        <el-form-item label="板块" prop="PartID">
          <el-select v-model="editForm.PartID" placeholder="请选择">
            <el-option v-for="item in parts" :key="item.value"
                       :label="item.label" :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="State">
          <el-radio v-model="editForm.State" label="0">未审核</el-radio>
          <el-radio v-model="editForm.State" label="1">已审核</el-radio>
        </el-form-item>
        <el-form-item label="内容" prop="Summary">
          <quill-editor v-model="editForm.Summary"></quill-editor>
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
          <el-descriptions-item label="用户名"> {{ editForm.Author.Account }}</el-descriptions-item>
          <el-descriptions-item label="收藏数"> {{ editForm.Favorite }}</el-descriptions-item>
          <el-descriptions-item label="板块"> {{ editForm.Part.PartName }}</el-descriptions-item>
          <el-descriptions-item label="发布时间"> {{ formattedPublishTime}}</el-descriptions-item>
          <el-descriptions-item label="状态">
            <el-tag v-if="editForm.State === '0'" type="warning">未审核</el-tag>
            <el-tag v-else type="success">已审核</el-tag>
          </el-descriptions-item>
        </el-descriptions>
        <quill-editor v-model="editForm.Content" @focus="focus($event)" class="mg"></quill-editor>
      </template>
    </Drawer>

  </div>

</template>

<script>


export default {
  name: "post",
  computed:{
    formattedPublishTime() {
      const publishTime = this.editForm.PublishTime;
      return this.$moment(publishTime).format('YYYY-MM-DD HH:mm:ss');
    },
  },
  data() {
    return {
      title: '帖子管理',
      drawerTitle: "帖子详情",
      //添加帖子对话框
      addDialogVisible: false,
      //控制详情抽屉可见否
      detailsDrawer: false,
      //控制修改帖子对话框的显示与隐藏
      editDialogVisible: false,
      //数据总数
      total: 0,
      //获取帖子列表的参数对象
      queryInfo: {
        query: '',
        pageNum: 1,
        pageSize: 10
      },
      //板块数据集合
      parts:[
        {value:'1',label:'加油站'},
        {value:'2',label:'求解答'},
      ],
      //查询到的帖子信息
      editForm: {
        Author:{},
        Part:{}
      },
      //帖子数据集合
      postList: [],
      //添加帖子的表单数据
      addForm: {
        account: '',
        part: '1',
        summary: '',
        state: '0',
        content:''
      },
      // 添加帖子规则
      addFormRules: {
        account: [
          {required: true, message: '请输入用户名', trigger: 'blur'},
          {min: 3, max: 10, message: '用户名的长度在3~10个字符间', trigger: 'blur'}
        ],
      },
      //表格配置
      columns: [
        {prop: 'Author.Account', label: '用户名', width: '150px'},
        {prop: 'PublishTime', label: '发布时间', width: '200px', sortable: true,
          formatter:(row, column)=>{
            const publishTime = row[column.property];
            return this.$moment(publishTime).format('YYYY-MM-DD HH:mm:ss');
          }},
        {prop: 'Part.PartName', label: '板块', width: '150px', sortable: true},
        {prop: 'Summary', label: '概要', width: '180px', showOverflowTooltip: true},
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
        query: this.queryInfo.query,
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
        params: {'id': id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) return this.$message.error('删除帖子失败！')
      this.$message.success('删除帖子成功！')
      await this.getPostList()
    },
    //审核帖子
    async checkPostById(id) {
      const result = await this.$confirm('确定要通过该帖子？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'success'
      }).catch(err => err)
      console.log(result)
      if (result !== 'confirm') {
        return this.$message.info('已取消通过')
      }
      console.log(id);
      // 发起修改帖子信息的数据请求
      const {data: res} = await this.axios.patch('post/update', {
        'id': this.editForm.ID,
        'partID': this.editForm.PartID,
        'summary': this.editForm.Summary,
        'content': this.editForm.Content,
        'state': this.editForm.State
      },{
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) {
        return this.$message.error('更新用户信息失败！')
      }
      // const {data: res} = await this.axios.delete('post/delete', {
      //   params: {'id': id},
      //   headers: {
      //     'Authorization': window.sessionStorage.getItem("token")
      //   }
      // })
      // if (res.code !== 200) return this.$message.error('删除帖子失败！')
      // this.$message.success('删除帖子成功！')
      await this.getPostList()
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
    //添加帖子对话框关闭
    addDialogClosed() {
      this.$refs.addFormRef.resetFields()
    },
    //点击按钮添加帖子
    addPost() {
      this.$refs.addFormRef.validate(async valid => {
        if (!valid) return
        console.log(this.addForm)
        // 发起添加帖子网络请求
        const {data: res} = await this.axios.post('post/add', this.addForm, {
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          this.$message.error('添加帖子失败！')
        } else this.$message.success('添加帖子成功！')
        //隐藏对话框
        this.addDialogVisible = false
        //刷新用户列表
        await this.getPostList()
      })
    },
    //展示编辑用户的对话框
    async showEditDialog(row) {
      console.log(row);
      // const {data: res} = await this.axios.get('post/searchByAccount', {
      //   params: {'account': id},
      //   headers: {
      //     'Authorization': window.sessionStorage.getItem("token")
      //   }
      // })
      // if (res.code !== 200) {
      //   return this.$message.error('查询用户信息失败！')
      // }
      this.editForm = row
      // console.log(this.editForm.Author.Account)
      this.editDialogVisible = true
    },
    // 修改帖子对话框关闭
    editDialogClosed() {
      this.$refs.editFormRef.resetFields()
    },
    // 修改信息并提交
    editPostInfo() {
      this.$refs.editFormRef.validate(async valid => {
        console.log(valid)
        if (!valid) return
        console.log(this.editForm)
        // 发起修改帖子信息的数据请求
        const {data: res} = await this.axios.patch('post/update', {
          'postID': this.editForm.PostID,
          'partID': this.editForm.PartID,
          'summary': this.editForm.Summary,
          'content': this.editForm.Content,
          'state': this.editForm.State
        },{
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          return this.$message.error('更新用户信息失败！')
        }
        // 关闭对话框
        this.editDialogVisible = false
        // 刷新数据列表
        await this.getUserList()
        this.$message.success('更新用户信息成功！')
      })
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