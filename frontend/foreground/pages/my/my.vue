<template>
	<view class="content">
		<!-- 个人信息 -->
		<u-cell-group>
			<u-cell>
				<u-icon @click="gotoPage('/pages/my/setting')" slot="value" name="setting" color="#6f6f6f"
					size="22"></u-icon>
			</u-cell>
			<u-cell isLink>
				<u-image @click="changeHead()" width='140rpx' height='140rpx' slot="icon" :src="user.avatarUrl"
					shape="circle"></u-image>
				<view class="box" slot="title">
					<view class="box1">
						<text style="margin-right:10px;">{{user.nickName}}</text>
						<u-tag :text="user.level" size="mini" shape="circle"></u-tag>
					</view>
					<text style="color: darkgray;">{{user.slogan}}</text>
				</view>
			</u-cell>
			<u-cell>
				<text slot="title">{{user.college}}</text>
				<text slot="title">{{user.major}}</text>
				<text slot="value">福研帮已经陪伴了您</text>
				<text slot="value" class="days">{{user.useageDays}}</text>
				<text slot="value">天了</text>
			</u-cell>
			<u-gap height="15" bg-color="#f9f9f9"></u-gap>
			<u-cell icon="edit-pen" title="我的创作" isLink url="/pages/my/posts"></u-cell>
			<u-cell icon="rmb-circle" title="我的学币" isLink url="/pages/my/coin"></u-cell>
			<u-cell icon="star" title="我的收藏" isLink url="/pages/my/favorites"></u-cell>
			<u-gap height="15" bg-color="#f9f9f9"></u-gap>
			<u-cell icon="question-circle" title="帮助与反馈" isLink url="/pages/my/helpAndFeedback/feedbackIndex"></u-cell>
		</u-cell-group>

	</view>
</template>

<script>
	export default {
		data() {
			return {
				id: "",
				user: {
					avatarUrl: '/static/my-assets/泰裤辣.png',
					nickName: '张三',
					level: 'Lv.10',
					slogan: '一定上岸！！',
					useageDays: '50',
					college: '福州大学',
					major: '软件工程',
				}
			}
		},
		mounted() {
			console.log("执行onLoad（）");
			var _this = this;
			const on = uni.$on('login1', function(data) {
				_this.id = data.userId;
				console.log("用户id是：");
				console.log(_this.id)
			})
			console.log("执行onLoad（）");
		},
		// onLoad() {
			
		// },
		methods: {
			gotoPage(url) {
				uni.navigateTo({
					url
				})
			},
			changeHead() {
				uni.chooseImage({
					count: 1,
					success: (res) => {
						this.headImg = res.tempFilePaths[0]
					}
				});
			},
		}
	}
</script>

<style lang="scss">
	.box {
		display: flex;
		flex-direction: column;
		padding: 10px;
	}

	.box1 {
		display: flex;
		flex-direction: row;
	}

	.days {
		font-size: 25px;
		color: crimson;
	}
</style>