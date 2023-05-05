<template>
	<view class="login-page">
		<view class="title">欢迎登录</view>
		<form class="form" @submit.prevent="login">
			<view class="form-item">
				<label for="username">请输入短信验证码：</label>
				<u-code-input mode="line" :space="15" :maxlength="4" hairline></u-code-input>
			</view>
			<view class="sendCode">
				<u-code :seconds="seconds" @end="end" @start="start" ref="uCode" @change="codeChange">重新获取</u-code>
				<u--text :text="tips" @tap="getCode"></u--text>
			</view>
			<view class="button">
				<button type="submit" @click="toHome()">下一步</button>
			</view>


		</form>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				username: "",
				password: "",
				tips: '',
				// refCode: null,
				seconds: 10,
			};
		},
		methods: {
			login() {
				// 在这里添加登录逻辑
				console.log("用户名：" + this.username);
				console.log("密码：" + this.password);
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
			end() {
				uni.$u.toast('倒计时结束');
			},
			start() {
				uni.$u.toast('倒计时开始');
			},
			toHome() {
				uni.switchTab({
					url: '../home/home'
				})
				console.log("123456");
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
		height: 80vh;
	}

	.title {
		font-size: 24px;
		font-weight: bold;
		margin-bottom: 20px;
	}

	.form {
		display: flex;
		flex-direction: column;
		align-items: center;
		/* justify-content: center; */
		/* width: 70%; */
		padding: 50px;
		height: 50vh;
		background-color: #f2f2f2;
		border-radius: 20px;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
	}

	.form-item {
		display: flex;
		flex-direction: column;
		margin-bottom: 10px;

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
		/* align-self: flex-end; */
		padding: 20px 0;
		/* margin-bottom: 50%; */
		/* justify-content: center; */
	}

	.handoff {
		display: flex;
		flex-direction: column;
		/* margin-top: 30px; */
		margin-left: -10px;
	}

	.resetPassword {
		margin-left: auto;
	}

	text {
		font-size: 1%;
	}

	label {
		font-weight: bold;
		margin-bottom: 5px;
	}

	input {
		font-size: large;
		padding: 10px;
		border-radius: 5px;
		border: 1px solid #ccc;
		width: 100%;
		height: auto;
		box-sizing: border-box;
	}

	button {
		background-color: #4CAF50;
		border: none;
		color: white;
		padding: 10px;
		width: 80%;
		text-align: center;
		text-decoration: none;
		display: inline-block;
		font-size: 16px;
		border-radius: 5px;
		cursor: pointer;
	}
</style>