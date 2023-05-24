<template>
	<view class="page">


		<view v-for="(item,index) in feedbacks" :key="index" @click="feedbackClick(item)">
			<view class="list-box">
				<view class="text-title">{{item.content}}</view>
				<view v-if="item.state=='0'" class="text-tips">您的反馈我们已经收到，我们会尽快处理。感谢您的支持。</view>
				<view v-if="item.state=='1'" class="text-tips">您的反馈我们已经收到，我们会尽快处理。感谢您的支持。</view>
				<view class="text-time">{{item.time}}</view>
				<view :class="item.state=='0'?'state-grey':'state-green'  ">
				<text v-if="item.state=='0'">未处理</text>
				<text v-else>已处理</text>
				
				</view>
			</view>
		</view>

	</view>
</template>

<script>
	export default {
		data() {
			return {
				userID:'',
				feedbacks: [],
			}
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
			uni.$u.http.get('/v1/frontend/user/myFeedback/' + this.userID, {
			}).then(res => {
				console.log(res.data.data);
				this.feedbacks=res.data.data;
			}).catch(err => {

			})
			// console.log("执行onLoad（）");
		},
		onLoad() {
			this.getUserFeedback();
		},
		methods: {
			//获取用户意见反馈列表
			getUserFeedback() {
				//演示数据  实际通过接口调用获得
				this.feedbacks = [{
						"time": "2022-03-07 11:31:51",
						"content": "界面显示错乱",
						"state": "1",
						"stateName": "已处理",
					},
					{
						"time": "2022-03-07 11:31:51",
						"content": "界面显示错乱",
						"state": "0",
						"stateName": "未处理",
					},
					{
						"time": "2022-03-07 11:31:51",
						"content": "界面显示错乱",
						"state": "0",
						"stateName": "未处理",
					},
					{
						"time": "2022-03-07 11:31:51",
						"content": "界面显示错乱",
						"state": "1",
						"stateName": "已处理",
					}
				];
			},
			feedbackClick(item) {

				uni.navigateTo({
					url: '/pages/my/helpAndFeedback/feedbackDetail?' + 'feedback=' + JSON.stringify(item),
					success: res => {},
					fail: () => {},
					complete: () => {}
				});
			},
		}
	}
</script>

<style lang="scss">
	page {
		background-color: #F8F8F8;
	}

	.text-title {
		color: #303133;
		font-size: 30rpx;
		font-weight: bold;
		margin-right: 100rpx;
	}

	.text-time {
		color: #909193;
		font-size: 24rpx;
		margin-top: 24rpx;
	}

	.text-tips {
		color: #bcbcbc;
		font-size: 24rpx;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		margin-top: 22rpx;
	}

	.state-green {
		position: absolute;
		right: 0;
		top: 0;
		color: #07C160;
		font-size: 28rpx;
		padding: 30rpx;
	}

	.state-grey {
		position: absolute;
		right: 0;
		padding: 30rpx;
		top: 0;
		color: #606266;
		font-size: 28rpx;
	}



	.list-box {
		position: relative;
		background-color: #FFFFFF;
		border-radius: 10rpx;
		margin-top: 30rpx;
		margin-left: 30rpx;
		margin-right: 30rpx;
		padding: 30rpx;
	}



	.flex-col {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		justify-content: space-between;
	}
</style>