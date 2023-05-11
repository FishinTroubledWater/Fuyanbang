<template>
  <div class="app-container">
    <el-row :gutter="20">
      <!-- 左边卡片 -->
      <el-col :span="8" :xs="24">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>个人信息</span>
          </div>
          <div>
            <div class="text-center">
              <el-avatar :src="require('../assets/images/user-profile.png')"></el-avatar>
            </div>
            <ul class="list-group list-group-striped">
              <li class="list-group-item">
                <i class="el-icon-user"></i>用户名
                <div class="pull-right">福研帮后台管理员</div>
              </li>
              <li class="list-group-item">
                <i class="el-icon-edit-outline"></i>密码
                <div class="pull-right">1989-6-4</div>
              </li>
              <li class="list-group-item">
                <i class="el-icon-reading"></i>邮箱
                <div class="pull-right">fyb@fzu.com</div>
              </li>
              <li class="list-group-item">
                <i class="el-icon-school"></i>开发团队
                <div class="pull-right">福州大学软工团队</div>
              </li>
              <li class="list-group-item">
                <i class="el-icon-edit"></i>简介
                <div class="pull-right">致力于打造福大学生的考研APP</div>
              </li>
            </ul>
          </div>
        </el-card>
      </el-col>
      <!-- 右边卡片 -->
      <el-col :span="16" :xs="24">
        <el-card>
          <div slot="header" class="clearfix">
            <span>基本资料</span>
          </div>
          <el-tabs v-model="activeTab">
            <el-tab-pane label="基本资料" name="userinfo">
              用户基本信息
            </el-tab-pane>
            <el-tab-pane label="修改密码" name="resetPwd">
              重置密码
            </el-tab-pane>
          </el-tabs>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script>
export default {
  data() {
    return {
      user: {},
      roleGroup: {},
      postGroup: {},
      activeTab: "userinfo"
    }
  },
  created() {
    this.getUser();
  },
  methods: {
    //   async getUser() {
    //     const { data: res } = await this.axios.post('user/searchByAccount', {
    //       //pageSize: this.queryInfo.pageSize,
    //       //pageNum: this.queryInfo.pageNum
    //     }, {
    //       headers: {
    //         'Authorization': window.sessionStorage.getItem("token")
    //       }
    //     });
    //     console.log(res)
    //     if (res.code !== 200) return this.$message.error('获取用户列表失败！')
    //     //this.userList = res.data.users
    //     //this.total = res.data.total
    //     console.log(this.userList);
    //   }
    // }
  }
}
</script>
<style scoped>
.text-center {
  /*flex 布局*/
  display: flex;
  /*实现垂直居中*/
  align-items: center;
  /*实现水平居中*/
  justify-content: center;

  text-align: justify;
  width: 100px;
  height: 100px;
  background: rgb(255, 255, 255);
  margin: 0 auto;
  color: rgb(0, 0, 0);
}

li {
  list-style: none;
}

.list-group-item {
  line-height: 60px;
  width: 290px;
  border-bottom: 1px solid rgb(70, 70, 70)
}

.list-group {
  justify-content: left;
}

.pull-right {
  float: right;
}
</style>