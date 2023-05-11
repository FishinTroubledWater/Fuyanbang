<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    用户查询-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold;">用户查询</div>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入内容" v-model="queryInfo.query">
            <el-button slot="append" icon="el-icon-search" @click="getUserList"></el-button>
          </el-input>
        </el-col>
        <el-col :span=2>
          <el-button @click="addDialogVisible = true" type="primary"
                     round icon="el-icon-plus">添加用户
          </el-button>
        </el-col>
      </el-row>
    </el-card>
    <!--    用户列表-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold"> 用户列表</div>
      <Table :table-data="userList" :columns="columns" >
        <template #operation="scope">
          <el-button size="mini" type="success" icon="el-icon-view" round
                     @click="showDetails(scope.row)">详情
          </el-button>
          <el-button size="mini" type="primary" icon="el-icon-edit" round
                     @click="showEditDialog(scope.row.Account)">编辑
          </el-button>
          <el-button size="mini" type="danger" icon="el-icon-delete" round
                     @click="removeUserById(scope.row.Account)">删除
          </el-button>
        </template>
      </Table>
    </el-card>
    <!--    分页器-->
    <Pagination :total="total" :query-info="queryInfo"
                @page-size-change="handlePageSizeChange"
                @page-change="handlePageChange"></Pagination>
    <!--    添加用户的对话框-->
    <el-dialog title="添加用户" :visible.sync="addDialogVisible" width="30%"
               @close="addDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="addFormRef" :model="addForm" label-width="80px"
               :rules="addFormRules">
        <el-form-item label="用户名" prop="account">
          <el-input v-model="addForm.account"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="addForm.password"></el-input>
        </el-form-item>
        <el-form-item label="手机号" prop="phonenumber">
          <el-input v-model="addForm.phonenumber"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
         <el-button @click="addDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="addUser">确 定</el-button>
      </span>
    </el-dialog>
    <!--    修改用户对话框-->
    <el-dialog title="修改用户" :visible.sync="editDialogVisible" width="30%"
               @close="editDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="editFormRef" :model="editForm" label-width="80px"
               :rules="editFormRules">
        <el-form-item label="用户名" prop="Account">
          <el-input v-model="editForm.Account" disabled></el-input>
        </el-form-item>
        <el-form-item label="手机号" prop="Phonenumber">
          <el-input v-model="editForm.PhoneNumber"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="editDialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="editUserInfo">确 定</el-button>
  </span>
    </el-dialog>
    <!--    抽屉-->
    <Drawer  :drawer="detailsDrawer" :title="drawerTitle" @closed="drawerClosed">
      <template #details="scope" >
        <div class="mg setcenter"><el-image :src="editForm.AvatarUrl" style="height: 150px"></el-image></div>
        <el-descriptions  direction="vertical" :column="2" border class="mg">
          <el-descriptions-item label="用户ID" > {{editForm.ID}}</el-descriptions-item>
          <el-descriptions-item label="用户名" > {{editForm.Account}}</el-descriptions-item>
          <el-descriptions-item label="昵称">{{editForm.NickName}}</el-descriptions-item>
          <el-descriptions-item label="手机号">{{editForm.PhoneNumber}}</el-descriptions-item>
          <el-descriptions-item label="地址">{{editForm.Area}}</el-descriptions-item>
          <el-descriptions-item label="性别">{{editForm.Sex}}</el-descriptions-item>
          <el-descriptions-item label="目标院校">{{editForm.TargetCollege}}</el-descriptions-item>
          <el-descriptions-item label="目标专业">{{editForm.TargetMajor}}</el-descriptions-item>
          <el-descriptions-item label="专业">{{editForm.Major}}</el-descriptions-item>
          <el-descriptions-item label="年级"><el-tag size="small">{{editForm.Year}}</el-tag></el-descriptions-item>
          <el-descriptions-item label="标语" :span="2">
            <el-input type="textarea" v-model="editForm.Slogan" :rows="5" readonly resize="none"></el-input>
            <quill-editor v-model="editForm.Slogan" @focus="focus($event)"></quill-editor>
          </el-descriptions-item>
        </el-descriptions>

      </template>
    </Drawer>
  </div>
</template>

