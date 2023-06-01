<!-- 帮助反馈页面 -->
<template>
	<view>
		<view class="text-top">常见问题</view>
		<view class="bg-list">
			<view v-for="(item,index) in problemBox" :key="index" :title="item.problemTitle" class="list-title">
				<view class="text-title">{{item.problemTitle}}</view>
				<view v-for="(problemItem,problemIndex) in item.problemList" :key="problemIndex" class="list-question"
					hover-class="hover" @click="onClick(problemItem)">
					<view class="text-item">{{problemItem.title}}</view>
					<view class="line" v-if="problemIndex!=item.problemList.length-1"></view>
				</view>
			</view>
		</view>
		<view class="bg-box">
			<view class="bg-white-box">
				<image src="../../../static/my-assets/helpAndFeedback/icon-letter.png" class="image"></image>
				<view class="text-feedback" hover-class="hover" @click="toFeedbackList">我的反馈</view>

				<view class="vertical-line"></view>

				<image src="../../../static/my-assets/helpAndFeedback/icon-edit.png" class="image"></image>
				<view class="text-feedback" hover-class="hover" @click="toFeedback">我要反馈</view>
			</view>

		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				problemBox: [{
						problemTitle: '福研帮',
						problemList: [{
								title: '这个App是用来干什么的',
								content: '后疫情时代，许多福大学子开始明确了考研意向，然而考研相关的资料和信息一方面是多而繁杂又十分零散，另一方面是有些院校的信息并不公开，如何获取考研的相关信息，成为无数福大学子迫在眉睫、渴望解决的问题，而福大校内也暂时没有一款靠谱的考研信息平台。基于这样的需求，我们认为退出一款针对本校学生的考研APP是必要的，而这样一款APP无论在资源信息获取，还是本校推广都具有一定竞争优势，且由于在本校，接受用户反馈、及时调整优化也具有本地优势。所以我们决定就福大考研群体，开发一个高效率、高质量、全方位的一站式考研信息平台，致力于为广大福大考生提供优质考研服务。'
							},
							{
								title: '这个App的用户群体是什么样的',
								content: '福研帮APP主要面向就读于福州大学应届或往届选择考研的广大考生，自然也能为以福州大学作为目标院校的他校考生提供详细的福州大学考研信息,当然也能一定程度上满足国内所有有考研意向、想要简单高效获取考研信息的考生，同时还为已经上岸了的学长学姐(更偏向于本科毕业于福州大学）提供了一个分享成功经验的、能将经验转换为一定价值的窗口。总的来说，福研帮APP作用于任何具有考研意向的、需要高效获取考研信息的、渴望分享成功经验并获得一定价值的群体。'
							}

						]
					},
					{
						problemTitle: '圈子发布',
						problemList: [{
								title: '圈子发布的内容哪些人可以看到',
								content: '每一个使用该APP的正式注册用户都可以看到，除此之外还有App的开发人员和后台管理人员可以看到'
							}

						]
					},
					{
						problemTitle: '学币',
						problemList: [{
								title: '学币可以用来干嘛',
								content: '学币可以用来1.发表提问并悬赏 2.兑换热门App的会员。'
							}

						]
					}, {
						problemTitle: '其他',
						problemList: [{
								title: '如何修改用户头像',
								content: '点击头像即可选择自己喜欢的头像。',
							},
							{
								title: '如何修改登录密码',
								content: '请点击[我的] - [设置] - [账号与安全]即可修改登录密码',
							},
							{
								title: '如何退出登录',
								content: '请点击[我的] - [设置] - [退出登录]即可退出登录',
							},
							{
								title: '系统反应迟钝',
								content: '如果出现系统反应迟钝、响应时间过长的问题，可查看本机的网络状态，或重启程序后再次尝试。'
							}
						]
					}
				]
			}
		},
		methods: {
			onClick(item) {

				uni.navigateTo({
					url: '/pages/my/helpAndFeedback/helpDetail?title=' + item.title + '&content=' + item.content,
					success: res => {},
					fail: () => {},
					complete: () => {}
				});

			},
			toFeedbackList() {
				uni.navigateTo({
					url: '/pages/my/helpAndFeedback/feedbackList',
					success: res => {},
					fail: () => {},
					complete: () => {}
				});
			},
			toFeedback() {
				uni.navigateTo({
					url: '/pages/my/helpAndFeedback/feedback',
					success: res => {},
					fail: () => {},
					complete: () => {}
				});
			}
		}
	}
</script>

<style>
	page {
		background-color: #fafafa;
		height: 100%;
	}



	.bg-box {
		background-color: #F8F8F8;
		position: fixed;
		bottom: 0;
		left: 0;
		right: 0;
	}

	.bg-list {
		/* margin-bottom: 100rpx; */
		/* background-color: #FFFFFF; */
		padding: 30rpx;
		width: auto;
		height: 1000rpx;
		margin: 25rpx;
		margin-top: 60rpx;
		margin-bottom: 100rpx;
		background-color: #fff;
		box-shadow: 12rpx 12rpx 10rpx #bfbfbf;
		border-radius: 20rpx;
	}

	.bg-white-box {
		background-color: #FFFFFF;
		margin: 30rpx;
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: center;
		padding: 10rpx;
		border-radius: 20rpx;
		color: #303133;
		font-size: 32rpx;
	}

	.vertical-line {
		height: 20rpx;
		background-color: #cecece;
		width: 2rpx;
		margin-left: 30rpx;
		margin-right: 30rpx;
	}

	.line {
		width: 100%;
		height: 1rpx;
		background-color: #F5F5F5;
	}



	.text-title {
		color: #303133;
		font-size: 32rpx;
		font-weight: bold;
	}

	.text-item {
		color: #999999;
		font-size: 28rpx;
		padding: 24rpx;
	}

	.list-title {
		margin-bottom: 30rpx;
	}

	.list-question {
		color: #606266;
		font-size: 28rpx;
	}

	.hover {
		background-color: #eee;
		opacity: 0.6;
	}

	.image {
		width: 40rpx;
		height: 40rpx;
		margin-left: 20rpx;
	}

	.text-feedback {
		padding: 20rpx;
	}

	.text-top {
		margin: 30rpx;
		color: #999999;
		font-size: 28rpx;
	}
</style>
