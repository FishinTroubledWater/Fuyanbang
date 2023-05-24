<template>
  <div class="login">
    <el-form ref="loginForm" :model="loginForm" :rules="loginRules" class="login-form">
      <h3 class="title">福研帮管理系统</h3>
      <el-form-item prop="account">
        <el-input
            v-model="loginForm.account"
            type="text"
            auto-complete="off"
            placeholder="账号"
            prefix-icon="el-icon-user-solid"
        >
        </el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input
            v-model="loginForm.password"
            type="password"
            auto-complete="off"
            placeholder="密码"
            prefix-icon="el-icon-lock"
        >
        </el-input>
      </el-form-item>
      <el-row :gutter="20">
      <el-col :span="6" :offset="2"><el-checkbox v-model="loginForm.rememberMe" style="margin:0 0 25px 0;">记住密码</el-checkbox></el-col>
      </el-row>

      <el-form-item style="width:100%;">
        <el-button
            :loading="loading"
            size="medium"
            type="primary"
            style="width:100%; background-image:linear-gradient(83deg,#ffd3a5,#fd6585)"
            @click.native.prevent="handleLogin"
            round
        >
          <span v-if="!loading">登 录</span>
          <span v-else>登 录 中...</span>
        </el-button>
      </el-form-item>
    </el-form>
    <!--  底部  -->
    <div class="el-login-footer">
      <span>Copyright © 2023-2023 福研帮 版权所有.</span>
    </div>
  </div>
</template>

<script>

export default {
  name: "Login",
  data() {
    return {
      //表单数据绑定对象
      loginForm: {
        account: "admin",
        password: "admin123",
      },
      //表单验证规则
      loginRules: {
        account: [
          {required: true, trigger: "blur", message: "请输入您的账号"},
          {min: 3, max: 10, trigger: "blur", message: "长度在3到10个字符"}
        ],
        password: [
          {required: true, trigger: "blur", message: "请输入您的密码"},
          {min: 6, max: 20, trigger: "blur", message: "长度在6到20个字符"}
        ],
      },
      loading: false,

    };
  },
  methods: {
      handleLogin() {
        this.loading=true;
        this.$refs.loginForm.validate(async valid =>{
          try {
            const response = await this.axios.post('/login', this.loginForm);
            const res = response.data;
            if (res.code === 200) {
              this.$message.success('登录成功');
              window.sessionStorage.setItem('token', res.token);
              window.sessionStorage.setItem('account', res.account);
              window.sessionStorage.setItem('id', res.id);
              this.$router.push({ path: '/home' });
            } else {
              this.loading=false;
              this.$message.error('登录失败！');
            }
          } catch (error) {
            // 处理请求错误
            console.error(error);
            this.loading=false;
            this.$message.error('登录失败！');
          }
        });
    }
  }
}

</script>
<style>
.login {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  background-image: url("../assets/images/login-background.jpg");
  background-size: cover;
}
.title {
  margin: 0 auto 30px auto;
  text-align: center;
  font-weight: bold;
  color: #f88181;
}
.login-form {
  border-radius: 30px;
  background: rgba(255, 255, 255, 0.80);
  width: 400px;
  padding: 25px 25px 5px 25px;
}
.el-login-footer {
  height: 40px;
  line-height: 40px;
  position: fixed;
  bottom: 0;
  width: 100%;
  text-align: center;
  color: #fff;
  font-family: Arial;
  font-size: 12px;
  letter-spacing: 1px;
}
</style>

