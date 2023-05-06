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
				<text @click="toRegister()">立即注册</text>
				<text class="resetPassword">忘记密码</text>
			</view>
			<view class="button">
				<button type="submit" @click="login()">登录</button>
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
				
			};
		},
		methods: {
			login() {
				// 在这里添加登录逻辑
				console.log("邮箱：" + this.email);
				console.log("密码：" + this.password);
				
			},
			toHome() {
				uni.switchTab({
					url: '../home/home'
				})
			},
			toRegister() {
				uni.navigateTo({
					url: './register'
				})
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
		padding: 50rpx 0;
		/* margin-bottom: 50%; */
		/* justify-content: center; */
	}

	.handoff {
		display: flex;
		flex-direction: row;
		margin-top: 80rpx;
		margin-left: -10rpx;
	}

	.resetPassword {
		margin-left: auto;

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
		padding: 15rpx;
		width: 400rpx;
		text-align: center;
		text-decoration: none;
		display: inline-block;
		font-size: 40rpx;
		border-radius: 30rpx;
		cursor: pointer;
	}
</style>