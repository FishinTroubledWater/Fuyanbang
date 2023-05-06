<template>
	<view class="content">
		<!-- 个人信息 -->
		<u-cell-group>
			<u-cell>
				<u-icon @click="gotoPage('/pages/my/setting')" slot="value" name="setting" color="#6f6f6f"
					size="22"></u-icon>
			</u-cell>
			<u-cell isLink >
				<u-image @click="changeHead()" width='140rpx' height='140rpx' slot="icon" :src="headImg"
					shape="circle"></u-image>
				<view class="box" slot="title">
					<view class="box1">
						<text style="margin-right:10px;">用户名</text>
						<u-tag  text="等级" size="mini" shape="circle"></u-tag>
					</view>
					<text style="color: darkgray;">个性签名</text>
				</view>
			</u-cell>
			<u-cell>
				<text slot="title">本科院校</text>
				<text slot="title">本科专业</text>
				<text slot="value">福研帮已经陪伴了您</text>
				<text slot="value" class="days">多少</text>
				<text slot="value">天了</text>
			</u-cell>
			<u-gap height="15" bg-color="#f9f9f9"></u-gap>
			<u-cell icon="edit-pen" title="我的创作" isLink url="/pages/my/posts"></u-cell>
			<u-cell icon="rmb-circle" title="我的学币" isLink url="/pages/my/coin"></u-cell>
			<u-cell icon="star" title="我的收藏" isLink url="/pages/my/favorites"></u-cell>
			<u-gap height="15" bg-color="#f9f9f9"></u-gap>
			<u-cell icon="question-circle" title="帮助" isLink url="/pages/my/help"></u-cell>
		</u-cell-group>

	</view>
</template>

<script>
	export default {
		data() {
			return {
				headImg: '/static/my-assets/泰裤辣.png',
			}
		},
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
	.box1{
		display: flex;
		flex-direction: row;
	}
	.days{
		font-size: 25px;
		color: crimson;
	}
</style>