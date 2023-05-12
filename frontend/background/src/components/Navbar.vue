<template>
  <el-container style="height: 100%; border: 1px solid #eee">
    <!--    侧边导航-->
    <el-aside width="200px" style="background-color: rgb(70,74,83)">
      <el-menu :default-openeds="['1', '3']" background-color=" #343c3f" text-color="#fff" default-active="1"
               :router="true">

        <div class="logo">
          <img src="../assets/logo.png" style="width: 40px;padding:10px">
          <span style="font-weight: bold;">福研帮管理系统</span>
        </div>

        <el-menu-item index="home"><i class="el-icon-s-home"/> <span class="el-menu-text">首页</span></el-menu-item>
        <el-menu-item index="user"><i class="el-icon-user"/> <span class="el-menu-text">用户管理</span></el-menu-item>
        <el-menu-item index="post"><i class="el-icon-tickets"/> <span class="el-menu-text">帖子管理</span></el-menu-item>
        <el-menu-item index="news"><i class="el-icon-discover"/> <span class="el-menu-text">资讯管理</span></el-menu-item>
        <el-menu-item index="recipe"><i class="el-icon-s-promotion"/> <span class="el-menu-text">学长学姐说管理</span>
        </el-menu-item>
        <!--        <el-menu-item index="comment"><i class="el-icon-s-comment"/> <span class="el-menu-text">评论管理</span></el-menu-item>-->
        <el-menu-item index="feedback"><i class="el-icon-message-solid"/> <span class="el-menu-text">反馈中心</span>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <!--    主体头部-->
    <el-container>
      <el-header style="text-align: right; font-size: 12px">
        <el-dropdown trigger="click">
          <div class="top_right">
            <el-avatar :size="'large'" :src="avatar" :fit="'cover'"
                       style="margin-right: 10px"></el-avatar>
            <span class="el-dropdown-link">{{ account }}<i class="el-icon-caret-bottom el-icon--right"></i></span></div>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item class="clearfix" @click.native="toAbout">个人页面</el-dropdown-item>
            <el-dropdown-item class="clearfix" @click.native="logout">退出</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </el-header>
      <el-main class="el-main">
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>
<script>
import avatar from "../assets/images/profile.jpg"

export default {
  data() {
    return {
      avatar: avatar,
      account: window.sessionStorage.getItem("account"),
      id: window.sessionStorage.getItem("id"),
    }
  },
  methods: {
    logout() {
      window.sessionStorage.clear();
      this.$confirm('确定要退出登录？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$router.push('/login');
      })
    },
    toAbout() {
      this.$router.push('/about');
    }
  }
};
</script>

<style>
.top_right {
  display: flex;
  justify-content: center;
  align-items: center;
}

.logo {
  height: 60px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.el-header {
  background-color: #fdf5f5;
  color: #333;
  line-height: 60px;

}

.el-aside {
  color: #ffffff;
}

.el-dropdown-link {
  font-size: 20px;
}

.el-main {
  background-color: #dfdfdf;
}

</style>