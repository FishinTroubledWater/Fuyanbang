<template>
	<view class="feedback-body">

		<text class="text-black">问题反馈和意见建议
			<text style="color: red;">*</text></text>
		<textarea placeholder="请描述您遇到的问题或对本产品的建议..." v-model="sendData.content" class="feedback-textare"
			maxlength="-1" />

<!-- 
		<text class="text-black">联系方式<text class="text-grey">(必填)</text> </text>
		<text style="color: red;">*</text></text>
		<input class="feedback-input" v-model="sendData.account" placeholder="请输入您的邮箱" /> -->
		<view class="btn">
			<button :disabled="!sendData.content" @click="upFeedback()" class="Btn">
				提交
			</button>
		</view>

	</view>
</template>

<script>
	export default {
		data() {
			return {
				userID: '',
				sendData: {
					content: '', //反馈内容
					// account: '', //联系邮箱
					state: '0', //默认刚提交的反馈状态为‘未处理’
					time: '', //当前设备当前时间
				},
				btnLoading: false
			};
		},
		mounted() {
			// console.log("执行onLoad（）");
			uni.getStorage({
				key: 'userId', // 储存在本地的变量名
				success: res => {
					// 成功后的回调
					// console.log(res.data);   // hello  这里可做赋值的操作
					this.userID = res.data;
					console.log(this.userID)
				}
			})
		},
		methods: {
			getTime:function(){
						var date = new Date(),
						year = date.getFullYear(),
						month = date.getMonth() + 1,
						day = date.getDate(),
						hour = date.getHours() < 10 ? "0" + date.getHours() : date.getHours(),
						minute = date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes(),
						second = date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds();
						month >= 1 && month <= 9 ? (month = "0" + month) : "";
						day >= 0 && day <= 9 ? (day = "0" + day) : "";
						var timer = year + '-' + month + '-' + day + ' ' + hour + ':' + minute + ':' + second;
						return timer;
					},
			upFeedback() {
				this.sendData.time=this.getTime();
				uni.$u.http.post('/v1/frontend/user/sendFeedback', {
					userID: this.userID,
					content: this.sendData.content,
					state: this.sendData.state,
					account: this.sendData.account,
					time: this.sendData.time,

				}).then(res => {
					console.log(res);
					uni.showToast({
						title: '发表成功',
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
					console.log("失败了........")
					uni.showToast({
						title: '发表失败',
						//将值设置为 success 或者直接不用写icon这个参数
						icon: 'error',
						//显示持续时间为 1.5秒
						duration: 1500
					})
				})
			}
		},
	}
</script>

<style>
	.text-black {
		color: #303133;
		font-size: 32rpx;
	}

	.text-grey {
		color: #BCBCBC;
		font-size: 24rpx;
		margin-left: 15rpx;
	}

	.feedback-quick {
		padding-right: 10rpx;
		color: #606266;
		font-size: 32rpx;
	}

	.feedback-body {
		padding: 30rpx;
	}

	.feedback-textare {
		margin-top: 30rpx;
		margin-bottom: 30rpx;
		height: 266rpx;
		color: #303133;
		font-size: 28rpx;
		line-height: 2em;
		width: 100%;
		box-sizing: border-box;
		padding: 20rpx 30rpx;
		border-radius: 20rpx;
		background-color: #F5F6F8;
	}

	.feedback-input {
		font-size: 28rpx;
		color: #303133;
		background-color: #F5F6F8;
		border-radius: 20rpx;
		height: 100rpx;
		min-height: 100rpx;
		padding: 0 30rpx;
		margin-top: 30rpx;
		margin-bottom: 60rpx;
	}



	.btn-submit {
		border-radius: 20rpx;
		color: #FFFFFF;
		margin-top: 100rpx;
		background-color: #007AFF;
		margin-bottom: 70rpx;
	}

	.image-title {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		color: #606266;

	}

	.filepicker {
		margin-top: 30rpx;
		margin-bottom: 30rpx;
	}

	.btn {
		position: fixed;
		bottom: 0;
		left: 0;
		right: 0;
		margin: 30rpx 30rpx 60rpx 30rpx;
	}
	.Btn{
		width: 80%;
		height: 50px;
		background-color: bisque;
		border-radius: 20px;
	}
</style>