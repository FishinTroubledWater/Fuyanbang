<template>
	<view>

		<view class="tabs-box">
			<view class="one-nav" style="padding: 10rpx;" :class="currentSwiperIndex === 0 ? 'nav-actived' : '' "
				@tap="swiperChange(0)">加油站
			</view>
			<view class="one-nav" style="padding: 10rpx;" :class="currentSwiperIndex === 1 ? 'nav-actived' : '' "
				@tap="swiperChange(1)">求解答
			</view>
		</view>
		<swiper class="swiper-box" style="height: 2000upx" :current="currentSwiperIndex">
			<swiper-item class="swiper-item sns-que">
				<view class="topic">
					<text class="contentTopWord">21天攒图计划招募令</text>
					<text>\n</text>
					<text style="color:blanchedalmond;">每日一画，不画就出局
						来吧，一起坚持21天，做行动派！</text>
					<navigator class="contentTopWord" url="/pages/data/activityDetail/activityDetail">了解详情</navigator>

				</view>

				<u-list>
					<u-list-item v-for="(item, index) in indexList" :key="index">
						<uni-card @click="clicknews(indexList[index].postId)" :title="indexList[index].name"
							sub-title="教育部" :extra="indexList[index].time" :thumbnail="indexList[index].icon"
							class="trends-box-item">
							<!-- <u--text :lines="3" :text="indexList[index].summary"></u--text>
								<image class="newsimage" :src="indexList[index].img[0]"></image> -->
							<view class="u-content">
								<u-parse :content="indexList[index].summary"></u-parse>
							</view>
						</uni-card>
					</u-list-item>
				</u-list>


			</swiper-item>
			<swiper-item class="swiper-item sns-oil">
				<u-list>
					<u-list-item v-for="(item, index) in questionsList" :key="index">
						<uni-card @click="clickquestions(questionsList[index].queId)" :title="questionsList[index].name"
							sub-title="教育部" :extra="questionsList[index].time" :thumbnail="questionsList[index].icon"
							class="trends-box-item">
							<view class="u-content">
								<u-parse :content="questionsList[index].summary"></u-parse>
							</view>
						</uni-card>
					</u-list-item>
				</u-list>
			</swiper-item>

		</swiper>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				content: `
					<p>露从今夜白，月是故乡明</p>
					<img src="../../static/background/activityDetails.png" />
				`,
				currentSwiperIndex: 0,
				list1: [{
					name: '求解答',
				}, {
					name: '加油站',
				}],


				indexList: [{

						name: '1',
						time: '2022-12-21',
						icon: '../../static/background/activityDetails.png',
						postId: '123456',
						summary: '<p>露从今夜白，月是故乡明</p>',
						isImage: true,
						img: ['../../static/background/activityDetails.png',
							'../../static/background/activityDetails.png',
							'../../static/background/activityDetails.png'
						]
					}, {
						name: '1',
						time: '2022-12-21',
						icon: '../../static/background/activityDetails.png',
						postId: '123457',
						summary: `
					<p>露从今夜白，月是故乡明</p>
					<img src="../../static/background/activityDetails.png" />
				`,
						isImage: true,
						img: ['www.baidu.com',
							'www.baidu.com',
							'www.baidu.com'
						]
					}, {
						name: '1',
						time: '2022-12-21',
						icon: '../../static/background/activityDetails.png',
						postId: '123456',
						summary: `
					<p>露从今夜白，月是故乡明</p>
					<img src="../../static/background/activityDetails.png" />
				`,
						isImage: true,
						img: ['www.baidu.com',
							'www.baidu.com',
							'www.baidu.com'
						]
					}, {
						name: '1',
						time: '2022-12-21',
						icon: '../../static/background/activityDetails.png',
						postId: '123456',
						summary: `
					<p>露从今夜白，月是故乡明</p>
					<img src="../../static/background/activityDetails.png" />
				`,
						isImage: true,
						img: ['www.baidu.com',
							'www.baidu.com',
							'www.baidu.com'
						]
					}, {
						name: '1',
						time: '2022-12-21',
						icon: '../../static/background/activityDetails.png',
						postId: '123456',
						summary: `
					<p>露从今夜白，月是故乡明</p>
					<img src="../../static/background/activityDetails.png" />
				`,
						isImage: true,
						img: ['www.baidu.com',
							'www.baidu.com',
							'www.baidu.com'
						]
					}, {
						name: '1',
						time: '2022-12-21',
						icon: '../../static/background/activityDetails.png',
						postId: '123456',
						summary: `
					<p>露从今夜白，月是故乡明</p>
					<img src="../../static/background/activityDetails.png" />
				`,
						isImage: true,
						img: ['www.baidu.com',
							'www.baidu.com',
							'www.baidu.com'
						]
					}, {
						name: '1',
						time: '2022-12-21',
						icon: '../../static/background/activityDetails.png',
						postId: '123456',
						summary: `
					<p>露从今夜白，月是故乡明</p>
					<img src="../../static/background/activityDetails.png" />
				`,
						isImage: true,
						img: ['www.baidu.com',
							'www.baidu.com',
							'www.baidu.com'
						]
					}, {
						name: '1',
						time: '2022-12-21',
						icon: '../../static/background/activityDetails.png',
						postId: '123456',
						summary: `
					<p>露从今夜白，月是故乡明</p>
					<img src="../../static/background/activityDetails.png" />
				`,
						isImage: true,
						img: ['www.baidu.com',
							'www.baidu.com',
							'www.baidu.com'
						]
					}, {
						name: '1',
						time: '2022-12-21',
						icon: '',
						postId: '123456',
						summary: `
					<p>露从今夜白，月是故乡明</p>
					<img src="../../static/background/activityDetails.png" />
				`,
						isImage: true,
						img: ['www.baidu.com',
							'www.baidu.com',
							'www.baidu.com'
						]
					}


				],
				questionsList: [{
						name: '1',
						time: '2022-12-21',
						icon: '../../static/background/activityDetails.png',
						queId: '123456',
						summary: `
							<p>露从今夜白</p>
							<img src="../../static/background/activityDetails.png" />
						`,
						isImage: true,
						img: ['../../static/background/activityDetails.png',
							'../../static/background/activityDetails.png',
							'../../static/background/activityDetails.png'
						]
					},
					{
						name: '2',
						time: '2022-12-21',
						icon: '../../static/background/activityDetails.png',
						queId: '123',
						summary: `
							<p>露从今夜白，月是故乡明</p>
							<img src="../../static/background/activityDetails.png" />
						`,
						isImage: true,
						img: ['../../static/background/activityDetails.png',
							'../../static/background/activityDetails.png',
							'../../static/background/activityDetails.png'
						]
					}
				]
			};
		},
		methods: {
			//求解答、加油站、关注切换方法
			swiperChange(index) {
				this.currentSwiperIndex = index
			},
			clicknews(index) {
				// console.log('1'),
				uni.navigateTo({
					url: '/pages/data/newsDetails/newsDetails?id=' + index,
				})
			},
			clickquestions(index) {
				uni.navigateTo({
					url: '/pages/data/questionsDetails/questionsDetails?id=' + index,
				})
			}

		}
	}
</script>

<style lang="scss">
	.u-content {
	        padding: 24rpx;
	}
	.newsimage {
		max-height: 200rpx;
		max-width: 200px;
	}

	.quanzi {
		font-size: 50rpx;
		font-weight: 800;
	}

	.tabs-box {
		flex-flow: row;
		justify-content: space-around;
		display: flex;
	}

	.nav-actived {
		color: blanchedalmond;
		font-weight: 700;
	}

	.appContentTop {
		z-index: -1;
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 200rpx;
	}

	.contentTopWord {
		color: #ffffff;
		font-size: 40rpx;
	}

	.contentPic {
		object-fit: cover;
		height: 100%;
		width: 100%;
	}

	.topic {
		background: url("../../static/background/bg2.png") center center;
	}
</style>