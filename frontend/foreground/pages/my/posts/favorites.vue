<template>
	<view class="page">
		<!-- 弹出提示窗 -->
		<u-toast ref="uToast"></u-toast>
		<view v-for="(item,index) in favorites" :key="index" @click="favoritesClik(item)">
			<view class="list-box">
				<view class="text-delete" @click.stop="postsDelete(item.id,userID)">×</view>
				<view class="text-title">{{item.title}}</view>
				<view class="text-tips">{{item.content}}</view>
				<view class="item-bottom">
					<view class="favorite-icon">
						<u-icon name="star-fill" color="#909193" labelColor="#909193" :label="item.favorite"
							labelSize="7rpx"></u-icon>
					</view>
					<view class="like-icon">
						<u-icon name="thumb-up-fill" color="#909193" labelColor="#909193" :label="item.like"
							labelSize="7rpx"></u-icon>
					</view>
					<view class="time-box">
						<view class="text-time">{{item.publishTime}}</view>
					</view>
				</view>

			</view>
		</view>

	</view>
</template>

<script>
	export default {
		data() {
			return {
				userID: '', //用户id
				favorites: [{
					id: '', //文章id
					partID: '1', //文章类型
					title: '我的收藏1', //文章标题
					content: '这是我的收藏1', //文章内容
					favorite: '204', //收藏数
					like: '850', //点赞数
					publishTime: '2020-6-3 18:00:00', //发表时间
				}, {
					id: '', //文章id
					partID: '2', //文章类型
					title: '我的收藏2', //文章标题
					content: '这是我的收藏2', //文章内容
					favorite: '125', //收藏数
					like: '840', //点赞数
					publishTime: '2020-6-3 18:00:00', //发表时间
				}, ],
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
						this.userID = res.data;
						console.log(this.userID)
					}
				})
				uni.$u.http.get('/v1/frontend/user/myFavorites/' + this.userID, {

				}).then(res => {
					console.log("获取收藏成功！");
					console.log(res.data.data);
					this.favorites = res.data.data;
				}).catch(err => {
					console.log("获取收藏失败！！！");
				})
			},
			// 删除收藏
			postsDelete(article_ID,user_ID) {
				
				var _this=this;
				// 基本用法，注意：post的第三个参数才为配置项
				uni.$u.http.post('/v1/frontend/user/deleteMyFavorite', {
					userID:user_ID,
					articleID:article_ID
				}).then(res => {
					console.log("userID=="+user_ID);
					console.log("articleID=="+article_ID);
					console.log(res.data);
					this.$refs.uToast.show({
						type: 'success',
						message: "删除成功",
					})
					setTimeout(function() {
					  // 这里写要延时执行的代码
					   _this.refresh();
					}, 1500);
					
			
				}).catch(err => {
					console.log(err);
					this.$refs.uToast.show({
						type: 'error',
						message: "删除失败",
					})
				})
			
			},
			favoritesClik(item) {
				if (item.partID == '1') { //假设是加油站
					uni.navigateTo({
						//TODO
						//参数部分未修改
						url: '/pages/data/newsDetails/newsDetails?' + 'id=' + item.id,
						success: res => {},
						fail: () => {},
						complete: () => {}
					});
				} else if (item.partID == '2') { //假设是求助
					uni.navigateTo({
						//TODO
						//参数部分未修改
						url: '/pages/data/questionsDetails/questionsDetails?' + 'id=' + item.id,
						success: res => {},
						fail: () => {},
						complete: () => {}
					});
				}

			},
		}
	}
</script>

<style lang="scss">
	page {
		background-color: #fafafa;
	}

	.list-box {
		position: relative;
		background-color: #FFFFFF;
		border-radius: 10rpx;
		margin-top: 30rpx;
		margin-left: 30rpx;
		margin-right: 30rpx;
		padding: 30rpx;
		
		box-shadow: 5rpx 10rpx 10rpx #bfbfbf;
		border-radius: 20rpx;
	}

	.text-title {
		color: #303133;
		font-size: 30rpx;
		font-weight: bold;
		margin-right: 100rpx;
	}

	.time-box {
		width: 300px;
		display: flex;
		flex-direction: row-reverse;
	}

	.text-time {
		color: #909193;
		font-size: 24rpx;
		margin-top: 24rpx;
		margin-left: 0%;
	}

	.text-tips {
		color: #bcbcbc;
		font-size: 24rpx;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		margin-top: 22rpx;
	}

	.favorite-icon {
		margin-top: 24rpx;
	}

	.like-icon {
		margin-left: 2%;
		margin-top: 24rpx;
	}

	.item-bottom {
		display: flex;
		flex-direction: row;
	}

	.text-delete {
		float: right;
	}
</style>