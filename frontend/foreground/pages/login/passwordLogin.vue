<!-- <template>
	<view class="login-page">
		<view class="title">欢迎登录</view>
		<form class="form" @submit.prevent="login">
			<view class="form-item">
				<label for="email">邮箱：</label>
				<input type="text" id="email" v-model.lazy="email" :class="{ invalid: !validEmail }">
				<p v-if="!validEmail && email !== ''" class="error">请输入有效的电子邮件地址</p>
			</view>
			<view class="form-item">
				<label for="password">密码：</label>
				<input type="password" id="password" v-model.lazy="password" :class="{ invalid: !validPassword }">
				<p v-if="!validPassword && password !== ''" class="error">密码必须至少为6个字符</p>
			</view>
			<view class="handoff">
				<text @click="toRegister()">立即注册</text>
				<text class="resetPassword" @click="toResetPassword()">忘记密码</text>
			</view>
			<view class="button">
				<button type="submit" @click="login()" :disabled="!validEmail || !validPassword">登录</button>
			</view>
		</form>
	</view>
</template> -->
<template>
	<view class="login-page">
		<view class="title">欢迎登录</view>
		<form class="form" @submit.prevent="login" :rules="rules" ref="form1">
			<view class="form-item" ref="item1">
				<label for="email">邮箱：</label>
				<input type="text" id="email" v-model="email">
			</view>
			<view class="form-item" ref="item1">
				<label for="password">密码：</label>
				<input type="password" id="password" v-model="password">
			</view>
			<view class="handoff">
				<text @click="toRegister()">立即注册</text>
				<text class="resetPassword" @click="toResetPassword()">忘记密码</text>
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
				stateCode: "",
				userId: "",
				rules: {
					email: [
						{
							type: 'string',
							required: true,
							message: '请填写邮箱',
							trigger: ['blur', 'change']
						},
						{
							pattern: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
							// 正则检验前先将值转为字符串
							transform(value) {
								return String(value);
							},
							message: '只能包含字母或数字',
							trigger: ['blur', 'change']
						},
						{
							min: 6,
							max: 8,
							message: '长度在6-8个字符之间',
							trigger: ['blur', 'change']
						},
					],
					'password': [{
							type: 'string',
							required: true,
							message: '请填写密码',
							trigger: ['blur', 'change']
						},
						{
							min: 6,
							max: 8,
							message: '长度在6-8个字符之间',
							trigger: ['blur', 'change']
						},
					]
				},
			};
		},

		methods: {
			login() {
				// 在这里添加登录逻辑
				console.log("邮箱：" + this.email);
				console.log("密码：" + this.password);
				uni.$u.http.post('http://localhost:8088/v1/frontend/passwordLogin', {
					account: this.email,
					password: this.password
				}).then(res => {
					console.log(res);
					// this.stateCode = res.statusCode;
					this.userId = res.data.data.User.ID;
					uni.setStorage({
						key: 'userId', // 存储值的名称
						data: this.userId, //   将要存储的数据
						success: res => {
							// 成功后的回调
							console.log(res);
						}
					});
					this.toHome();
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
				this.validEmail = emailRegex.test(this.email);
			},
			validatePassword() {
				this.validPassword = this.password.length >= 6;
			},
		},
		// watch: {
		// 	email: {
		// 		handler: 'validateEmail',
		// 		immediate: false,
		// 	},
		// 	password: {
		// 		handler: 'validatePassword',
		// 		immediate: false,
		// 	},
		// },
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