<script>
export default {
  name: "UserView",

  data() {
    // 验证手机号的规则
    var checkPhoneNumber = (rule, value, cb) => {
      const regPhoneNumber = /^1(3\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\d|9[0-35-9])\d{8}$/
      if (regPhoneNumber.test(value)) {
        return cb()
      }
      cb(new Error('请输入合法的手机号'))
    };
    return {
      title: "用户管理",
      drawerTitle:"用户详情",
      total: 0,
      //用户数据集合
      userList: [],
      //查询到的用户信息
      editForm: {},
      //添加用户对话框
      addDialogVisible: false,
      //控制详情抽屉可见否
      detailsDrawer: false,
      //控制修改用户对话框的显示与隐藏
      editDialogVisible: false,
      //列表配置
      columns: [
        {prop: 'Account', label: '用户名', width: '150px'},
        {prop: 'NickName', label: '昵称', width: '150px'},
        {prop: 'PhoneNumber', label: '手机号', width: '150px'},
        {prop: 'Year', label: '年级', width: '100px', sortable: true},
        {prop: 'College', label: '大学', width: '100px'},
        {prop: 'Balance', label: '学币', width: '100px', sortable: true},
      ],
      //添加用户的表单数据
      addForm: {
        account: '',
        password: '',
        phonenumber: '',
      },
      // 添加用户的规则
      addFormRules: {
        account: [
          {required: true, message: '请输入用户名', trigger: 'blur'},
          {min: 3, max: 10, message: '用户名的长度在3~10个字符间', trigger: 'blur'}
        ],
        password: [
          {required: true, message: '请输入密码', trigger: 'blur'},
          {min: 6, max: 20, message: '密码的长度在6~20个字符间', trigger: 'blur'}
        ],
        phonenumber: [
          {required: true, message: '请输入手机号', trigger: 'blur'},
          {validator: checkPhoneNumber, trigger: 'blur'}
        ]
      },
      //编辑用户的规则
      editFormRules: {
        phonenumber: [
          {required: true, message: '请输入手机号', trigger: 'blur'},
          {validator: checkPhoneNumber,  trigger: 'blur'}
        ]
      },
      //假数据
      // 获取用户列表的参数对象
      queryInfo: {
        query: '',
        pageNum: 1,
        pageSize: 10
      },
    }
  },
  created() {
    this.getUserList()
  },
  methods: {
    // 查询方法
    async getUserList() {
      const {data: res} = await this.axios.post('user/list', {
        query: this.queryInfo.query,
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
    //处理每页显示数量变化
    handlePageSizeChange(newSize) {
      this.queryInfo.pageSize = newSize;
      this.getUserList()
    },
    //处理页码变化
    handlePageChange(newPage) {
      this.queryInfo.pageNum = newPage;
      this.getUserList()
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

    //添加用户对话框关闭
    addDialogClosed() {
      this.$refs.addFormRef.resetFields()
    },

    //点击按钮添加用户
    addUser() {
      this.$refs.addFormRef.validate(async valid => {
        if (!valid) return
        console.log(this.addForm)
        // 发起添加用户网络请求
        const {data: res} = await this.axios.post('user/add', this.addForm, {
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          this.$message.error('添加用户失败！')
        } else this.$message.success('添加用户成功！')
        //隐藏对话框
        this.addDialogVisible = false
        //刷新用户列表
        await this.getUserList()
      })
    },

    //展示编辑用户的对话框
    async showEditDialog(id) {
      console.log(id);
      const {data: res} = await this.axios.get('user/searchByAccount', {
        params: {'account': id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) {
        return this.$message.error('查询用户信息失败！')
      }
      this.editForm = res.data
      this.editDialogVisible = true
    },

    // 修改用户对话框关闭
    editDialogClosed() {
      this.$refs.editFormRef.resetFields()
    },

    // 修改信息并提交
    editUserInfo() {
      this.$refs.editFormRef.validate(async valid => {
        console.log(valid)
        if (!valid) return
        console.log(this.editForm.Account)
        console.log(this.editForm.PhoneNumber)
        // 发起修改用户信息的数据请求
        const {data: res} = await this.axios.patch('user/update', {
          'account': this.editForm.Account,
          'phoneNumber': this.editForm.PhoneNumber
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

    //删除用户
    async removeUserById(id) {
      const result = await this.$confirm('确定要删除该用户？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).catch(err => err)
      console.log(result)
      if (result !== 'confirm') {
        return this.$message.info('已取消删除')
      }
      console.log(id);
      const {data: res} = await this.axios.delete('user/delete', {
        params:{'account': id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) return this.$message.error('删除用户失败！')
      this.$message.success('删除用户成功！')
      await  this.getUserList()
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