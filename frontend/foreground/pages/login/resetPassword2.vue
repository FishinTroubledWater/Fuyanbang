<template>
	<view class="login-page">
		<view class="title">修改密码</view>
		<form class="form" @submit.prevent="login">
			<view class="form-item">
				<view class="label">请输入新密码：</view>
				<input class="input" v-model="password" type="password" placeholder="请输入密码" />
			</view>
			<view class="form-item">
				<view class="label">确认密码：</view>
				<input class="input" v-model="confirmPassword" type="password" placeholder="请再次输入密码" />
			</view>
			<button class="button" type="primary" @tap="submitForm">确定</button>
		</form>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				email: '',
				password: '',
				confirmPassword: ''
			};
		},
		onLoad: function(option) { //option为object类型，会序列化上个页面传递的参数
			// console.log(option.email); //打印出上个页面传递的参数。
			this.email = option.email;
			// console.log('123'+this.email); //打印出上个页面传递的参数。
		},
		methods: {
			submitForm() {
				if (this.password.length < 6) {
					uni.showToast({
						title: '密码至少6位',
						icon: 'none'
					});
					return;
				}
				if (this.password !== this.confirmPassword) {
					uni.showToast({
						title: '两次输入的密码不一致',
						icon: 'none'
					});
					return;
				}

				// 如果两次密码一致，可以在这里进行注册逻辑
				// ...124.222.141.238
				uni.$u.http.post('/v1/frontend/user/resetPassword', {
					account: this.email,
					password: this.password,
				}).then(res => {
					console.log(res);
					uni.showToast({
						title: '密码修改成功，请登录',
						icon: 'success'
					});
					// this.stateCode = res.statusCode;
					setTimeout(() => {
						this.toPasswordLogin();
					}, 1000);

				}).catch(err => {
					console.log(err);
					var str = err.data.msg
					if (str.includes("要修改的记录不存在")) {
						uni.showToast({
							title: '邮箱输入错误，该邮箱还未被注册',
							icon: 'none'
						});
						return;
					}
				});
			},
			toPasswordLogin() {
				uni.navigateTo({
					url: './passwordLogin'
				})
			}
		}
	};
</script>


<style>
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