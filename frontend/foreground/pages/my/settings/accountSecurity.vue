<template>
	<view class="backGround">
		<view class="infoCard">
			<view class="content">
				<u-popup :show="show" mode="center" @close="close">
					<view class="box">
						<view class="division">
							<text class="notion">旧密码</text>
						</view>
						<u-cell>
							<input slot="value" password="true" v-model="oldPassword" placeholder="请输入旧密码"></input>
						</u-cell>
						<view class="division">
							<text class="notion">新密码</text>
						</view>
						<u-cell>
							<input slot="value" password="true" v-model="newPassword" placeholder="请输入新密码"></input>
						</u-cell>
						<button class="Btn" @click="changePassword()">确定</button>
					</view>
				</u-popup>
				<u-cell-group>
					<view class="division">
						<text class="notion">密码设置</text>
					</view>
					<u-cell>
						<text slot="title">重置密码</text>
						<button class="changePasswordBtn" slot="value"
							@click="openPopup()">*****</button>
					</u-cell>
					<u-gap height="15" bg-color="#f9f9f9"></u-gap>
					<view class="division">
						<text class="notion">绑定邮箱账号 ( 有助于账号找回 )</text>
					</view>
					<u-cell>
						<text slot="title">绑定邮箱</text>
						<input class="right" slot="value" v-model="user.account"
							placeholder="请输入邮箱">{{user.account}}</input>
					</u-cell>
				</u-cell-group>
				<button class="upButton" @click="upInfo()">确定</button>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				show: false,
				id: '',
				oldPassword: '',
				newPassword: '',
				user: {
					password: '',
					account: '',
				}
			}
		},
		onShow() {
			console.log("执行onShow()函数")
			this.refresh();
		},
		mounted() {
			console.log("执行onLoad（）1");
			this.refresh();
			console.log("执行onLoad（）2");
		},
		methods: {
			refresh() {
				uni.getStorage({
					key: 'userId', // 储存在本地的变量名
					success: res => {
						// 成功后的回调
						// console.log(res.data);   // hello  这里可做赋值的操作
						this.id = res.data;
						console.log(this.id)
					}
				})
				uni.$u.http.get('v1/frontend/user/basicUserInfo?id=' + this.id, {}).then(res => {
					console.log("获取数据成功！！！！")
					console.log(res.data.data);
					this.user.password = res.data.data.user.Password;
					this.user.account = res.data.data.user.Account;
				}).catch(err => {
					console.log("获取数据失败！！！！")

				})

			},
			openPopup() {
				this.show = true;
			},
			changePassword() {
				//核对旧密码输入是否正确 即oldPassword ?=user.password
				if (this.oldPassword == this.user.password) //如果正确，修改user.password=newPassword 并关闭弹窗
				{
					uni.showToast({
						title: '修改成功请点击确定键',
						//将值设置为 success 或者直接不用写icon这个参数
						icon: 'success',
						//显示持续时间为 2秒
						duration: 1500
					})
					this.show = false;
					this.user.password = this.newPassword;
					console.log(this.user.password);
				} else { //如果错误，提示旧密码不正确
					uni.showToast({
						title: '旧密码错误',
						//将值设置为 success 或者直接不用写icon这个参数
						icon: 'error',
						//显示持续时间为 2秒
						duration: 1500
					})
				}
			},
			close() {
				this.show = false
			},
			upInfo() {
				uni.$u.http.post('/v1/frontend/user/resetPassword', {
					account: this.user.account,
					password: this.user.password,
				}).then(res => {
					console.log(res);
					uni.showToast({
						title: '成功',
						//将值设置为 success 或者直接不用写icon这个参数
						icon: 'success',
						//显示持续时间为 1.5秒
						duration: 1500
					})
					this.timer = setInterval(() => {
						//TODO 
						uni.navigateBack({
							delta: 1, //返回层数，2则上上页
						})
					}, 1500);
				}).catch(err => {
					console.log("失败了。。。");
				})

			},
			onUnload: function() {
				if (this.timer) { //在页面卸载时清除定时器有时会清除不了，可在页面跳转时清除
					clearInterval(this.timer);
					this.timer = null;
				}
			}

		}
	}
</script>

<style>
	.backGround {
		width: 100%;
		height: 1600rpx;
		background-color: #fafafa
	}

	.box {
		border-radius: 20rpx;
		height: 500rpx;
	}

	.division {
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		height: 30px;
	}

	.notion {
		margin-left: 20px;
		width: 80%;
		font-size: 13px;
		color: darkgrey;
	}

	.upButton {
		
		width: 80%;
		height: 50px;
		background-image: linear-gradient(112deg, #08507880, #85d8ce);
		border-radius: 20px;
		margin-top: 200rpx;
		/* margin-bottom: 50rpx; */
		color: #fff;
		border: none;
		box-shadow: 5rpx 10rpx 5rpx #bfbfbf;
	}

	.changePasswordBtn {
		background-color: white;
	}

	.changePasswordBtn::after {
		border: none;
	}

	.Btn {
		width: 80%;
		height: 50px;
		background-image: linear-gradient(112deg, #08507880, #85d8ce);
		border-radius: 20px;
		margin-top: 50rpx;
		color: #fff;
		border: none;
		box-shadow: 5rpx 10rpx 5rpx #bfbfbf;
	}

	.right {
		text-align: right;
	}
	
	.u-cell {
		padding-left: 20px;
		padding-right: 20px;
	}
	
	.infoCard {
		width: auto;
		height: 400rpx;
		margin: 25rpx;
		margin-top: 60rpx;
		margin-bottom: 60rpx;
		background-color: #fff;
		box-shadow: 12rpx 12rpx 10rpx #bfbfbf;
		border-radius: 20rpx;
	}
</style>