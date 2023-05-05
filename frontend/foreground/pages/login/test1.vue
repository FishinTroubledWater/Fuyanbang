<template>
    <view>

        <view class="input-wrap">
            <view v-if="isShow">
                <input type="text" placeholder="手机号/用户名" v-model="userName">
                <input type="text" placeholder="密码" v-model="userPwd">
            </view>
            <view v-if="!isShow" class="login-phone">
                <input type="text" placeholder="手机号" v-model="userNamePhone">
                <input type="text" placeholder="验证码" v-model="userPwdPhone">
                <view class="phone-code" @tap="getPhonecode">{{codeBtn.codeText}}</view>
            </view>
            <view>
                <view class="input-switch" @click="loginSwitch">{{loginWay}}</view>
                <view class="input-forget">忘记密码</view>
            </view>
        </view>
        <view class="login" @tap="submit" v-if="isShow">
            登录
        </view>
        <view class="login" @tap="submitPhone" v-if="!isShow">
            登录
        </view>
        <view class="register">
            注册
        </view>
        
    </view>
</template>

<script>
    export default {
        data() {
            return {
                isShow: true,
                loginWay: "短信验证码登陆",
                //用户输入内容
                userName: "",
                userPwd: "",
                userNamePhone: "",
                userPwdPhone: "",
                //验证规则
                rules: {
                    userName: {
                        rule:/\S/,
                        msg: "账号不能为空"
                    },
                    userPwd: {
                        rule: /^[0-9a-zA-Z]{6,16}$/,
                        msg: "密码应该为6-16位"
                    },
                    userNamePhone: {
                        rule: /^1[3456789]\d{9}$/,
                        msg: "手机号格式错误"
                    },
                    userPwdPhone: {
                        rule: /^[0-9]{6}$/,
                        msg: "请输入6位数字验证码"
                    }
                },
                //验证码按钮
                codeBtn: {
                    codeTime: 60,
                    codeText: "获取验证码",
                    codeStatus: true
                }
            }
        },
        methods:{
            //切换登录方式
            loginSwitch() {
                this.isShow = !this.isShow
                if(this.isShow) {
                    this.loginWay = "短信验证码登陆"
                }else{
                    this.loginWay = "账号密码登陆"
                }
            },
            //账号密码点击登录
            submit() {
                if(!this.validate('userName')) return;
                if(!this.validate("userPwd"))  return;
                uni.showLoading({
                    title:"登录中..."
                });
                //向服务器发送登陆请求
                setTimeout(()=>{
                    //隐藏登录状态
                    uni.hideLoading();
                    uni.navigateBack({
                        delta:1
                    })
                },2000)
            },
            //判断验证是否符合要求，合法性校验
            validate(key){
                let bool=true;
                if(!this.rules[key].rule.test(this[key])){
                    //提示信息
                    uni.showToast({
                        title:this.rules[key].msg,
                    })
                    //取反
                    bool=false;
                    return false;
                }
                return bool;
            },
            //手机验证码登录    
            submitPhone() {
                if(!this.validate('userNamePhone')) return;
                if(!this.validate("userPwdPhone"))  return;
                //向服务器发送登陆请求
                uni.request({
                    url: 'http://.....',//接口地址
                    data: {
                        user_MobilePhone: this.userNamePhone,
                        user_Password: this.userPwdPhone
                    },
                    method: "POST",
                    success: (res) => {
                        console.log("之前的数据状态" + "账号：" + this.userNamePhone + "密码：" + this.userPwdPhone)
                        let list = JSON(stringify(res.data))
                        console.log("数据状态:" + list)
                        if(list == "[]"){
                            uni.showToast({
                                icon: 'none',
                                title: '用户名密码错误'
                            })
                            return
                        }
                        uni.showToast({
                            icon: "none",
                            title: "登陆成功"
                        })
                        uni.switchTab({
                            url: "../course/index"
                        })
                    },
                    fail: () => {
                        uni.showToast({
                            icon: "none",
                            title: "网络异常，请稍后再试"
                        })
                    }
                })
            },
            //获取验证码按钮点击计时事件
            getPhonecode() {
                if(this.validate('userNamePhone') && this.codeBtn.codeStatus) {
                    this.codeBtn.codeStatus = false
                    let timerId = setInterval(() => {
                        let codetime = this.codeBtn.codeTime
                        codetime--
                        this.codeBtn.codeTime = codetime
                        this.codeBtn.codeText = codetime + "s"
                        if(codetime < 1) {
                            clearInterval(timerId)
                            this.codeBtn.codeText = "重新获取"
                            this.codeBtn.codeTime = 60
                            this.codeBtn.codeStatus = true
                        }
                    },1000)
                }
            },
            //微信登录
            // wechatLogin() {
                
            // }
        }
    }
</script>

<style>
</style>