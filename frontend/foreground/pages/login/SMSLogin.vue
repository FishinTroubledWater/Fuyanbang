<template>
	<view class="login-page">
		<view class="title">欢迎登录</view>
		<form class="form" @submit.prevent="login">
			<view class="form-item">
				<label for="username">手机号：</label>
				<input type="text" id="username" v-model="username">
			</view>
			<view class="button">
				<button type="submit" @click="toVerify">下一步</button>
			</view>
			<view class="handoff">
				<text @click="toPasswordLogin()">账号密码登录</text>
				<text class="register">立即注册</text>
			</view>

		</form>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				username: "",
				password: ""
			};
		},
		methods: {
			login() {
				// 在这里添加登录逻辑
				console.log("用户名：" + this.username);
				console.log("密码：" + this.password);
			},
			toVerify() {
				uni.navigateTo({
					url: './verify'
				});
			},
			toPasswordLogin() {
				uni.navigateTo({
					url: './passwordLogin'
				});
			}
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