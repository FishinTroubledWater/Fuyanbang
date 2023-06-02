<template>
	<view class="login-page">
		<view class="title">立即注册</view>
		<form class="form" @submit.prevent="login">
			<view class="form-item">
				<label for="email">邮箱：</label>
				<input type="text" id="email" v-model="email">
				<!-- <p v-if="!validEmail && email !== ''" class="error">请输入有效的电子邮件地址</p> -->
			</view>
			<!-- <view class="sendCode">
				<u-code :seconds="seconds" ref="uCode" @change="">后重新获取</u-code>
				<u-button type="submit" @tap="">{{tips}}</u-button>
			</view> -->
			<view class="sendCode">
				<u-code :seconds="seconds" ref="uCode" @change="codeChange">后重新获取</u-code>
				<u-button style="
				width: 75%;
				height: 35px;
				border-radius: 20px;
				margin-top: 20rpx;
				margin-bottom: 20rpx;
				font-size: 28rpx;
				border: none;
				box-shadow: 5rpx 10rpx 5rpx #bfbfbf;" @tap="getCode()">{{tips}}</u-button>
			</view>
			<view class="form-item" style="margin-top: 20rpx;">
				<label style="margin-top: 30rpx;" for="code">请输入验证码：</label>
				<u-code-input mode="line" :space="20" :maxlength="4" v-model="code"></u-code-input>
			</view>


			<view class="button">
				<button class="nextBtn" type="submit" @click="toSetPassword()">下一步</button>
			</view>
		</form>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				email: "",
				code: "",
				tips: '',
				// refCode: null,
				seconds: 60,
			};
		},
		methods: {
			validateEmail() {
				const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
				return emailRegex.test(this.email);
			},
			toSetPassword() {

				if (!this.validateEmail()) {
					uni.showToast({
						title: '邮箱格式错误',
						icon: 'none'
					});
					return;
				}
				uni.$u.http.post('/v1/frontend/codeVerify', {
					account: this.email,
					code: this.code
				}).then(res => {
					// console.log("这里有吗？");/
					if (res.data.code == 200) {
						// console.log("这里呢？");
						uni.navigateTo({
							url: './setPassword?email=' + this.email.toString(),
						})
					}
				}).catch(err => {
					uni.$u.toast('验证码错误');
					// this.loginError = true; // 设置登录错误标志
				});
			},
			codeChange(text) {
				this.tips = text;
			},
			getCode() {
				if (!this.validateEmail()) {
					uni.showToast({
						title: '邮箱格式错误',
						icon: 'none'
					});
					return;
				}
				if (this.$refs.uCode.canGetCode) {
					// 模拟向后端请求验证码
					uni.showLoading({
						title: '正在获取验证码'
					})
					uni.$u.http.post('/v1/frontend/sendEmail', {
						account: this.email,
					}).then(res => {
						setTimeout(() => {
							uni.hideLoading();
							// 这里此提示会被this.start()方法中的提示覆盖
							uni.$u.toast('验证码已发送,有效期300s');
							// 通知验证码组件内部开始倒计时
							this.$refs.uCode.start();
						}, 2000);
					}).catch(err => {
						uni.$u.toast('验证码发送失败');
						// this.loginError = true; // 设置登录错误标志
					});

				} else {
					uni.$u.toast('倒计时结束后再发送');
				}
			},

		},

		watch: {
			email: {
				handler: 'validateEmail',
				immediate: false,
			},

		},
	};
</script>
``
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
		height: 800rpx;
		background-color: #f2f2f2;
		opacity: 0.8;
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
		padding: 20rpx 0;
		/* margin-bottom: 50%; */
		/* justify-content: center; */
	}

	.handoff {
		display: flex;
		flex-direction: column;
		/* align-items: center; */
		margin-top: 50rpx;
		margin-left: -10rpx;
	}

	.register {
		margin-top: 20px;
	}

	.error {
		color: #FF5252;
		font-size: 12px;
		margin-top: 5px;
	}

	.nextBtn {
		width: 70%;
		height: 50px;
		background-image: linear-gradient(112deg, #08507880, #85d8ce);
		border-radius: 20px;
		margin-top: 40rpx;
		color: #fff;
		font-size: 36rpx;
		border: none;
		box-shadow: 5rpx 10rpx 5rpx #bfbfbf;
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
		padding: 12rpx;
		width: 350rpx;
		text-align: center;
		text-decoration: none;
		display: inline-block;
		font-size: 38rpx;
		border-radius: 30rpx;
		cursor: pointer;

	}
</style>