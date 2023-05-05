<template>
  <div class="login">
    <el-form ref="loginForm" :model="loginForm" :rules="loginRules" class="login-form">
      <h3 class="title">福研帮管理系统</h3>
      <el-form-item prop="username">
        <el-input
            v-model="loginForm.username"
            type="text"
            auto-complete="off"
            placeholder="账号"
        >
        </el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input
            v-model="loginForm.password"
            type="password"
            auto-complete="off"
            placeholder="密码"
        >
        </el-input>
      </el-form-item>
      <el-row :gutter="20">
      <el-col :span="6" :offset="2"><el-checkbox v-model="loginForm.rememberMe" style="margin:0px 0px 25px 0px;">记住密码</el-checkbox></el-col>
        <el-col :span="6" :offset="10"><el-link to="/">忘记密码</el-link></el-col>
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
      loginForm: {
        username: "admin",
        password: "admin123",
        rememberMe: false,
      },
      loginRules: {
        username: [
          {required: true, trigger: "blur", message: "请输入您的账号"}
        ],
        password: [
          {required: true, trigger: "blur", message: "请输入您的密码"}
        ],
      },
      loading: false,
      // 验证码开关
      captchaEnabled: false,
      // 注册开关
      register: false,
      redirect: undefined
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$router.push("/Main");
        } else {
          this.$message.error('请输入正确用户名或密码!');
          return false;
        }
      });
    },
    handleLogin() {
      this.$router.push({path:'/home'})
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
  margin: 0px auto 30px auto;
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

