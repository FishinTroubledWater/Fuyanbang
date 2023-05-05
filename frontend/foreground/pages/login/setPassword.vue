<template>
	<view class="login-page">
		
		<form class="form" @submit.prevent="login">
			<view class="title">设置密码</view>
			<view class="form-item">
				<view class="label">密码：</view>
				<input class="input" v-model="password" type="password" placeholder="请输入密码" />
			</view>
			<view class="form-item">
				<view class="label">确认密码：</view>
				<input class="input" v-model="confirmPassword" type="password" placeholder="请再次输入密码" />
			</view>
			<button class="button" type="primary" @tap="submitForm">注册</button>
		</form>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				password: '',
				confirmPassword: ''
			};
		},
		methods: {
			submitForm() {
				if (this.password !== this.confirmPassword) {
					uni.showToast({
						title: '两次输入的密码不一致',
						icon: 'none'
					});
					return;
				}

				// 如果两次密码一致，可以在这里进行注册逻辑
				// ...
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
		margin-top: 50rpx;
	}

	.sendCode {
		/* flex: 1 0 auto; */
		display: flex;
		flex-direction: column;
		align-items: center;
		/* align-self: flex-end; */
		padding: 20px;
		justify-content: center;
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