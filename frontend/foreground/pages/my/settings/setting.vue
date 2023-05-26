<template>
	<view class="container">
		<view class="list-cell b-b m-t" @click="gotoPage('/pages/my/settings/userInfo')" hover-class="cell-hover"
			:hover-stay-time="50">
			<text class="cell-tit">个人资料</text>
			<text class="cell-more yticon icon-you"></text>
			<text>></text>
		</view>
		<view class="list-cell b-b" @click="gotoPage('/pages/my/settings/accountSecurity')" hover-class="cell-hover"
			:hover-stay-time="50">
			<text class="cell-tit">账号与安全</text>
			<text class="cell-more yticon icon-you"></text>
			<text>></text>
		</view>
<!-- 
		<view class="list-cell m-t">
			<text class="cell-tit">消息推送</text>
			<switch checked color="#fa436a" @change="switchChange" />
		</view> -->
		<view class="list-cell m-t b-b" @click="cleanCache" hover-class="cell-hover" :hover-stay-time="50">
			<text class="cell-tit">清除缓存</text>
			<text class="cell-tip">{{cache}}M</text>
			<text class="cell-more yticon icon-you"></text>
		</view>
		<view class="list-cell b-b" @click="gotoPage('/pages/my/settings/aboutUs')" hover-class="cell-hover"
			:hover-stay-time="50">
			<text class="cell-tit">关于我们</text>
			<text class="cell-more yticon icon-you"></text>
			<text>></text>
		</view>
		<view class="list-cell log-out-btn" @click="toLogout">
			<text class="cell-tit">退出登录</text>
		</view>
	</view>
</template>

<script>
	import {
		mapMutations
	} from 'vuex';
import { getTransitionRawChildren } from "vue";
	export default {
		data() {
			return {
				max:100,
				min:12,
				rand:0,
				cache:105.2

			};
		},
		onLoad() {
			this.getRand();
			this.cache +=this.rand;
		},
		methods: {
			cleanCache() {
				this.cache = 12.8;
				uni.showToast({
					title: '清除缓存成功',
					//将值设置为 success 或者直接不用写icon这个参数
					icon: 'success',
					//显示持续时间为 1.5秒
					duration: 1500
				})
			},
			logout() {
				uni.clearStorage();
				setTimeout(() => {
					uni.navigateTo({
						url: '/pages/login/passwordLogin'
					})
				}, 200)
			},
			gotoPage(url) {
				uni.navigateTo({
					url
				})
			},
			//退出登录
			toLogout() {
				uni.showModal({
					content: '确定要退出登录么',
					success: (e) => {
						if (e.confirm) {
							this.logout();
						}
					}
				});
			},
			//switch
			switchChange(e) {
				let statusTip = e.detail.value ? '打开' : '关闭';
				this.$api.msg(`${statusTip}消息推送`);
			},
			getRand() {
				console.log(this.min)
				console.log(this.max)
				this.rand = Math.floor(Math.random() * (this.max - this.min + 1)) + this.min

			}

		}
	}
</script>

<style lang='scss'>
	page {
		background: $page-color-base;
	}

	.list-cell {
		display: flex;
		align-items: baseline;
		padding: 20upx $page-row-spacing;
		line-height: 60upx;
		position: relative;
		background: #fff;
		justify-content: center;

		&.log-out-btn {
			margin-top: 40upx;

			.cell-tit {
				color: $uni-color-primary;
				text-align: center;
				margin-right: 0;
			}
		}

		&.cell-hover {
			background: #fafafa;
		}

		&.b-b:after {
			left: 30upx;
		}

		&.m-t {
			margin-top: 16upx;
		}

		.cell-more {
			align-self: baseline;
			font-size: $font-lg;
			color: $font-color-light;
			margin-left: 10upx;
		}

		.cell-tit {
			flex: 1;
			font-size: $font-base + 2upx;
			color: $font-color-dark;
			margin-right: 10upx;
		}

		.cell-tip {
			font-size: $font-base;
			color: $font-color-light;
		}

		switch {
			transform: translateX(16upx) scale(.84);
		}
	}
</style>