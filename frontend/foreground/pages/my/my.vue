<template>
	<view>
		<view class="content">
			<!-- 个人信息 -->
			<u-cell-group>

				<view class="cellWithRadius1">
					<u-cell isLink url="/pages/my/settings/userInfo">
						<u-image width='140rpx' height='140rpx' slot="icon" :src="user.avatarUrl"
							shape="circle"></u-image>
						<view class="box" slot="title">
							<view class="box1">
								<text style="margin-right:10px;">{{user.nickName}}</text>
								<u-tag borderColor="#00bcd4" bgColor="#00bcd4" :text="user.level" size="mini"
									shape="circle"></u-tag>
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
				</view>
				<!-- <u-gap height="15" bg-color="#f9f9f933"></u-gap> -->
				<view class="cellWithRadius">
					<u-cell iconStyle="color:#876800;" icon="edit-pen" titleStyle="margin-left: 20rpx;" title="我的创作"
						isLink url="/pages/my/posts/posts"></u-cell>
					<u-cell iconStyle="color:#faad14;" icon="rmb-circle" titleStyle="margin-left: 20rpx;" title="我的学币"
						isLink url="/pages/my/coin/coin"></u-cell>
					<u-cell iconStyle="color:#fa8c16;" icon="star-fill" titleStyle="margin-left: 20rpx;" title="我的收藏"
						isLink url="/pages/my/posts/favorites"></u-cell>
					<!-- <u-gap height="15" bg-color="#f9f9f933"></u-gap> -->
				</view>
				<view class="cellWithRadius">
					<u-cell iconStyle="color:#73d13d;" icon="question-circle" titleStyle="margin-left: 20rpx;"
						title="帮助与反馈" isLink url="/pages/my/helpAndFeedback/feedbackIndex"></u-cell>
					<u-cell iconStyle="color:#7cb305;" icon="setting" titleStyle="margin-left: 20rpx;" title="设置" isLink
						url="/pages/my/settings/setting"></u-cell>
				</view>
			</u-cell-group>

		</view>
	</view>

</template>

<script>
	export default {
		data() {
			return {
				id: "",
				user: {
					avatarUrl: '/static/my-assets/taiku.png',
					nickName: '张三',
					level: 'Lv.10',
					slogan: '一定上岸！！',
					useageDays: '50',
					college: '福州大学',
					major: '软件工程',
				}
			}
		},
		onShow() {
			this.refresh();
		},
		onPullDownRefresh() {
			setTimeout(() => {
				this.refresh();
				uni.stopPullDownRefresh();

			}, 1000)
		},
		mounted() {
			this.refresh();
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
				uni.$u.http.get('v1/frontend/user/basicUserInfo?id=' + this.id, {

				}).then(res => {
					console.log("请求数据成功！");
					console.log(res.data.data);
					this.user.avatarUrl = res.data.data.user.AvatarUrl;
					this.user.nickName = res.data.data.user.NickName;
					this.user.level = res.data.data.level;
					this.user.slogan = res.data.data.user.Slogan;
					this.user.useageDays = res.data.data.userDay;
					this.user.college = res.data.data.user.College;
					this.user.major = res.data.data.user.Major;
				}).catch(err => {
					console.log("请求数据失败！！！");
				})
			},
			gotoPage(url) {
				uni.navigateTo({
					url
				})
			},
		}
	}
</script>

<style lang="scss">
	page {
		background-color: #fafafa
	}

	.cellWithRadius1 {
		width: auto;
		height: 100%;
		margin: 25rpx;
		margin-top: 100rpx;
		background-color: #fff;
		box-shadow: 5rpx 10rpx 10rpx #bfbfbf;
		border-radius: 20rpx;
	}

	.cellWithRadius {
		width: auto;
		height: 100%;
		margin: 25rpx;
		margin-top: 35rpx;
		background-color: #fff;
		box-shadow: 5rpx 10rpx 10rpx #bfbfbf;
		border-radius: 20rpx;
	}


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
		font-size: 70rpx;
		color: #00bcd4;
	}

	.cellTitle {
		margin-left: 20rpx;
	}

	.u-cell {

		border-color: rgba(0, 0, 0, 0.1);

	}
</style>