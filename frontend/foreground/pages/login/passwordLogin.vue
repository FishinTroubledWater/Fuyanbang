<template>
  <view class="login-page">
    <view class="title">欢迎登录</view>
    <form class="form" @submit.prevent="login">
      <view class="form-item">
        <label for="email">邮箱：</label>
        <input type="text" id="email" v-model="email">
      </view>
      <view class="form-item">
        <label for="password">密码：</label>
        <input type="password" id="password" v-model="password">
      </view>
      <view class="handoff">
        <text  style="margin-left: 60rpx;" @click="toRegister()">立即注册</text>
        <text style="margin-left: 100rpx;" @click="toResetPassword()">忘记密码</text>
      </view>
      <view class="button">
        <button type="submit" class="loginButton" @click="login()">登录</button>
      </view>
    </form>
  </view>
</template>

<script>
  export default {
    data() {
      return {
        email: "",
        password: "",
        stateCode: "",
        userId: "",
      };
    },
    computed: {
      isFormValid() {
        const valid = Object.keys(this.rules).every((key) => {
          return this.rules[key].every((rule) => {
            if (rule.required) {
              return !!this[key];
            }
          });
        });
        return valid;
      },
    },
    methods: {
      login() {
        // 在这里添加登录逻辑
        if (!this.validateEmail()) {
          uni.showToast({
            title: '邮箱格式错误',
            icon: 'none'
          });
          return;
        }
        if (!this.validatePassword()) {
          uni.showToast({
            title: '密码不得少于6位',
            icon: 'none'
          });
          return;
        }

        uni.$u.http.post('/v1/frontend/passwordLogin', {
          account: this.email,
          password: this.password
        }).then(res => {
          this.userId = res.data.data.User.ID;
          uni.setStorage({
            key: 'userId', // 存储值的名称
            data: this.userId, //   将要存储的数据
            success: res => {
              console.log(res);
              console.log("已成功发送userId");
            }
          });
          uni.showToast({
            title: '登陆成功',
            //将值设置为 success 或者直接不用写icon这个参数
            icon: 'success',
            //显示持续时间为 1.5秒
            // duration: 1000
          });

          setTimeout(() => {
            this.toHome();
          }, 1000);
        }).catch(err => {
          console.log(err);
          var str = err.data.msg
          if (str.includes("密码错误")) {
            uni.showToast({
              title: '密码错误',
              icon: 'none'
            });
            return;
          } else if (str.includes("账户不存在")) {
            uni.showToast({
              title: '账户不存在',
              icon: 'none'
            });
            return;
          }
          // this.loginError = true; // 设置登录错误标志
        });

      },
      toHome() {
        uni.switchTab({
          url: '../home/home'
        })
      },
      toTest() {
        uni.navigateTo({
          url: './test'
        })
      },
      toRegister() {
        uni.navigateTo({
          url: './register'
        })
      },
      toResetPassword() {
        uni.navigateTo({
          url: './resetPassword1'
        })
      },
      validateEmail() {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(this.email);
      },
      validatePassword() {
        return this.password.length >= 6;
      },
    },
    watch: {
      email: {
        handler: 'validateEmail',
        immediate: false,
      },
      password: {
        handler: 'validatePassword',
        immediate: false,
      },
    },
  };
</script>

<style>
  page {
    background-color: #fafafa;
  }

  .login-page {
    margin-top: 180rpx;
    display: flex;
    flex-direction: column;
    align-items: center;
    /* justify-content: center; */
    height: 800rpx;
  }

  .loginButton {
    width: 70%;
    height: 50px;
    background-image: linear-gradient(112deg, #08507880, #85d8ce);
    border-radius: 20px;
    margin-top: 16rpx;
    color: #fff;
    font-size: 36rpx;
    border: none;
    box-shadow: 5rpx 10rpx 5rpx #bfbfbf;
  }

  .title {
    font-size: 50rpx;
    font-weight: bold;
    margin-bottom: 30rpx;
  }

  .form {
    display: flex;
    flex-direction: column;
    align-items: center;
    /* justify-content: center; */
    padding: 70rpx;
    height: 750rpx;
    background-color: #f2f2f2;
    border-radius: 24rpx;
    box-shadow: 12rpx 12rpx 10rpx #bfbfbf;
  }

  .form-item {
    display: flex;
    flex-direction: column;
    margin-bottom: 20rpx;

  }

  .button {
    flex: 1 0 auto;
    display: flex;
    flex-direction: column;
    align-items: center;
    align-self: flex-end;
    padding: 50rpx 0;
    /* margin-bottom: 50%; */
    /* justify-content: center; */
  }

  .handoff {
    display: flex;
    flex-direction: row;
    margin-top: 50rpx;
    margin-left: -10rpx;
  }

  .resetPassword {
    /* margin-left: auto; */
  }

  .error {
    color: #FF5252;
    font-size: 12px;
    margin-top: 5px;
  }

  label {
    font-weight: bold;
    margin-bottom: 10rpx;
  }

  input {
    font-size: large;
    padding: 20rpx;
    border-radius: 5rpx;
    border: 1rpx solid #ccc;
    /* width: 100%; */
    height: auto;
    box-sizing: border-box;
  }

  button {
    background-color: #4CAF50;
    border: none;
    color: white;
    padding: 10rpx;
    width: 400rpx;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 40rpx;
    border-radius: 30rpx;
    cursor: pointer;
  }
</style>