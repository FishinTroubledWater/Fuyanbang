<template>
	<view>
		<view class="white">
			<view class="text-black">{{feedback.content}}</view>

			<view class="text-time">{{feedback.time}}提交</view>

		</view>
		<view class="box">
			<view class="text-black">客服回复：</view>
			<view class="text-grey">您的反馈我们已经收到，我们会尽快处理，给您造成的不便敬请谅解。感谢您的反馈与支持！</view>
			<view class="text-time">{{feedback.time}}</view>

		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {

				feedback: {},

			}
		},
		onLoad(options) {
			this.feedback = JSON.parse(options.feedback);
		},

		methods: {


		}
	}
</script>

<style>
	page {
		background-color: #fafafa;
		padding-bottom: 30rpx;
	}

	.white {
		background-color: #FFFFFF;
		padding: 45rpx 40rpx;
		box-shadow: 0 12rpx 10rpx #d0d0d0;
	}

	.text-grey {
		color: #606266;
		font-size: 32rpx;
		line-height: 1.5em;
	}

	.text-time {
		color: #909193;
		font-size: 24rpx;
		text-align: end;
		margin-top: 30rpx;
	}

	.text-black {
		color: #303133;
		font-size: 32rpx;
		margin-bottom: 30rpx;
		font-weight: bold;
	}

	.box {
		background-color: #FFFFFF;
		border-radius: 30rpx;
		margin: 50rpx 30rpx;
		padding: 30rpx 30rpx;
		
		box-shadow: 12rpx 12rpx 10rpx #bfbfbf;
		border-radius: 20rpx;
	}

	.icon-normal {
		width: 25rpx;
		height: 25rpx;
		margin-right: 20rpx;
	}


	.flex-row {
		display: flex;
		flex-direction: row;
		justify-content: space-around;
		align-items: center;
		margin-top: 30rpx;
	}

	.btn-normal {
		border-radius: 40rpx;
		border: solid 1rpx #D2D2D2;
		color: #999999;
		padding: 20rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		font-size: 28rpx;
		width: 32%;
	}

	.btn-resolve {
		background-color: #E1FFFB;
		border: solid 1rpx #19CFB7;
		color: #19CFB7;

	}

	.btn-unResolve {
		border: solid 1rpx #FF5F5F;
		background-color: #FFF2F2;
		color: #FF5F5F;

	}



	.image {
		width: 200rpx;
		height: 200rpx;
		border-radius: 20rpx;
		margin: 10rpx;
	}

	.img-row {
		display: flex;
		flex-direction: row;

		flex-wrap: wrap;
		justify-content: flex-start;
		align-items: center;

	}
</style>
