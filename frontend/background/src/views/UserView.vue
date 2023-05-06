<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    用户查询-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold;"> 用户查询</div>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入内容" v-model="queryInfo.query">
            <el-button slot="append" icon="el-icon-search" @click="getUserList"></el-button>
          </el-input>
        </el-col>
        <el-col :span=2>
          <el-button @click="addDialogVisible = true" type="primary"
                     size="small" round icon="el-icon-plus">添加用户
          </el-button>
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
        <el-table-column label="联系方式" width="200">
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
                       @click="showEditDialog(scope.row.account)">编辑
            </el-button>
            <el-button size="mini" type="danger" icon="el-icon-delete" round
                       @click="removeUserById(scope.row.account)">删除
            </el-button>
          </template>
        </el-table-column>

      </el-table>
    </el-card>
    <Pagination></Pagination>
    <!--    <Pagination :total="total" :query-info="queryInfo"></Pagination>-->
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
        <el-form-item label="用户类别" prop="type">
          <el-select v-model="addForm.type" placeholder="请选择">
            <el-option label="普通用户" value="普通用户"></el-option>
            <el-option label="管理员" value="管理员"></el-option>
          </el-select>
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
        <el-form-item label="用户名" prop="account" >
          <el-input v-model="editForm.account" disabled></el-input>
        </el-form-item>
        <el-form-item label="手机号" prop="phonenumber">
          <el-input v-model="editForm.phonenumber"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="editDialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="editUserInfo">确 定</el-button>
  </span>
    </el-dialog>
  </div>
</template>

<script>
import Pagination from "@/components/Pagination";
import Breadcrumb from "@/components/Breadcrumb";

export default {
  name: "UserView",
  components: {Pagination, Breadcrumb},
  data() {
    // 验证手机号的规则
    var checkPhoneNumber = (rule, value, cb) => {
      const regPhoneNumber = /^1(3\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\d|9[0-35-9])\d{8}$/
      if (regPhoneNumber.test(value)) {
        return cb()
      }
      cb(new Error('请输入合法的手机号'))
    };
    const item = {
      date: '2016-05-02',
      name: '王小虎',
      phoneNumber: '110',
      role: '普通用户',
      college: '福州大学',
    };
    return {
      title: "用户管理",
      total: 0,
      userList: [],
      //添加用户对话框
      addDialogVisible: false,
      //添加用户的表单数据
      addForm: {
        account: '',
        password: '',
        phonenumber: '',
        type: '普通用户',
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
          {validator: checkPhoneNumber, trigger: 'blur'}
        ]
      },
      //假数据
      tableData: Array(10).fill(item),
      // 获取用户列表的参数对象
      queryInfo: {
        query: '',
        pagenum: 1,
        pagesize: 2
      },
      //控制修改用户对话框的显示与隐藏
      editDialogVisible: false,
      //查询到的用户信息
      editForm: {},

    }
  },
  created() {
    // this.getUserList()
  },
  methods: {
    // 查询方法
    async getUserList() {
      const {data: res} = await this.axios.get('user/list', {
        params: this.queryInfo
      })
      if (res.code !== 200) return this.$message.error('获取用户列表失败！')
      this.userList = res.data.users
      this.total = res.data.total
      console.log(res)
    },

    // 编辑方法
    handleEdit(index, row) {
      console.log(index, row);
    },

    // 删除方法
    handleDelete(index, row) {
      console.log(index, row);
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
        //发起添加用户网络请求
        // const {data: res} = await this.axios.post('user/add', this.addForm)
        //
        // if(res.code !== 201){
        //   this.$message.error('添加用户失败！')
        // }
        // this.$message.success('添加用户成功！')
        // //隐藏对话框
        // this.addDialogVisible = false
        // //刷新用户列表
        // this.getUserList()
      })
    },

    //展示编辑用户的对话框
    async showEditDialog(id) {
      // const {data:res} = await this.axios.get('user/' + id)
      // if(res.code !== 200){
      //   return this.$message.error('查询用户信息失败！')
      // }
      // this.editForm = res.data
      console.log(id);
      this.editDialogVisible = true
    },

    // 修改用户对话框关闭
    editDialogClosed(){
      this.$refs.editFormRef.resetFields()
    },

    // 修改信息并提交
    editUserInfo(){
      this.$refs.editFormRef.validate(async valid => {
        console.log(valid)
        if(!valid) return
        // 发起修改用户信息的数据请求
        // const {data : res} = await this.axios.put('users/' + this.editForm.id,{
        //   phonenumber: this.editForm.phonenumber,
        // })
        //
        // if(res.code !== 200){
        //   return this.$message.error('更新用户信息失败！')
        // }
        // // 关闭对话框
        // this.editDialogVisible = false
        // // 刷新数据列表
        // this.getUserList()
        // this.$message.success('更新用户信息成功！')
      })
    },

    //删除用户
    removeUserById(id){
      const result = this.$confirm('确定要删除该用户？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      console.log(result)
    }

  }


}
</script>

<style scoped>

</style>