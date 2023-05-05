<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    用户查询-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold;"> 用户查询</div>
      <el-row :gutter="20">
        <el-col :span="24" style="margin: 5px">用户类别</el-col>
        <el-col :span="5">
          <el-button-group>
            <el-button size="mini" type="primary" autofocus>不限</el-button>
            <el-button size="mini" type="primary">普通用户</el-button>
            <el-button size="mini" type="primary">管理员</el-button>
          </el-button-group>
        </el-col>
        <el-col :span=2>
          <el-button @click="addDialogVisible = true" type="primary" size="small" round icon="el-icon-plus">添加用户
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
                       @click="handleEdit(scope.$index, scope.row)">编辑
            </el-button>
            <el-button size="mini" type="danger" icon="el-icon-delete" round
                       @click="handleDelete(scope.$index, scope.row)">删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <Pagination></Pagination>
    <!--    <Pagination :total="total" :query-info="queryInfo"></Pagination>-->
    <!--    添加用户的对话框-->
    <el-dialog title="提示" :visible.sync="addDialogVisible" width="30%">
      <!--      内容主体区域-->
      <el-form ref="addFormRef" :model="addForm" label-width="80px" :rules="addFormRules">
        <el-form-item label="用户名" prop="account">
          <el-input v-model="addForm.account"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="addDialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="addDialogVisible = false">确 定</el-button>
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

      },

      addFormRules:{
        account:[
          {required:true, message:'请输入用户名',trigger:'blur'},
          {min:3, max:10, message: '用户名的长度在3~10个字符间',trigger: 'blur'}
        ]
      },
      tableData: Array(10).fill(item),
      // 获取用户列表的参数对象
      queryInfo: {
        query: '',
        pagenum: 1,
        pagesize: 2
      },

    }
  },
  created() {
    // this.getUserList()
  },
  methods: {
    // 查询方法
    async getUserList() {
      const {data: res} = await this.axios.get('users', {
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
    }

  }


}
</script>

<style scoped>

</style>