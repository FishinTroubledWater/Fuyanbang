<template>
  <el-container style="height: 100%; border: 1px solid #eee">
    <el-aside width="200px" style="background-color: rgb(70,74,83)">
      <el-menu :default-openeds="['1', '3']" background-color=" #343c3f" text-color="#fff" default-active="1" :router="true">
        <div class="logo">
            <img src="../assets/logo.png" style="width: 40px;padding:10px">
            <span style="font-weight: bold;">福研帮管理系统</span>
        </div>
        <el-menu-item index="home" ><i class="el-icon-s-home"/> <span class="el-menu-text">首页</span></el-menu-item>
        <el-menu-item index="user"><i class="el-icon-user"/> <span class="el-menu-text">用户管理</span></el-menu-item>
        <el-menu-item index="content"><i class="el-icon-tickets"/> <span class="el-menu-text">内容管理</span></el-menu-item>
        <el-menu-item index="trends"><i class="el-icon-discover"/> <span class="el-menu-text">动态管理</span></el-menu-item>
        <el-menu-item index="recipe"><i class="el-icon-discover"/> <span class="el-menu-text">学长学姐说管理</span></el-menu-item>
      </el-menu>
    </el-aside>

    <!--    主体头部-->
    <el-container>
      <el-header style="text-align: right; font-size: 12px">
        <el-dropdown style="width:100px">
          <span>王小虎</span><i class="el-icon-arrow-down"></i>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item>查看</el-dropdown-item>
            <el-dropdown-item>新增</el-dropdown-item>
            <el-dropdown-item>删除</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>

      </el-header>

      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>
<script>
export default {


};


</script>

<style>
.logo {
   height:60px ;
   display: flex;
  justify-content: center;
  align-items: center;
}

.el-header {
  background-color: #B3C0D1;
  color: #333;
  line-height: 60px;
}

.el-aside {
  color: #ffffff;
}

</style>