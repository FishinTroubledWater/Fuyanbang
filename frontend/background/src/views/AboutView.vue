<template>
  <div class="app-container">
    <el-row :gutter="20">
<!--       左边卡片 -->
      <el-col :span="8" :xs="24">
        <el-card class="box-card round15" >
          <div slot="header" class="clearfix">
            <span>个人信息</span>
          </div>
          <div>
            <div class="text-center">
              <el-avatar :src="require('../assets/images/user-profile.png')"></el-avatar>
            </div>
            <div class="setcenter">
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
          </div>
        </el-card>
      </el-col>
      <!-- 右边卡片 -->
      <el-col :span="16" :xs="24">
        <el-card class="round15">
          <div slot="header" class="clearfix">
            <span>基本资料</span>
          </div>
          <div class="info">
            <div class="img-center">
              <el-image :src="require('../assets/images/FYB.png')" style="width: 290px;"
                        fit="fill"></el-image>
            </div>
            <el-row style="line-height: 42px;">应用简介：由来自福州大学20级软件工程专业的学生打造</el-row>
            <el-row style="line-height: 42px;">主要功能：提供各类考研复习资料</el-row>
            <el-row style="line-height: 42px;">开发目的：为福大考研学子提供易用的考研复习APP</el-row>
            <el-row style="line-height: 42px;">涉及方向：Vue、uniapp、Go、Gin</el-row>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script>


export default {
  components: {},

  data() {
    return {
    }
  },
  methods:{

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
  width: 120px;
  height: 120px;
  background: rgb(255, 255, 255);
  margin: 0 auto;
  color: rgb(0, 0, 0);
}

.img-center {
  /*flex 布局*/
  display: flex;
  /*实现垂直居中*/
  align-items: center;
  /*实现水平居中*/
  justify-content: center;

  text-align: justify;
  margin: 0 auto;
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
  margin-left: -40px;
}

.pull-right {
  float: right;
}

.info {
  font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
  font-size: 18px;
  color: rgb(75, 75, 75);
}
</style>
