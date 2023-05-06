<template>
	<view class="content">
		<u-popup :show="show" mode="center" @close="close">
			<view class="box">
				<view class="division">
					<text class="notion">旧密码</text>
				</view>
				<u-cell>
					<input slot="value" v-model="oldPassword" placeholder="请输入旧密码"></input>
				</u-cell>
				<view class="division">
					<text class="notion">新密码</text>
				</view>
				<u-cell>
					<input slot="value" v-model="newPassword" placeholder="请输入新密码"></input>
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
				<button class="changePasswordBtn" slot="value" @click="openPopup()">*****</button>
			</u-cell>
			<u-gap height="15" bg-color="#f9f9f9"></u-gap>
			<view class="division">
				<text class="notion">绑定邮箱账号 ( 有助于账号找回 )</text>
			</view>

			<u-cell>
				<text slot="title">绑定邮箱</text>
				<input  class="right" slot="value" v-model="user.email" placeholder="请输入邮箱"></input>
			</u-cell>
		</u-cell-group>
		<button class="upButton" @click="upInfo()">确定</button>
	</view>

</template>

<script>
	export default {
		data() {
			return {
				show: false,
				account: '',
				oldPassword: '',
				newPassword: '',
				user: {
					password: '',
					email: '',
				}
			}
		},
		methods: {
			openPopup() {
				this.show = true;
			},
			changePassword() {
				//核对旧密码输入是否正确 即oldPassword ?=user.password
				//如果错误，提示旧密码不正确
				//如果正确，修改user.password=newPassword 并关闭弹窗
				uni.showToast({
					title: '修改成功',
					//将值设置为 success 或者直接不用写icon这个参数
					icon: 'success',
					//显示持续时间为 2秒
					duration: 1500
				}) 
			},
			close() {
				this.show = false
			},
			upInfo() {
				uni.showToast({
					title: '成功',
					//将值设置为 success 或者直接不用写icon这个参数
					icon: 'success',
					//显示持续时间为 2秒
					duration: 1500
				}) 
				// // 点击上传信息按钮触发的方法
				// var that = this
				// uni.request({
				// 	//api地址
				// 	url: 'http://localhost:3000/web/api/rest/user/',
				// 	header: {
				// 		'content-type': 'application/x-www-form-urlencoded'
				// 	},
				// 	method: 'POST',
				// 	data: {
				// 		// 将json数据转化成字符串格式进行上传
				// 		information: JSON.stringify(that.user)
				// 	},
				// 	success: (res) => {
				// 		console.log(res)
				// 	},
				// 	error(err) {
				// 		console.log(err)
				// 	}
				// })
			}
			
		}
	}
</script>

<style>
	.box {
		border-radius: 10px;
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
		background-color: bisque;
		border-radius: 20px;
		margin-top: 200px;
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
		background-color: bisque;
		border-radius: 20px;
		margin-top: 50px;
	}
	.right {
		text-align: right;
	}
	
</style>