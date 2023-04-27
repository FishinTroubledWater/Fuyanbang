<template>
	<view class="login-page">
		<view class="title">身份验证</view>
		<form class="form" @submit.prevent="login">
			<view class="form-item">
				<label for="email">邮箱：</label>
				<input type="text" id="email" v-model="email">
			</view>
			<!-- <view class="sendCode">
				<u-code :seconds="seconds" ref="uCode" @change="">后重新获取</u-code>
				<u-button type="submit" @tap="">{{tips}}</u-button>
			</view> -->
			<view class="sendCode">
				<u-code :seconds="seconds" ref="uCode" @change="codeChange">后重新获取</u-code>
				<u-button @tap="getCode" >{{tips}}</u-button>
			</view>
			<view class="form-item" style="margin-top: 20rpx;">
				<label for="code">请输入验证码：</label>
				<u-code-input mode="line" :space="20" :maxlength="4" hairline></u-code-input>
			</view>
			<view class="button">
				<button type="submit" @click="toResetPassword2()">下一步</button>
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
				code: "",
				tips: '',
				// refCode: null,
				seconds: 60,
			};
		},
		methods: {
			toSetPassword() {
				uni.navigateTo({
					url: './setPassword?email='+this.email.toString(),
				})
			},
			toResetPassword2() {
				uni.navigateTo({
					url: './resetPassword2?email='+this.email.toString(),
				})
			},
			codeChange(text) {
				this.tips = text;
			},
			getCode() {
				if (this.$refs.uCode.canGetCode) {
					// 模拟向后端请求验证码
					uni.showLoading({
						title: '正在获取验证码'
					})
					setTimeout(() => {
						uni.hideLoading();
						// 这里此提示会被this.start()方法中的提示覆盖
						uni.$u.toast('验证码已发送');
						// 通知验证码组件内部开始倒计时
						this.$refs.uCode.start();
					}, 2000);
				} else {
					uni.$u.toast('倒计时结束后再发送');
				}
			},
		}
	};
</script>

<style>
	.login-page {
		display: flex;
		flex-direction: column;
		align-items: center;
		/* justify-content: center; */
		height: 80rpx;
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
		border-radius: 20rpx;
		box-shadow: 0 0 30rpx rgba(0, 0, 0, 0.2);
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