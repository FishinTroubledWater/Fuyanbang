<template>
	<view class="content">
		<u-notify ref="uNotify"></u-notify>
		<!-- 个人信息 -->
		<u-cell-group>
			<u-cell class="blankBox">
			</u-cell>
			<u-cell isLink url="/pages/my/settings/userInfo">
				<u-image width='140rpx' height='140rpx' slot="icon" :src="user.avatarUrl" shape="circle"></u-image>
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
			<u-cell icon="edit-pen" title="我的创作" isLink url="/pages/my/posts/posts"></u-cell>
			<u-cell icon="rmb-circle" title="我的学币" isLink url="/pages/my/coin/coin"></u-cell>
			<u-cell icon="star" title="我的收藏" isLink url="/pages/my/posts/favorites"></u-cell>
			<u-gap height="15" bg-color="#f9f9f9"></u-gap>
			<u-cell icon="question-circle" title="帮助与反馈" isLink url="/pages/my/helpAndFeedback/feedbackIndex"></u-cell>
			<u-cell icon="setting" title="设置" isLink url="/pages/my/settings/setting"></u-cell>
		</u-cell-group>

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
				this.$refs.uNotify.show({
					top: 10,
					type: 'success',
					color: '#000',
					bgColor: '#b2cf87',
					message: '刷新成功',
					duration: 1000 * 2,
					fontSize: 20,
					safeAreaInsetTop: true
				})
			}, 1000)
		},
		mounted(){
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
	.blankBox {
		height: 100rpx;
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
		font-size: 25px;
		color: crimson;
	}
</style>